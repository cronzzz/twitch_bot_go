package handler

import "github.com/gempir/go-twitch-irc/v3"

type IHandler interface {
	processMessage(message twitch.PrivateMessage)
	setClient(client twitch.Client)
}
