package handler

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"regexp"
)

type Regexp struct {
	client *twitch.Client
}

func (handler *Regexp) SetClient(client *twitch.Client) {
	handler.client = client
}

func (handler *Regexp) ProcessMessage(message *twitch.PrivateMessage) {
	regexp, err := regexp.Compile("[\\@a-zA-Z\\d_\\-]*[\\_\\-\\.]\\d{1,2}")
	if err != nil {
		fmt.Print(err)
	}
	if regexp.MatchString(message.Message) {
		handler.client.Say(message.Channel, fmt.Sprintf("/timeout %s 5", message.User.DisplayName))
	}
}
