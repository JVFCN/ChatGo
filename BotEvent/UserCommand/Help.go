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
4. !post

开发者云湖ID:3161064, 邮箱:j3280891657@gmail.com
关于付费版:本机器人GPT3.5, GPT3.5-16k模型完全免费使用 GPT4需要充值后使用
价格:10元/月,25元/季度,100元/年(无限制使用GPT4模型, 以及GPT4-32K模型)
注意:付费后请联系开发者,并发送付费截图,以便开通权限`

func Help(Type string, UGid Type.Id) {
	_, err := Sends.SendTextMessage(UGid.MainId, Type, HelpContent)
	if err != nil {
		log.Println(err)
	}
}
