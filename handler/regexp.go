package handler

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"regexp"
	"strings"
	"twitch_bot_go/config"
)

type Regexp struct {
	client         *twitch.Client
	configurations *config.Config
	expressions    []*regexp.Regexp
}

func (handler *Regexp) Init(configurations *config.Config) {
	handler.configurations = configurations
	handler.expressions = []*regexp.Regexp{}
	for _, re := range handler.configurations.Blacklist {
		reCompiled, err := regexp.Compile(re)
		if err != nil {
			fmt.Print(err)
		} else {
			handler.expressions = append(handler.expressions, reCompiled)
		}
	}
}

func (handler *Regexp) SetClient(client *twitch.Client) {
	handler.client = client
}

func (handler *Regexp) ProcessMessage(message *twitch.PrivateMessage) {
	for _, re := range handler.expressions {
		if re.MatchString(strings.Trim(message.Message, " \t\n\r.")) {
			fmt.Println("------------------TO-BAN------------------")
			fmt.Println(fmt.Sprintf("%s: %s", message.User.DisplayName, message.Message))
			fmt.Println("------------------/TO-BAN------------------")
			handler.client.Say(message.Channel, fmt.Sprintf("/timeout %s %d", message.User.DisplayName, handler.configurations.General.TimeoutDuration))
		}
	}
}
