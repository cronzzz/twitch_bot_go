package command

import (
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"strconv"
	"strings"
	"twitch_bot_go/config"
)

func executeList(client *twitch.Client) {
	for i, re := range config.GetConfig().Blacklist {
		client.Say("klavdius", fmt.Sprintf("%d: %s", i, re))
	}
}

func executeAdd(command string, client *twitch.Client) {
	config.GetConfig().Blacklist = append(config.GetConfig().Blacklist, command[4:])
	config.GetConfig().Save()
	client.Say("klavdius", fmt.Sprintf("Added %s", command[4:]))
}

func executeRm(command string, client *twitch.Client) {
	idxToRemove, err := strconv.Atoi(command[3:])
	if err != nil {
		client.Say("klavdius", err.Error())
	}

	if len(config.GetConfig().Blacklist) < idxToRemove+1 {
		client.Say("klavdius", "Index out of range")
		return
	}

	reToRemove := config.GetConfig().Blacklist[idxToRemove]
	before := config.GetConfig().Blacklist[:idxToRemove]
	after := config.GetConfig().Blacklist[idxToRemove+1:]
	config.GetConfig().Blacklist = append(before, after...)
	config.GetConfig().Save()
	client.Say("klavdius", fmt.Sprintf("Removed %s", reToRemove))
}

func Execute(command string, client *twitch.Client) {
	client.Say("klavdius", fmt.Sprintf("Executing: %s", command))

	if strings.HasPrefix(command, "list") {
		executeList(client)
	}

	if strings.HasPrefix(command, "add") {
		executeAdd(command, client)
	}

	if strings.HasPrefix(command, "rm") {
		executeRm(command, client)
	}
}
