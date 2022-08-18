package main

import (
	"flag"
	"fmt"
	"github.com/gempir/go-twitch-irc/v3"
	"strings"
	"twitch_bot_go/command"
	"twitch_bot_go/config"
	"twitch_bot_go/handler"
)

func catchParams() {
	channelFlag := flag.String("channel", "", "Channel name")
	usernameFlag := flag.String("username", "", "Bot Username")
	tokenFlag := flag.String("token", "", "Oauth token")
	timeoutDurationFlag := flag.Int("timeout", 600, "Timeout duration")
	flag.Parse()
	configsChanged := false
	if *channelFlag != "" {
		config.GetConfig().General.Channel = *channelFlag
		configsChanged = true
	}

	if *usernameFlag != "" {
		config.GetConfig().General.Username = *usernameFlag
		configsChanged = true
	}

	if *tokenFlag != "" {
		config.GetConfig().General.Token = *tokenFlag
		configsChanged = true
	}

	if *timeoutDurationFlag != 600 {
		config.GetConfig().General.TimeoutDuration = *timeoutDurationFlag
		configsChanged = true
	}

	if configsChanged {
		config.GetConfig().Save()
	}
}

func initializeClient() *twitch.Client {
	client := twitch.NewClient(config.GetConfig().General.Username, fmt.Sprintf("oauth:%s", config.GetConfig().General.Token))
	client.Join(config.GetConfig().General.Channel)
	return client
}

func setConnectHandler(client *twitch.Client) {
	client.OnConnect(func() {
		fmt.Println("--------CONNECTED--------")
	})
}

func setPrivMsgHandler(client *twitch.Client) {
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		fmt.Println(
			fmt.Sprintf(
				">>> %s (%s) @ %s: %s",
				message.User.DisplayName,
				message.User.Name,
				message.Channel,
				message.Message,
			),
		)

		for _, h := range handler.Get() {
			go h.ProcessMessage(&message)
		}
	})
}

func setDirectMessageHandler(client *twitch.Client) {
	client.OnWhisperMessage(func(message twitch.WhisperMessage) {
		fmt.Println(
			fmt.Sprintf(
				"WHISPER >>> %s (%s): %s",
				message.User.DisplayName,
				message.User.Name,
				message.Message,
			),
		)

		if message.User.Name != "klavdius" {
			client.Whisper(message.User.Name, "beep-bop")
			return
		}

		if strings.HasPrefix(message.Message, "list") ||
			strings.HasPrefix(message.Message, "add") ||
			strings.HasPrefix(message.Message, "rm") {
			command.Execute(message.Message, client)
		}
	})
}

func setReconnectHandler(client *twitch.Client) {
	client.OnReconnectMessage(func(message twitch.ReconnectMessage) {
		fmt.Println("--------RECEIVED RECONNECT MESSAGE--------")
		fmt.Println(
			fmt.Sprintf("%s %s", message.RawType, message.Raw),
		)
		fmt.Println("--------/RECEIVED RECONNECT MESSAGE--------")
		fmt.Println("--------DISCONNECTING--------")
		err := client.Disconnect()
		if err != nil {
			panic(err)
		} else {
			fmt.Println("--------DISCONNECTING--------")
		}
		fmt.Println("--------CONNECTING--------")
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	})
}

func main() {
	config.GetConfig().Init()
	catchParams()
	client := initializeClient()
	setConnectHandler(client)
	setPrivMsgHandler(client)
	setDirectMessageHandler(client)
	setReconnectHandler(client)
	handler.InitializeFilters(client)
	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
