package BotEvent

import (
	"ChatGPTBot/OpenAI"
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"fmt"
	"github.com/yhchat/bot-go-sdk/subscription"
	"log"
	"os"
)

func Command(event subscription.MessageEvent) {
	CommandName := event.Message.CommandName
	CommandId := event.Message.CommandId
	SenderId := event.Sender.SenderId
	ChatType := event.Chat.ChatType

	if ChatType == "bot" {
		ChatType = "user"
	} else {
		ChatType = "group"
	}

	UGid := Type.Id{}
	if ChatType == "group" {
		UGid = Type.Id{
			MainId:   event.Chat.ChatId,
			MainType: "group",
			User:     SenderId,
			Group:    event.Chat.ChatId,
			Name:     event.Sender.SenderNickname,
		}
	} else {
		UGid = Type.Id{
			MainId:   SenderId,
			MainType: "user",
			User:     SenderId,
			Group:    event.Chat.ChatId,
			Name:     event.Sender.SenderNickname,
		}
	}

	if CommandName == "设置私有ApiKey" || CommandId == 348 {
		fmt.Println(event.Message.Content["text"].(string))
		err := SQLite.UpdateUserApiKey(SenderId, event.Message.Content["text"].(string))
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的ApiKey已更新为:"+event.Message.Content["text"].(string))
	} else if CommandName == "查看ApiKey" || CommandId == 351 {
		Key, err := SQLite.GetUserApiKey(SenderId)
		if err != nil {
			log.Println(err)
			return
		}

		if Key == os.Getenv("DEFAULT_API") {
			_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的ApiKey为默认值")
		} else {
			_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的ApiKey为:"+Key)
		}
	} else if CommandName == "重置ApiKey" || CommandId == 371 {
		err := SQLite.UpdateUserApiKey(SenderId, "DefaultApiKey")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的ApiKey已重置为默认值")
	} else if CommandName == "AI生成图像" || CommandId == 352 {
		_, err := Sends.SendTextMessage(UGid.MainId, UGid.MainType, "正在生成图像...")
		if err != nil {
			log.Println(err)
			return
		}
		ImageUrl, err := OpenAI.CreateImage(UGid, event.Message.Content["text"].(string))
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendImageMessage(UGid.MainId, UGid.MainType, ImageUrl)
		if err != nil {
			log.Println(err)
			return
		}
	} else if CommandName == "添加期望功能/反馈Bug" || CommandId == 355 {
		_, err := Sends.SendTextMessage("3161064", "user", "用户"+UGid.Name+"("+UGid.User+")反馈了问题:\n"+event.Message.Content["text"].(string))
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "反馈成功, 管理员会尽快处理, 请等待回复")
	}
}
