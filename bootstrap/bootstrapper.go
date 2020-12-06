package bootstrap

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Bootstrap struct{}

func (b Bootstrap) CheckOrInitDirectory() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)
	_, err = os.Stat(configDir)
	if os.IsNotExist(err) {
		err := os.Mkdir(configDir, os.FileMode(755))
		if err != nil {
			log.Fatalln("couldn't get permission for creating "+configDir+" err: ", err)
		}
	}
	return configDir
}

func (b Bootstrap) CheckEmptyDir(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Readdir(1)
	if err == io.EOF {
		return err
	}
	return nil
}

func (b Bootstrap) Destroy() error {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)
	return os.RemoveAll(configDir)
}
