package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "ChatBotCLI",
	Short: "ChatBotCLI可以帮助你快速的创建云湖ChatGPT机器人",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version: 1.0.1")
		fmt.Println(cmd.Help())
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
