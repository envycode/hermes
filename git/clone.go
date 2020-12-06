package git

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

type Git struct {
	Uri string
}

func (g Git) Clone() error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)

	_ = syscall.Rmdir(configDir)

	cmd := exec.Command("bash", "-c", fmt.Sprintf("git clone %v %v", g.Uri, configDir))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (g Git) Update() error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)

	_ = syscall.Rmdir(configDir)

	cmd := exec.Command("bash", "-c", "git reset --hard HEAD")
	cmd.Dir = configDir
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmdPull := exec.Command("bash", "-c", "git pull")
	cmdPull.Dir = configDir
	cmdPull.Stderr = os.Stderr
	cmdPull.Stdin = os.Stdin
	cmdPull.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return cmdPull.Run()
}
