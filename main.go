package main

import "fmt"

type WSMessage struct {
	EventType string `json:"eventType"`
}

type PingMessage struct {
	WSMessage
	UserExternalId string `json:"userExternalId"`
}

func main() {
	respChan := make(chan string)

	for i := 0; i < 250; i++ {
		go work(respChan)
	}

	for {
		select {
		case resp := <-respChan:
			fmt.Println(resp)
		}
	}
}
