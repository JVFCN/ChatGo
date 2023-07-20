package BotEvent

import (
	"ChatGPTBot/OpenAI"
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/subscription"
	"log"
)

func GroupJoin(event subscription.GroupJoinEvent) {
	err := SQLite.AddUser(event.UserId)
	if err != nil {
		log.Println(err)
		return
	}

	Response, err := Sends.SendTextMessage(event.ChatId, "group", "正在思考中...")
	MessageId := Response.Data.(map[string]interface{})["messageInfo"].(map[string]interface{})["msgId"].(string)

	UGid := Type.Id{
		MainId:   event.ChatId,
		MainType: "group",
		User:     event.UserId,
		Group:    event.ChatId,
	}

	err = OpenAI.GetGPTAnswer("有一位新成员进入了我们的群聊,请你随机用一种方式和语气欢迎新成员"+UGid.Name+"的到来", UGid, MessageId)
	if err != nil {
		log.Println(err)
		return
	}
}
