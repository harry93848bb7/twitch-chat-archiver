package archiver

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/harry93848bb7/chat-archiver/badges"
	"github.com/harry93848bb7/chat-archiver/emotes"
	"github.com/harry93848bb7/chat-archiver/messages"
	"github.com/harry93848bb7/chat-archiver/protobuf"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// Emotes ...
func Emotes(channelID string) ([]*protobuf.Emote, error) {
	var emote = []*protobuf.Emote{}

	bttvglobal, err := emotes.BTTVGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, bttvglobal...)
	bttvuser, err := emotes.BTTVChannel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, bttvuser...)
	ffzglobal, err := emotes.FFZGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, ffzglobal...)
	ffzuser, err := emotes.FFZChannel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, ffzuser...)
	twitchglobal, err := emotes.TwitchGlobal()
	if err != nil {
		return nil, err
	}
	emote = append(emote, twitchglobal...)
	robot, err := emotes.Robot()
	if err != nil {
		return nil, err
	}
	emote = append(emote, robot...)
	channel, err := emotes.Channel(channelID)
	if err != nil {
		return nil, err
	}
	emote = append(emote, channel...)

	// Provide warning logs if emotes are duplicate across different sources
	for i := range emote {
		emoteFound := 0
		for d := range emote {
			if emote[i].Code == emote[d].Code {
				emoteFound++
			}
		}
		if emoteFound > 1 {
			log.Printf(`Warning: Duplicate emote "%s" found across sources`, emote[i].Code)
		}
	}
	return emote, nil
}

// Badges ...
func Badges(channelID string) ([]*protobuf.Badge, error) {
	var badge = []*protobuf.Badge{}

	global, err := badges.TwitchGlobal()
	if err != nil {
		return nil, err
	}
	channel, err := badges.Channel(channelID)
	if err != nil {
		return nil, err
	}

	// Remove twitch global badges when a channel badge exists

	// - A channel can replace the global subscription badge with custom channel subscription badges
	// - A channel can replace the global bit badge with custom channel bit badges
	// - A channel can replace the global cheer badge with custom channel cheer badges

	deleted := 0
	for g := range global {
		for c := range channel {
			i := g - deleted

			if (global[i].Code == channel[c].Code) && (global[i].Version == channel[c].Version) {
				global = append(global[:i], global[i+1:]...)
				deleted++
			}
		}
	}
	log.Printf("Replace %d twitch badges in favor of channel badges\n", deleted)

	badge = append(badge, global...)
	badge = append(badge, channel...)

	// Provide warning logs if badges are duplicate across different sources
	for i := range badge {
		badgeCount := 0
		for d := range badge {
			if (badge[i].Code == badge[d].Code) && (badge[i].Version == badge[d].Version) {
				badgeCount++
				if badgeCount > 1 {
					log.Printf(`Warning: Duplicate badge "%s" version "%s" found across sources`+"\n", badge[i].Code, badge[i].Version)
				}
			}
		}
	}
	return badge, nil
}

// Messages ...
func Messages(vodID, clientID string) ([]*protobuf.Message, error) {
	var message = []*protobuf.Message{}

	client := messages.NewClient(clientID)

	// Continue to move through the message chunks until next is not present
	next := ""
	for {
		chunk, err := client.GetMessageChunk(next, vodID)
		if err != nil {
			return nil, err
		}
		for _, c := range chunk.Comments {
			var badges = []*protobuf.MessageBadge{}

			for _, b := range c.Message.UserBadges {
				badges = append(badges, &protobuf.MessageBadge{
					Code:    b.ID,
					Version: b.Version,
				})
			}
			message = append(message, &protobuf.Message{
				ContentOffset: c.ContentOffsetSeconds,
				DisplayName:   c.Commenter.DisplayName,
				DisplayColor:  c.Message.UserColor,
				Content:       c.Message.Body,
				Badges:        badges,
			})
		}
		if chunk.Next == "" {
			break
		}
		next = chunk.Next
	}
	return message, nil
}

// Archive ...
func Archive(vodID, clientID string) (*protobuf.Archive, error) {
	client := messages.NewClient(clientID)

	vod, err := client.GetVODInfo(vodID)
	if err != nil {
		log.Println("Failed to get VOD information", err)
		return &protobuf.Archive{}, err
	}
	vID, err := strconv.ParseInt(vodID, 10, 64)
	if err != nil {
		return &protobuf.Archive{}, err
	}
	badge, err := Badges(fmt.Sprintf("%d", vod.Channel.ID))
	if err != nil {
		return &protobuf.Archive{}, err
	}
	emote, err := Emotes(fmt.Sprintf("%d", vod.Channel.ID))
	if err != nil {
		return &protobuf.Archive{}, err
	}
	message, err := Messages(vodID, clientID)
	if err != nil {
		return &protobuf.Archive{}, err
	}

	// Remove badges which are not being used in any of the chat messages to reduce file size as they are redundant
	deleted := 0
	for g := range badge {
		i := g - deleted
		badgeOccurance := 0

		for _, m := range message {
			for _, messageBadge := range m.Badges {
				if (messageBadge.Code == badge[i].Code) && (messageBadge.Version == badge[i].Version) {
					badgeOccurance++
				}
			}
		}
		if badgeOccurance == 0 {
			badge = append(badge[:i], badge[i+1:]...)
			deleted++
		}
	}
	log.Printf("Successfully remove %d badges not being used in the archive messages", deleted)

	// Remove emotes which are not being used in any of the chat messages to reduce file size as they are redundant
	deleted = 0
	for g := range emote {
		i := g - deleted
		emoteOccurance := 0

		for _, m := range message {
			if strings.Index(m.Content, emote[i].Code) != -1 {
				emoteOccurance++
			}
		}
		if emoteOccurance == 0 {
			emote = append(emote[:i], emote[i+1:]...)
			deleted++
		}
	}
	log.Printf("Successfully remove %d emotes not being used in the archive messages", deleted)

	return &protobuf.Archive{
		VodId:      vID,
		Title:      vod.Title,
		Category:   vod.Game,
		Length:     timestamppb.New(time.Unix(int64(vod.Length), 0)),
		RecordedAt: timestamppb.New(vod.RecordedAt),
		Channel: &protobuf.Channel{
			Name: vod.Channel.DisplayName,
			Id:   int64(vod.Channel.ID),
		},
		Badges:   badge,
		Emotes:   emote,
		Messages: message,
	}, nil
}
