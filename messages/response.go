package messages

import "time"

// VODInfo ...
type vodInfo struct {
	Title      string    `json:"title"`
	ID         string    `json:"_id"`
	RecordedAt time.Time `json:"recorded_at"`
	Game       string    `json:"game"`
	Length     int       `json:"length"`
	Channel    channel   `json:"channel"`
}

// Channel ...
type channel struct {
	DisplayName string `json:"display_name"`
	ID          int    `json:"_id"`
}

// MessageChunk ...
type messageChunk struct {
	Comments []comment `json:"comments"`
	Next     string    `json:"_next"`
}

// Comment ...
type comment struct {
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
