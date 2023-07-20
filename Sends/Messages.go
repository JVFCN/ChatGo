package Sends

import (
	"github.com/joho/godotenv"
	"github.com/yhchat/bot-go-sdk/openapi"
	"log"
	"os"
)

func SendTextMessage(recvId string, recvType string, text string) (openapi.BasicResponse, error) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
	}
	return openApi.SendTextMessage(textMessage)
}

func SendTextMessage1(recvId string, recvType string, text string, buttons interface{}) (openapi.BasicResponse, error) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))
	textMessage := openapi.TextMessage{
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	return openApi.SendTextMessage(textMessage)
}

func EditTextMessage(msgId string, recvId string, recvType string, text string, buttons interface{}) (openapi.BasicResponse, error) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))
	newTextMessage := openapi.EditTextMessage{
		MsgId:    msgId,
		RecvId:   recvId,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	return openApi.EditTextMessage(newTextMessage)
}

func SendImageMessage(recvId string, recvType string, Url string) (openapi.BasicResponse, error) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))

	content := map[string]interface{}{"imageUrl": Url}
	return openApi.SendMessage(recvId, recvType, "image", content)
}
func BatchSendMessages(recvIds []string, recvType string, text string, buttons interface{}) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))
	batchSendMessage := openapi.BatchTextMessage{
		RecvIds:  recvIds,
		RecvType: recvType,
		Text:     text,
		Buttons:  buttons,
	}
	openApi.BatchSendTextMessage(batchSendMessage)
}

func SetBoardAllUser(ContentType, Content string) {
	_ = godotenv.Load("data/.env")
	openApi := openapi.NewOpenApi(os.Getenv("TOKEN"))
	_, err := openApi.SetBotBoardAll(ContentType, Content)
	if err != nil {
		log.Println(err)
		return
	}
}
