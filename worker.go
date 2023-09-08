package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/google/uuid"
)

func work(respChan chan<- string) {
	conn, _, _, err := ws.DefaultDialer.Dial(context.Background(), "ws://localhost:3000")
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	id := uuid.New().String()

	for {
		pingMessage := PingMessage{
			WSMessage:      WSMessage{EventType: "ping"},
			UserExternalId: id,
		}

		messageJson, err := json.Marshal(pingMessage)

		err = wsutil.WriteClientMessage(conn, ws.OpText, messageJson)
		if err != nil {
			panic(err)
		}

		var pongMessage *WSMessage = nil
		for pongMessage == nil || pongMessage.EventType != "pong" {
			resp, err := wsutil.ReadServerText(conn)
			if err != nil {
				panic(err)
			}

			if err = json.Unmarshal(resp, &pongMessage); err != nil {
				continue
			}
		}

		t := time.Now().UnixMicro()
		respChan <- fmt.Sprintf("%s -- %s", fmt.Sprint(t), id)

		time.Sleep(20 * time.Millisecond)
	}
}
