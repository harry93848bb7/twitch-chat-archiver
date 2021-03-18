package badges

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/harry93848bb7/chat-archiver/protobuf"
	"github.com/harry93848bb7/chat-archiver/sterilise"
)

// TwitchGlobal ...
func TwitchGlobal() ([]*protobuf.Badge, error) {
	response, err := http.Get("https://badges.twitch.tv/v1/badges/global/display?language=en")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		log.Println("Error retrieving all Global Twitch Badges")
		return nil, nil
	}
	var data badges
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var badges = []*protobuf.Badge{}
	for code, badge := range data.BadgeSets {
		for number, version := range badge["versions"] {
			r, err := http.Get(version.ImageURL1X)
			if err != nil {
				return nil, err
			}
			if r.StatusCode != 200 {
				log.Println("Error retrieving Global Twitch Badge", code)
				continue
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			data, format, err := sterilise.SteriliseImage(b)
			if err == sterilise.UnknownFormat {
				log.Println("Unknown badge image file format:", version.Title)
				continue
			} else if err != nil {
				return nil, err
			}
			badges = append(badges, &protobuf.Badge{
				Code:      code,
				Version:   number,
				Title:     version.Title,
				ImageType: format,
				ImageData: data,
			})
		}
	}
	return badges, nil
}

// Channel ...
func Channel(userID string) ([]*protobuf.Badge, error) {
	response, err := http.Get(fmt.Sprintf("https://badges.twitch.tv/v1/badges/channels/%s/display?language=en", userID))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		log.Println("Error retrieving all User Twitch Badges")
		return nil, nil
	}
	var data badges
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	if len(data.BadgeSets) == 0 {
		log.Println("No User Twitch Badges found for user", userID)
		return nil, nil
	}
	var badges = []*protobuf.Badge{}
	for code, badge := range data.BadgeSets {
		for number, version := range badge["versions"] {
			r, err := http.Get(version.ImageURL1X)
			if err != nil {
				return nil, err
			}
			if r.StatusCode != 200 {
				log.Println("Error retrieving User Twitch Badge", code)
				continue
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			data, format, err := sterilise.SteriliseImage(b)
			if err == sterilise.UnknownFormat {
				log.Println("Unknown badge image file format:", version.Title)
				continue
			} else if err != nil {
				return nil, err
			}
			badges = append(badges, &protobuf.Badge{
				Code:      code,
				Version:   number,
				Title:     version.Title,
				ImageType: format,
				ImageData: data,
			})
		}
	}
	return badges, nil
}
