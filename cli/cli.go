package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"hermes/bootstrap"
	"hermes/executor"
	"hermes/git"
	"hermes/reader"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "hermes",
	Short: "Hermes is a SSH Config Manager",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var execCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("require hostname for connecting to server")
		}
		bootstrapper := bootstrap.Bootstrap{}
		homeDir := bootstrapper.CheckOrInitDirectory()
		if err := bootstrapper.CheckEmptyDir(homeDir); err != nil {
			log.Fatalln("please use `init` command before using this cli app")
		}
		configs, err := reader.ReadYaml()
		if err != nil {
			log.Fatalln(err)
		}
		exec := executor.Executor{Configs: configs}
		if err := exec.Execute(args[0]); err != nil{
		    log.Fatalln(err)
        }
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hermes client library",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("require private gituri repository for getting inventories")
		}
		bootstrapper := bootstrap.Bootstrap{}
		clone := git.Git{Uri: args[0]}
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

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Remove hermes client inventory",
	Run: func(cmd *cobra.Command, args []string) {
		bootstrapper := bootstrap.Bootstrap{}
		if err := bootstrapper.Destroy(); err != nil {
			log.Fatalln(err)
		}
		log.Println("destroy hermes config successfully")
	},
}

func Execute() {
	rootCmd.AddCommand(initCmd, execCmd, destroyCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
