// Package slack implements sending and getting messages to and from Slack
package slack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/callerobertsson/gotools/mentalnote/config"
	"github.com/callerobertsson/gotools/mentalnote/model"
)

// APIBaseURL is the base URL for the Slack API
const APIBaseURL = "https://slack.com/api/"

// SendMessage sends a message to Slack
func SendMessage(message string, conf config.Config) {

	var u *url.URL
	u, _ = url.Parse(APIBaseURL)
	u.Path += "/chat.postMessage"

	params := url.Values{}
	params.Add("token", conf.APIToken)
	params.Add("channel", conf.ChannelID)
	params.Add("text", message)
	params.Add("username", conf.Username)
	params.Add("icon_url", conf.IconURL)

	u.RawQuery = params.Encode()

	client := &http.Client{}
	r, _ := http.NewRequest("GET", u.String(), nil)

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}

// GetMessages retrieves messages from Slack
func GetMessages(conf config.Config) ([]model.Message, error) {

	var u *url.URL
	u, _ = url.Parse(APIBaseURL)
	u.Path += "/groups.history"

	params := url.Values{}
	params.Add("token", conf.APIToken)
	params.Add("channel", conf.ChannelID)

	u.RawQuery = params.Encode()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", u.String(), nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	type MessagesResponse struct {
		Ok       bool            `json:"ok"`
		Error    string          `json:"error"`
		Messages []model.Message `json:"messages"`
	}

	var messageResponse MessagesResponse

	err = json.Unmarshal(body, &messageResponse)

	if err != nil {
		return nil, err
	}

	if messageResponse.Ok {
		return messageResponse.Messages, nil
	}

	return nil, errors.New(messageResponse.Error)
}
