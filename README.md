# ChatGo 运行于[云湖](https://yhchat.com)上的智能助手
该项目可实现大部分ChatGPT的操作(上下文, 清除记忆, 快捷复制等)  
部署起来也十分方便, `data/.env`文件示例:  
```
TOKEN=云湖机器人Token  
DEFAULT_API=默认API(可使用chatbot --base命令切换API代理)  
DEFAULT_MODEL=机器人默认使用模型  
PROXY=网络代理地址(如http://127.0.0.1:7890)
```

使用该项目, 您可以**自由, 快捷, 方便**的创建属于您自己的云湖聊天机器人  
而不必再去编写重复的代码!  

> 对于机器人的vip功能(如不需要, 可将该部分代码删除), 您可以在部署后使用.help命令查看  
或直接查看`BotEvent/OnMessageNormal.go`中的`RunCommand`函数  

## 总之
该项目可以帮助您快速部署任意支持OpenAI API格式的智能助手, 希望可以帮到您!
