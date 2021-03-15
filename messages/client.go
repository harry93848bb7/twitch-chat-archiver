package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client ...
type Client struct {
	clientID string
	Client   *http.Client
}

// NewClient ...
func NewClient(clientID string) *Client {
	return &Client{
		clientID: clientID,
		Client:   http.DefaultClient,
	}
}

// GetVODInfo ..
func (c *Client) GetVODInfo(vodID string) (*VODInfo, error) {
	request, err := http.NewRequest(http.MethodGet, "https://api.twitch.tv/kraken/videos/"+vodID, nil)
	if err != nil {
		return &VODInfo{}, err
	}
	request.Header.Add("Client-ID", c.clientID)
	request.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return &VODInfo{}, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &VODInfo{}, err
	}
	var data VODInfo
	if err := json.Unmarshal(b, &data); err != nil {
		return &VODInfo{}, err
	}
	return &data, nil
}

// GetMessageChunk ...
func (c *Client) GetMessageChunk(next string, vodID string) (*MessageChunk, error) {
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.twitch.tv/v5/videos/%s/comments?cursor="+next, vodID), nil)
	if err != nil {
		return &MessageChunk{}, nil
	}
	r.Header.Add("Client-ID", c.clientID)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return &MessageChunk{}, nil
	}
	var data MessageChunk
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &MessageChunk{}, nil
	}
	if err := json.Unmarshal(b, &data); err != nil {
		return &MessageChunk{}, nil
	}
	return &data, nil
}
