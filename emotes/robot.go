package emotes

import "embed"

//go:embed robot/*
var robotEmotes embed.FS

var emoteMapping = map[string]string{
	"robot1.png":  `:)`,
	"robot2.png":  `:(`,
	"robot3.png":  `:D`,
	"robot4.png":  `>(`,
	"robot5.png":  `:|`,
	"robot6.png":  `O_o`,
	"robot7.png":  `B)`,
	"robot8.png":  `:O`,
	"robot9.png":  `<3`,
	"robot10.png": `:/`,
	"robot11.png": `;)`,
	"robot12.png": `:P`,
	"robot13.png": `;P`,
	"robot14.png": `R)`,
}
