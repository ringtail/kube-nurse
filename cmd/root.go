package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)


func init(){
	rootCmd.AddCommand(diagnoseCmd)
}

var rootCmd = &cobra.Command{
	Use:   "kube-nurse",
	Short: "kube-nurse is a kubernetes system diagnostic tool.",
	Long:  `You can use "kube-nurse diagnose [cluster-dump.log]" to check the status.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("kube-nurse is working!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
