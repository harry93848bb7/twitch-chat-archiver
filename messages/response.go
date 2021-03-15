package messages

import "time"

// GlobalBadges ...
type GlobalBadges struct {
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

// VODInfo ...
type VODInfo struct {
	Title      string    `json:"title"`
	ID         string    `json:"_id"`
	RecordedAt time.Time `json:"recorded_at"`
	Game       string    `json:"game"`
	Length     int       `json:"length"`
	Channel    Channel   `json:"channel"`
}

// Channel ...
type Channel struct {
	DisplayName string `json:"display_name"`
	ID          int    `json:"_id"`
}

// MessageChunk ...
type MessageChunk struct {
	Comments []Comment `json:"comments"`
	Next     string    `json:"_next"`
}

// Comment ...
type Comment struct {
	ContentOffsetSeconds float64 `json:"content_offset_seconds"`
	Commenter            struct {
		DisplayName string `json:"display_name"`
	} `json:"commenter"`
	Message struct {
		Body       string `json:"body"`
		UserBadges []struct {
			ID      string `json:"_id"`
			Version string `json:"version"`
		} `json:"user_badges"`
		UserColor string `json:"user_color"`
	} `json:"message"`
}
