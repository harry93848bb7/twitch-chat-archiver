package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Client ...
type Client struct {
	clientID string
}

// NewClient ...
func NewClient(clientID string) *Client {
	return &Client{
		clientID: clientID,
	}
}

// GetVODInfo ..
func (c *Client) GetVODInfo(vodID string) (*vodInfo, error) {
	request, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/videos/"+vodID, nil)
	if err != nil {
		return &vodInfo{}, err
	}
	request.Header.Add("Client-ID", c.clientID)
	request.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &vodInfo{}, err
	}
	if resp.StatusCode != 200 {
		log.Println("Failed to get VOD information from VODID", vodID)
		return &vodInfo{}, fmt.Errorf("Failed to get VOD information")
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &vodInfo{}, err
	}
	var data vodInfo
	if err := json.Unmarshal(b, &data); err != nil {
		return &vodInfo{}, err
	}
	return &data, nil
}

// GetMessageChunk ...
func (c *Client) GetMessageChunk(next string, vodID string) (*messageChunk, error) {
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.twitch.tv/v5/videos/%s/comments?cursor="+next, vodID), nil)
	if err != nil {
		return &messageChunk{}, nil
	}
	r.Header.Add("Client-ID", c.clientID)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return &messageChunk{}, nil
	}
	if resp.StatusCode != 200 {
		log.Println("Failed to get chat message chunk from VOD", vodID)
		return &messageChunk{}, fmt.Errorf("Failed to get chat message chunk")
	}
	var data messageChunk
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &messageChunk{}, nil
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return &messageChunk{}, nil
	}
	return &data, nil
}
