package badges

// Badges ...
type badges struct {
	BadgeSets map[string]versions `json:"badge_sets"`
}

// Version ...
type version map[string]twitchBadge

// Versions ...
type versions map[string]map[string]twitchBadge

// TwitchBadge ...
type twitchBadge struct {
	ImageURL1X  string `json:"image_url_1x"`
	ImageURL2X  string `json:"image_url_2x"`
	ImageURL4X  string `json:"image_url_4x"`
	Description string `json:"description"`
	Title       string `json:"title"`
	ClickAction string `json:"click_action"`
	ClickURL    string `json:"click_url"`
}
