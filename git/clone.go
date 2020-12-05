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

	if err := syscall.Rmdir(configDir); err != nil {
		return err
	}

	cmd := exec.Command("bash", "-c", fmt.Sprintf("git clone %v %v", g.Uri, configDir))
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
