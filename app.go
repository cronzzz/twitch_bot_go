package main

import (
	"fmt"
	"time"
	"twitch_bot_go/process"
)

/**
@todo lock for threads
*/

func main() {

	dChanProc := make(chan int)
	dChanQuit := make(chan int)

	go func() {
		for {
			dChanProc <- 1
		}
	}()

	dispatcher := process.Dispatcher{}
	dispatcher.SetProcessChannel(dChanProc)
	dispatcher.SetQuitChannel(dChanQuit)

	for i := 0; i < 10; i++ {
		batch := process.Batch{}
		batch.SetId(i)
		batch.Add(func() {
			time.Sleep(1 * time.Second)
		})
		fmt.Println(fmt.Sprintf("added %d", i))
		dispatcher.Add(&batch)
	}

	go func() {
		dispatcher.Dispatch()
	}()
	time.Sleep(7 * time.Second)
	dChanQuit <- 1

	//client := twitch.NewClient("cronzzz", "oauth:rnodyn44o60hm77dntwsu4ik57114p")
	//client.Join("klavdius")

	//client.OnPrivateMessage(func(message twitch.PrivateMessage) {
	//	fmt.Println(
	//		fmt.Sprintf(
	//			">>> %s (%s): %s",
	//			message.User.DisplayName,
	//			message.User.Name,
	//			message.Message,
	//		),
	//	)
	//
	//	var batch Batch
	//	batch.add(func() {
	//		regexpHandler := handler.Regexp{}
	//		regexpHandler.SetClient(client)
	//		regexpHandler.ProcessMessage(&message)
	//	})
	//	dispatcher.add(&batch)
	//})
	//
	//err := client.Connect()
	//if err != nil {
	//	panic(err)
	//}
}
