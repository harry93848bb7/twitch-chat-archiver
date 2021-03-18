package emotes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/harry93848bb7/chat-archiver/protobuf"
	"github.com/harry93848bb7/chat-archiver/sterilise"
)

// BTTVGlobal ...
func BTTVGlobal() ([]*protobuf.Emote, error) {
	r, err := http.Get("https://api.betterttv.net/3/cached/emotes/global")
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		log.Println("Error retrieving all Global BetterTTV emotes")
		return nil, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data := []struct {
		ID        string `json:"id"`
		Code      string `json:"code"`
		ImageType string `json:"imageType"`
		UserID    string `json:"userId"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var emotes = []*protobuf.Emote{}
	for _, emote := range data {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/1x", emote.ID))
		if err != nil {
			return nil, err
		}
		if r.StatusCode != 200 {
			log.Println("Error retrieving BetterTTV emoticon id", emote.ID)
			continue
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		data, format, err := sterilise.SteriliseImage(b)
		if err == sterilise.UnknownFormat {
			log.Println("Unknown emote image file format:", emote.Code)
			continue
		} else if err != nil {
			return nil, err
		}
		emotes = append(emotes, &protobuf.Emote{
			Code:      emote.Code,
			Source:    "BetterTTV Global Emotes",
			ImageType: format,
			ImageData: data,
		})
	}
	return emotes, nil
}

// BTTVChannel ...
func BTTVChannel(channelID string) ([]*protobuf.Emote, error) {
	r, err := http.Get(fmt.Sprintf("https://api.betterttv.net/3/cached/users/twitch/%s", channelID))
	if err != nil {
		return nil, err
	}
	if r.StatusCode == http.StatusNotFound {
		log.Println("No BetterTTV emotes found for channel", channelID)
		return nil, nil
	}
	if r.StatusCode != 200 {
		log.Println("Error retrieving BetterTTV channel emotes")
		return nil, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data := struct {
		ID            string   `json:"id"`
		Bots          []string `json:"bots"`
		ChannelEmotes []struct {
			ID        string `json:"id"`
			Code      string `json:"code"`
			ImageType string `json:"imageType"`
			UserID    string `json:"userId"`
		} `json:"channelEmotes"`
		SharedEmotes []struct {
			ID        string `json:"id"`
			Code      string `json:"code"`
			ImageType string `json:"imageType"`
			User      struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				DisplayName string `json:"displayName"`
				ProviderID  string `json:"providerId"`
			} `json:"user"`
		} `json:"sharedEmotes"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var emotes = []*protobuf.Emote{}
	for _, emote := range data.ChannelEmotes {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/1x", emote.ID))
		if err != nil {
			return nil, err
		}
		if r.StatusCode != 200 {
			log.Println("Error retrieving BetterTTV emoticon id", emote.ID)
			continue
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		data, format, err := sterilise.SteriliseImage(b)
		if err == sterilise.UnknownFormat {
			log.Println("Unknown emote image file format:", emote.Code)
			continue
		} else if err != nil {
			return nil, err
		}
		emotes = append(emotes, &protobuf.Emote{
			Code:      emote.Code,
			Source:    "BetterTTV Channel Emotes",
			ImageType: format,
			ImageData: data,
		})
	}
	for _, emote := range data.SharedEmotes {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/1x", emote.ID))
		if err != nil {
			return nil, err
		}
		if r.StatusCode != 200 {
			log.Println("Error retrieving BetterTTV emoticon id", emote.ID)
			continue
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		data, format, err := sterilise.SteriliseImage(b)
		if err == sterilise.UnknownFormat {
			log.Println("Unknown emote image file format:", emote.Code)
			continue
		} else if err != nil {
			return nil, err
		}
		emotes = append(emotes, &protobuf.Emote{
			Code:      emote.Code,
			Source:    "BetterTTV Channel Emotes",
			ImageType: format,
			ImageData: data,
		})
	}
	return emotes, nil
}

// FFZGlobal ...
func FFZGlobal() ([]*protobuf.Emote, error) {
	r, err := http.Get("https://api.frankerfacez.com/v1/set/global")
	if err != nil {
		return nil, err
	}
	if r.StatusCode != 200 {
		log.Println("Error retrieving all Global FrankerFaceZ emotes")
		return nil, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data := struct {
		Sets map[string]struct {
			Title     string `json:"title"`
			Emoticons []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"emoticons"`
		} `json:"sets"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var emotes = []*protobuf.Emote{}
	for _, set := range data.Sets {
		if set.Title == "Global Emotes" {
			for _, emote := range set.Emoticons {
				r, err := http.Get(fmt.Sprintf("https://cdn.frankerfacez.com/emoticon/%d/1", emote.ID))
				if err != nil {
					return nil, err
				}
				if r.StatusCode != 200 {
					log.Println("Error retrieving FrankerFaceZ emoticon id", emote.ID)
					continue
				}
				b, err := ioutil.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				data, format, err := sterilise.SteriliseImage(b)
				if err == sterilise.UnknownFormat {
					log.Println("Unknown emote image file format:", emote.Name)
					continue
				} else if err != nil {
					return nil, err
				}
				emotes = append(emotes, &protobuf.Emote{
					Code:      emote.Name,
					Source:    "FrankerFaceZ Global Emotes",
					ImageType: format,
					ImageData: data,
				})
			}
		}
	}
	return emotes, nil
}

// FFZChannel ...
func FFZChannel(channelID string) ([]*protobuf.Emote, error) {
	r, err := http.Get(fmt.Sprintf("https://api.frankerfacez.com/v1/room/id/%s", channelID))
	if err != nil {
		return nil, err
	}
	if r.StatusCode == http.StatusNotFound {
		log.Println("No FrankerFaceZ emotes found for channel", channelID)
		return nil, nil
	}
	if r.StatusCode != 200 {
		log.Println("Error retrieving FrankerFaceZ channel emotes")
		return nil, nil
	}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	data := struct {
		Sets map[string]struct {
			Title     string `json:"title"`
			Emoticons []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"emoticons"`
		} `json:"sets"`
	}{}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var emotes = []*protobuf.Emote{}
	for _, channel := range data.Sets {
		for _, emote := range channel.Emoticons {
			r, err := http.Get(fmt.Sprintf("https://cdn.frankerfacez.com/emoticon/%d/1", emote.ID))
			if err != nil {
				return nil, err
			}
			if r.StatusCode != 200 {
				log.Println("Error retrieving FrankerFaceZ emoticon id", emote.ID)
				continue
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			data, format, err := sterilise.SteriliseImage(b)
			if err == sterilise.UnknownFormat {
				log.Println("Unknown emote image file format:", emote.Name)
				continue
			} else if err != nil {
				return nil, err
			}
			emotes = append(emotes, &protobuf.Emote{
				Code:      emote.Name,
				Source:    "FrankerFaceZ " + channel.Title,
				ImageType: format,
				ImageData: data,
			})
		}
	}
	return emotes, nil
}

// TwitchGlobal ...
func TwitchGlobal() ([]*protobuf.Emote, error) {
	var emotes = []*protobuf.Emote{}
	directory, err := twitchGlobal.ReadDir("twitchglobal")
	if err != nil {
		return nil, err
	}
	for _, entry := range directory {
		code := emoteMapping[entry.Name()]
		if code == "" {
			return nil, fmt.Errorf("emote mapping not found")
		}
		b, err := twitchGlobal.ReadFile("twitchglobal/" + entry.Name())
		if err != nil {
			return nil, err
		}
		data, format, err := sterilise.SteriliseImage(b)
		if err == sterilise.UnknownFormat {
			log.Println("Unknown emote image file format:", code)
			continue
		} else if err != nil {
			return nil, err
		}
		emotes = append(emotes, &protobuf.Emote{
			Code:      code,
			Source:    "Twitch Global",
			ImageType: format,
			ImageData: data,
		})
	}
	return emotes, nil
}

// Channel ...
func Channel(channelID string) ([]*protobuf.Emote, error) {
	r, err := http.Get("https://api.twitchemotes.com/api/v4/channels/" + channelID)
	if err != nil {
		return nil, err
	}
	if r.StatusCode == http.StatusNotFound {
		log.Println("No Subscription emotes found for channel", channelID)
		return nil, nil
	}
	if r.StatusCode != 200 {
		log.Println("Error retrieving Channel Subscription emotes")
		return nil, nil
	}
	d := struct {
		DisplayName string `json:"display_name"`
		Emotes      []struct {
			Code        string `json:"code"`
			EmoticonSet int64  `json:"emoticon_set"`
			ID          int64  `json:"id"`
		} `json:"emotes"`
	}{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &d); err != nil {
		return nil, err
	}
	var emotes = []*protobuf.Emote{}
	for _, e := range d.Emotes {

		r, err := http.Get(fmt.Sprintf("https://static-cdn.jtvnw.net/emoticons/v1/%d/1.0", e.ID))
		if err != nil {
			return nil, err
		}
		if r.StatusCode != 200 {
			log.Println("Error retrieving Channel Subscription emoticon id", e.ID)
			continue
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		data, format, err := sterilise.SteriliseImage(b)
		if err == sterilise.UnknownFormat {
			log.Println("Unknown emote image file format:", e.Code)
			continue
		} else if err != nil {
			return nil, err
		}
		emotes = append(emotes, &protobuf.Emote{
			Code:      e.Code,
			Source:    "Channel: " + d.DisplayName,
			ImageType: format,
			ImageData: data,
		})
	}
	return emotes, nil
}
