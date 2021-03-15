package messages

import (
	"github.com/harry93848bb7/chat-archiver/badges"
	"github.com/harry93848bb7/chat-archiver/emotes"
)

// Archive ...
type Archive struct {
	VODID      string `json:"vodid"`
	Title      string `json:"title"`
	Length     int    `json:"length"`
	Category   string `json:"category"`
	RecordedAt int64  `json:"recorded_at"`

	ChannelName string `json:"channel_name"`
	ChannelID   string `json:"channel_id"`

	Badge    []badges.Badge `json:"badges,omitempty"`
	Emotes   []emotes.Emote `json:"emotes,omitempty"`
	Messages []Message      `json:"messages,omitempty"`
}

// Message ...
type Message struct {
	ContentOffSet float64 `json:"content_offset"`
	DisplayName   string  `json:"display_name"`
	DisplayColor  string  `json:"display_color"`
	Badges        []Badge `json:"badges,omitempty"`
	Message       string  `json:"message"`
}

// Badge ...
type Badge struct {
	ID      string
	Version string
}
