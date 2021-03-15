package messages

import (
	"fmt"

	"github.com/harry93848bb7/chat-archiver/badges"
	"github.com/harry93848bb7/chat-archiver/emotes"
)

// ArchiveChat ...
func ArchiveChat(vodID, clientID string) (*Archive, error) {
	c := NewClient(clientID)

	// Basic VOD Info
	vodInfo, err := c.GetVODInfo(vodID)
	if err != nil {
		return &Archive{}, err
	}

	// User and Global VOD Badges
	userBadges, err := badges.UserBadges(fmt.Sprintf("%d", vodInfo.Channel.ID))
	if err != nil {
		return &Archive{}, err
	}
	globalBadges, err := badges.TwitchGlobalBadges()
	if err != nil {
		return &Archive{}, err
	}

	archive := &Archive{
		VODID:       vodID,
		Title:       vodInfo.Title,
		Length:      vodInfo.Length,
		Category:    vodInfo.Game,
		RecordedAt:  vodInfo.RecordedAt.Unix(),
		ChannelName: vodInfo.Channel.DisplayName,
		ChannelID:   fmt.Sprintf("%d", vodInfo.Channel.ID),
		Badge:       userBadges,
	}
	archive.Badge = append(archive.Badge, globalBadges...)

	// VOD Messages
	next := ""
	for {
		data, err := c.GetMessageChunk(next, vodID)
		if err != nil {
			return &Archive{}, err
		}
		for _, c := range data.Comments {
			m := Message{
				ContentOffSet: c.ContentOffsetSeconds,
				DisplayName:   c.Commenter.DisplayName,
				DisplayColor:  c.Message.UserColor,
				Message:       c.Message.Body,
			}
			for _, b := range c.Message.UserBadges {
				m.Badges = append(m.Badges, Badge{
					ID:      b.ID,
					Version: b.Version,
				})
			}
			archive.Messages = append(archive.Messages, m)
		}
		if data.Next == "" {
			break
		}
		next = data.Next
	}

	// All Global and Channel Emotes
	emote, err := emotes.ArchiveEmotes(fmt.Sprintf("%d", vodInfo.Channel.ID))
	if err != nil {
		return &Archive{}, err
	}
	archive.Emotes = emote

	return archive, nil
}
