package cli

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"hermes/bootstrap"
	"hermes/git"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Hermes is a SSH Config Manager",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapper := bootstrap.Bootstrap{}
		homeDir := bootstrapper.CheckOrInitDirectory()
		if err := bootstrapper.CheckEmptyDir(homeDir); err != nil {
			log.Fatalln("please use `init` command before using this cli app")
		}
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hermes client library",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapper := bootstrap.Bootstrap{}
		clone := git.Git{Uri: gitUri}
		if err := clone.Clone(); err != nil {
			log.Fatalln(fmt.Sprintf("failed to clone repository with details %v", err))
		}
		homeDir := bootstrapper.CheckOrInitDirectory()
		if err := bootstrapper.CheckEmptyDir(homeDir); err != nil {
			log.Fatalln(fmt.Sprintf("config file is not found, with error %v ", err))
		}
		log.Println("initialize hermes success")
	},
}

var (
	gitUri string
)

func Execute() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&gitUri, "gituri", "g", "", "SSH URI Git Inventory Key")
	flag.Parse()
	rootCmd.Execute()
}
