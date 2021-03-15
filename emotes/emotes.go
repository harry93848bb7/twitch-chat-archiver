package emotes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Emote ...
type Emote struct {
	Code          string `json:"code"`
	Source        string `json:"source"`
	ImageType     string `json:"image_type"`
	Base64Encoded string `json:"base64_encoded"`
}

// ArchiveEmotes ...
func ArchiveEmotes(channelID string) ([]Emote, error) {
	var emotes = []Emote{}

	bttvUser, err := BTTVUser(channelID)
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, bttvUser...)

	ffzUser, err := FFZUser(channelID)
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ffzUser...)

	ffzGlobal, err := FFZGlobal()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ffzGlobal...)

	bttvGlobal, err := BTTVGlobal()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, bttvGlobal...)

	ttvGlobal, err := TwitchGlobal()
	if err != nil {
		return nil, err
	}
	emotes = append(emotes, ttvGlobal...)

	return emotes, nil
}

// BTTVGlobal ...
func BTTVGlobal() ([]Emote, error) {
	r, err := http.Get("https://api.betterttv.net/3/cached/emotes/global")
	if err != nil {
		return nil, err
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
	var emotes = []Emote{}
	for _, emote := range data {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/3x", emote.ID))
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		emotes = append(emotes, Emote{
			Code:          emote.Code,
			Source:        "BetterTTV Global Emotes",
			ImageType:     emote.ImageType,
			Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
		})
	}
	return emotes, nil
}

// BTTVUser ...
func BTTVUser(userID string) ([]Emote, error) {
	r, err := http.Get(fmt.Sprintf("https://api.betterttv.net/3/cached/users/twitch/%s", userID))
	if err != nil {
		return nil, err
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
	var emotes = []Emote{}
	for _, emote := range data.ChannelEmotes {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/3x", emote.ID))
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		emotes = append(emotes, Emote{
			Code:          emote.Code,
			Source:        "BetterTTV Channel Emotes",
			ImageType:     emote.ImageType,
			Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
		})
	}
	for _, emote := range data.SharedEmotes {
		r, err := http.Get(fmt.Sprintf("https://cdn.betterttv.net/emote/%s/3x", emote.ID))
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		emotes = append(emotes, Emote{
			Code:          emote.Code,
			Source:        "BetterTTV Channel Emotes",
			ImageType:     emote.ImageType,
			Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
		})
	}
	return emotes, nil
}

// FFZGlobal ...
func FFZGlobal() ([]Emote, error) {
	r, err := http.Get("https://api.frankerfacez.com/v1/set/global")
	if err != nil {
		return nil, err
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
	var emotes = []Emote{}
	for _, set := range data.Sets {
		if set.Title == "Global Emotes" {
			for _, emote := range set.Emoticons {
				r, err := http.Get(fmt.Sprintf("https://cdn.frankerfacez.com/emoticon/%d/1", emote.ID))
				if err != nil {
					return nil, err
				}
				b, err := ioutil.ReadAll(r.Body)
				if err != nil {
					return nil, err
				}
				emotes = append(emotes, Emote{
					Code:          emote.Name,
					Source:        "FrankerFaceZ Global Emotes",
					ImageType:     "png",
					Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
				})
			}
		}
	}
	return emotes, nil
}

// FFZUser ...
func FFZUser(userID string) ([]Emote, error) {
	r, err := http.Get(fmt.Sprintf("https://api.frankerfacez.com/v1/room/id/%s", userID))
	if err != nil {
		return nil, err
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
	var emotes = []Emote{}
	for _, channel := range data.Sets {
		for _, emote := range channel.Emoticons {
			r, err := http.Get(fmt.Sprintf("https://cdn.frankerfacez.com/emoticon/%d/1", emote.ID))
			if err != nil {
				return nil, err
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			emotes = append(emotes, Emote{
				Code:          emote.Name,
				Source:        "FrankerFaceZ " + channel.Title,
				ImageType:     "png",
				Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
			})
		}
	}
	return emotes, nil
}

// TwitchGlobal ...
func TwitchGlobal() ([]Emote, error) {
	var emotes = []Emote{}

	directory, err := twitchGlobal.ReadDir("twitchglobal")
	if err != nil {
		return nil, err
	}
	for _, entry := range directory {

		code := emoteMapping[entry.Name()]
		if code == "" {
			fmt.Println(entry.Name())
			return nil, fmt.Errorf("emote mapping not found")
		}

		b, err := twitchGlobal.ReadFile("twitchglobal/" + entry.Name())
		if err != nil {
			return nil, err
		}

		emotes = append(emotes, Emote{
			Code:          code,
			Source:        "Twitch Global",
			ImageType:     "png",
			Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
		})
	}
	return emotes, nil
}
