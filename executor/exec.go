package executor

import (
	"fmt"
	"hermes/reader"
	"log"
	"os"
	"os/exec"
)

type Executor struct {
	Configs reader.SshConfigs
}

func (e Executor) Execute(hostname string) {
	config, found := e.Configs.Config[hostname]
	if !found {
		log.Fatalln(fmt.Sprintf("no configuration found for host: %v", hostname))
	}
	ssh := exec.Command("bash", "-c", fmt.Sprintf("sshpass -p '%v' ssh %v@%v -p %v", config.Password, config.User, config.Hostname, config.Port))
	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	ssh.Stdin = os.Stdin

	if err := ssh.Run(); err != nil {
		log.Fatalln(err)
	}
}
