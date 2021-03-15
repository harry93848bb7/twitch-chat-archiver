package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/harry93848bb7/chat-archiver/messages"
)

func main() {
	archive, err := messages.ArchiveChat("", "")
	if err != nil {
		panic(err)
	}
	b, err := json.Marshal(&archive)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("archive.json", b, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
