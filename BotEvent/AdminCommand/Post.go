package AdminCommand

import (
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
)

func Post(Message string, Type string, UGid Type.Id) {
	IsAmin, err := SQLite.IsAmin(UGid.User)
	if err != nil {
		return
	}
	if IsAmin {
		Sends.BatchSendMessages(SQLite.GetAllUserIds(), Type, Message, nil)
	} else {
		_, err := Sends.SendTextMessage(UGid.MainId, Type, "无权执行本命令")
		if err != nil {
			return
		}
	}
}
