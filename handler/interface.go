package handler

import (
	"github.com/gempir/go-twitch-irc/v3"
	"twitch_bot_go/config"
)

type IHandler interface {
	Init(configurations *config.Config)
	SetClient(client *twitch.Client)
	ProcessMessage(message *twitch.PrivateMessage)
}

var handlers []IHandler

func InitializeFilters(client *twitch.Client) {
	handlers = []IHandler{}

	//regexp
	regexpHandler := &Regexp{}
	regexpHandler.Init(config.GetConfig())
	regexpHandler.SetClient(client)

	handlers = append(handlers, regexpHandler)
}

func Get() []IHandler {
	return handlers
}
