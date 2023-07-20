package BotEvent

import (
	"ChatGPTBot/OpenAI"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/subscription"
	"log"
)

func GroupLeave(event subscription.GroupLeaveEvent) {
	Response, err := Sends.SendTextMessage(event.ChatId, "group", "正在思考中...")
	MessageId := Response.Data.(map[string]interface{})["messageInfo"].(map[string]interface{})["msgId"].(string)

	UGid := Type.Id{
		MainId:   event.ChatId,
		MainType: "group",
		User:     event.UserId,
		Group:    event.ChatId,
		Name:     event.Nickname,
	}

	err = OpenAI.GetGPTAnswer("有一位成员退出了我们的群聊,请你随机用一种方式和语气送别"+UGid.Name+"这位成员", UGid, MessageId)
	if err != nil {
		log.Println(err)
		return
	}
}
