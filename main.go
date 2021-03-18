package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	proto "github.com/golang/protobuf/proto"
	"github.com/harry93848bb7/chat-archiver/archiver"
)

func main() {
	r, err := archiver.Archive("934709752", "")
	if err != nil {
		panic(err)
	}
	b, err := proto.Marshal(r)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%d", r.VodId), b, os.ModePerm)
	if err != nil {
		panic(err)
	}
	b, err = json.Marshal(r)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(fmt.Sprintf("%d.json", r.VodId), b, os.ModePerm)
}
