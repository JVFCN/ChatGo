package BotEvent

import (
	"ChatGPTBot/OpenAI"
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/subscription"
	"log"
	"strings"
)

func ButtonClicked(event subscription.ButtonReportInlineEvent) {
	Value := event.Value
	UGid := Type.Id{
		MainId:   event.RecvId,
		MainType: event.RecvType,
		User:     event.UserId,
		Group:    event.RecvType,
	}

	if strings.HasPrefix(Value, "buy") == true {
		_, err := Sends.SendTextMessage("3161064", "user", "用户"+UGid.User+" 昵称"+UGid.Name+"\n需要充值会员, 请及时处理")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "您的订单已提交, 请等待管理员处理")
		if err != nil {
			log.Println(err)
			return
		}
	} else if strings.HasPrefix(Value, "translate") == true {
		err := OpenAI.GetGPTAnswer(Value[9:]+"\n请把上面这个回答翻译成中文", UGid, event.MsgId)
		if err != nil {
			log.Println(err)
			return
		}
	} else if strings.HasPrefix(Value, "AgainReply") == true {
		_, err := Sends.EditTextMessage(event.MsgId, UGid.MainId, UGid.MainType, "正在重新生成回答...", nil)
		if err != nil {
			return
		}
		err = OpenAI.GetGPTAnswer(Value[10:], UGid, event.MsgId)
		if err != nil {
			log.Println(err)
			return
		}
	} else if Value == "gpt-3.5-turbo" {
		err := SQLite.UpdateUserModel(event.UserId, "gpt-3.5-turbo")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将模型切换为gpt-3.5-turbo")
		if err != nil {
			log.Println(err)
			return
		}
	} else if Value == "gpt-3.5-turbo-16k" {
		err := SQLite.UpdateUserModel(event.UserId, "gpt-3.5-turbo-16k")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将模型切换为gpt-3.5-turbo-16k")
		if err != nil {
			log.Println(err)
			return
		}
	} else if Value == "gpt-4" {
		err := SQLite.UpdateUserModel(event.UserId, "gpt-4")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将模型切换为gpt-4")
		if err != nil {
			log.Println(err)
			return
		}
	} else if Value == "gpt-4-32k" {
		err := SQLite.UpdateUserModel(event.UserId, "gpt-4-32k")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = Sends.SendTextMessage(UGid.MainId, UGid.MainType, "已将模型切换为gpt-4-32k")
		if err != nil {
			log.Println(err)
			return
		}
	}
}
