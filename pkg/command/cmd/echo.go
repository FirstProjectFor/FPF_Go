package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	initEchoFlag()
	rootCmd.AddCommand(echoCmd)
}

func initEchoFlag() {
	flags := echoCmd.PersistentFlags()
	flags.StringVarP(&echo, "echo", "e", "ping", "linux echo")
	err := echoCmd.MarkPersistentFlagRequired("echo")
	if err != nil {

	}
}

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echo short description",
	Long:  "echo long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("echo run, echo: ", echo)
	},
}
