package badges

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Badge ...
type Badge struct {
	Code          string `json:"code"`
	Version       string `json:"version"`
	Title         string `json:"title"`
	ImageType     string `json:"image_type"`
	Base64Encoded string `json:"base64_encoded"`
}

// Badges ...
type Badges struct {
	BadgeSets map[string]Versions `json:"badge_sets"`
}

// Version ...
type Version map[string]TwitchBadge

// Versions ...
type Versions map[string]map[string]TwitchBadge

// TwitchBadge ...
type TwitchBadge struct {
	ImageURL1X  string `json:"image_url_1x"`
	ImageURL2X  string `json:"image_url_2x"`
	ImageURL4X  string `json:"image_url_4x"`
	Description string `json:"description"`
	Title       string `json:"title"`
	ClickAction string `json:"click_action"`
	ClickURL    string `json:"click_url"`
}

// GlobalBadges ...
func TwitchGlobalBadges() ([]Badge, error) {
	response, err := http.Get("https://badges.twitch.tv/v1/badges/global/display?language=en")
	if err != nil {
		return nil, err
	}
	var data Badges
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var badges = []Badge{}
	for code, badge := range data.BadgeSets {
		for number, version := range badge["versions"] {
			r, err := http.Get(version.ImageURL1X)
			if err != nil {
				return nil, err
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			badges = append(badges, Badge{
				Code:          code,
				Version:       number,
				Title:         version.Title,
				ImageType:     "png",
				Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
			})
		}
	}
	return badges, nil
}

// UserBadges ...
func UserBadges(userID string) ([]Badge, error) {
	response, err := http.Get(fmt.Sprintf("https://badges.twitch.tv/v1/badges/channels/%s/display?language=en", userID))
	if err != nil {
		return nil, err
	}
	var data Badges
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	var badges = []Badge{}
	for code, badge := range data.BadgeSets {
		for number, version := range badge["versions"] {
			r, err := http.Get(version.ImageURL1X)
			if err != nil {
				return nil, err
			}
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return nil, err
			}
			badges = append(badges, Badge{
				Code:          code,
				Version:       number,
				Title:         version.Title,
				ImageType:     "png",
				Base64Encoded: base64.RawStdEncoding.EncodeToString(b),
			})
		}
	}
	return badges, nil
}
