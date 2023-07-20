package cmd

import (
	"ChatGPTBot/BotEvent"
	"ChatGPTBot/SQLite"
	"ChatGPTBot/Type"
	"github.com/spf13/cobra"
	"github.com/yhchat/bot-go-sdk/subscription"
)

func init() {
	rootCmd.AddCommand(Server)

	rootCmd.PersistentFlags().IntVar(&Type.Port, "port", 7888, "监听端口")
	rootCmd.PersistentFlags().StringVar(&Type.Base, "base", "https://api.mctools.online/v1", "ApiBase地址")
}

var Server = &cobra.Command{
	Use:   "server",
	Short: "Start the ChatBot server",
	Run: func(cmd *cobra.Command, args []string) {
		SQLite.Init()
		Subscription := subscription.NewSubscription(Type.Port)
		Subscription.OnMessageNormal = BotEvent.Normal
		Subscription.OnMessageInstruction = BotEvent.Command
		Subscription.OnGroupJoin = BotEvent.GroupJoin
		Subscription.OnBotFollowed = BotEvent.Followed
		Subscription.OnGroupLeave = BotEvent.GroupLeave
		Subscription.OnButtonReportInline = BotEvent.ButtonClicked
		Subscription.Start()
	},
}
