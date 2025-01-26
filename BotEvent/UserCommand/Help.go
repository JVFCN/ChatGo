package UserCommand

import (
	"ChatGPTBot/Sends"
	"ChatGPTBot/Type"
	"log"
)

const HelpContent = `1.输入.clear清空上下文(上下文保存三段对话)
2.输入.ChangeModel切换模型
3.输入.Model查看当前模型
4.输入.Pre查看会员到期时间
5.输入.free查看剩余免费次数

管理员命令:
1. !SetBoard
2. !SetModel
3  !SetPre
4. !post`

func Help(Type string, UGid Type.Id) {
	_, err := Sends.SendTextMessage(UGid.MainId, Type, HelpContent)
	if err != nil {
		log.Println(err)
	}
}
