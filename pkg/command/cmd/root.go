package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func init() {
	cobra.OnInitialize(initFlag)
}

func initFlag() {
	rootCmd.PersistentFlags().StringVarP(&version, "version", "v", "1.0.0", "linux version")
}

var rootCmd = &cobra.Command{
	Use:   "linux",
	Short: "linux short description",
	Long:  `linux long description`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("linux command run, version is: ", version)
		fmt.Println("linux command run, args is: ", strings.Join(args, ","))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
