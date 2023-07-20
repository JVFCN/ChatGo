package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"github.com/yhchat/bot-go-sdk/openapi"
)

func Post(Message string, Type string, UGid Type.Id) {
	IsAmin, err := SQLite.IsAmin(UGid.User)
	if err != nil {
		return
	}
	if IsAmin {
		Buttons := []openapi.Button{{Text: "复制公告", ActionType: 2, Value: Message}}

		Sends.BatchSendMessages(SQLite.GetAllUserIds(), Type, Message, Buttons)
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, Type, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
