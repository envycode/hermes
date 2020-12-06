package executor

import (
    "errors"
    "fmt"
    "hermes/reader"
    "io/ioutil"
    "log"
    "os"
    "os/exec"
)

type Executor struct {
	Configs reader.SshConfigs
}

func (e Executor) Execute(hostname string) error {
	config, found := e.Configs.Config[hostname]
	if !found {
	    return errors.New(fmt.Sprintf("no configuration found for host: %v", hostname))
	}

    home, err := os.UserHomeDir()
    if err != nil {
        return err
    }
    configDir := fmt.Sprintf("%v/.hermes", home)

    file, err := ioutil.TempFile(configDir, "file-temp-key-")
    if err != nil {
        return err
    }

    if err:= ioutil.WriteFile(file.Name(), []byte(config.Key), 777); err != nil{
         return err
    }

    defer func() {
        err := os.Remove(file.Name())
        if err != nil {
            log.Print(err)
        }
    }()

    ssh, err := e.generateSshCommand(config, file)
    if err != nil{
        return err
    }

	ssh.Stdout = os.Stdout
	ssh.Stderr = os.Stderr
	ssh.Stdin = os.Stdin

    return ssh.Run()
}

func (e Executor) generateSshCommand(config reader.SshConfig, file *os.File) (*exec.Cmd, error) {

    if config.Key != ""{
        return exec.Command("bash", "-c", fmt.Sprintf("ssh %v@%v -p %v -i %v", config.User, config.Hostname, config.Port, file.Name())) , nil
    }
    if config.Password != ""{
        return exec.Command("bash", "-c", fmt.Sprintf("sshpass -p '%v' ssh %v@%v -p %v", config.Password, config.User, config.Hostname, config.Port)), nil
    }

    return &exec.Cmd{}, errors.New("required key or password")

}
