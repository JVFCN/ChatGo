package cmd

import (
	"ChatGPTBot/BotEvent"
	"ChatGPTBot/SQLite"
	"github.com/spf13/cobra"
	"github.com/yhchat/bot-go-sdk/subscription"
)

var Port int

func init() {
	rootCmd.AddCommand(Server)

	rootCmd.PersistentFlags().IntVar(&Port, "port", 7888, "监听端口")
}

var Server = &cobra.Command{
	Use:   "server",
	Short: "Start the ChatBot server",
	Run: func(cmd *cobra.Command, args []string) {
		SQLite.Init()
		Subscription := subscription.NewSubscription(Port)
		Subscription.OnMessageNormal = BotEvent.Normal
		Subscription.OnMessageInstruction = BotEvent.Command
		Subscription.OnGroupJoin = BotEvent.GroupJoin
		Subscription.OnBotFollowed = BotEvent.Followed
		Subscription.OnGroupLeave = BotEvent.GroupLeave
		Subscription.OnButtonReportInline = BotEvent.ButtonClicked
		Subscription.Start()
	},
}
