package reader

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type SshConfigs struct {
	Config map[string]SshConfig
}

type SshConfig struct {
	Hostname string `yaml:"hostname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
	Key      string `yaml:"key"`
}

func ReadYaml() (SshConfigs, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln(err)
	}
	configDir := fmt.Sprintf("%v/.hermes", home)
	files, err := ioutil.ReadDir(configDir)
	if err != nil {
		return SshConfigs{}, err
	}
	configs := map[string]SshConfig{}
	for _, f := range files {
		if !strings.Contains(f.Name(), "yaml") && !strings.Contains(f.Name(), "yml") {
			continue
		}
		yamlFile, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", configDir, f.Name()))
		if err != nil {
			return SshConfigs{}, err
		}
		var config SshConfig
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			return SshConfigs{}, err
		}
		configs[config.Hostname] = config
	}
	if len(configs) == 0 {
		return SshConfigs{}, errors.New("no configuration yaml found")
	}
	return SshConfigs{Config: configs}, nil
}
