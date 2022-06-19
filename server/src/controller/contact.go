package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Contact() {
	bot, err := linebot.New(
		"7abc6cda334f08891c17ab82b090b38e",
		"VeyFwt/tWBR2fHge2+Q7Cvq+0f45E5UIFBP6NkMJUMTyDuGNrWEZtQUey8b9xH7ftIQ/7y02CcsuhfPal1DmPbdpAOWoOyG/hs7F7EQeeHrtLOlx9iOgbHhJPZDUiRDl4Q2Zb8JJiY9SzwSjYQplKAdB04t89/1O/w1cDnyilFU=",
	)
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf(
						"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":", nil); err != nil {
		log.Fatal(err)
	}

}
