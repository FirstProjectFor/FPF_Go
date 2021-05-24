package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	versionCmd.AddCommand(version1)
	versionCmd.AddCommand(version2)
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show linux version short description",
	Long:  "show linux version long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version run, version: ", version)
	},
}

var version1 = &cobra.Command{
	Use:   "version1",
	Short: "show linux version1 short description",
	Long:  "show linux version1 long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version1 run, version: ", version)
	},
}

var version2 = &cobra.Command{
	Use:   "version2",
	Short: "show linux version2 short description",
	Long:  "show linux version2 long description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version2 run, version: ", version)
	},
}
