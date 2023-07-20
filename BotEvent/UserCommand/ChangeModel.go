package UserCommand

import (
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/openapi"
	"log"
)

func ChangeModel(Type string, UGid Type.Id) {
	Buttons := []openapi.Button{
		{
			Text:       "GPT-4",
			ActionType: 3,
			Value:      "gpt-4",
		},
		{
			Text:       "GPT-4-32k",
			ActionType: 3,
			Value:      "gpt-4-32k",
		},
		{
			Text:       "GPT-3-turbo",
			ActionType: 3,
			Value:      "gpt-3.5-turbo",
		},
		{
			Text:       "GPT-3-turbo-16k",
			ActionType: 3,
			Value:      "gpt-3.5-turbo-16k",
		},
	}
	_, err := Sends.SendTextMessage1(UGid.MainId, Type, "请选择模型", Buttons)
	if err != nil {
		log.Println(err)
	}
}
