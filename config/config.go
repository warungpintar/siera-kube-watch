package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Webhook struct {
		Enabled bool   `yaml: enabled`
		Url     string `yaml: url`
	}

	Slack struct {
		Enabled bool   `yaml: enabled`
		Url     string `yaml: url`
	}
}

var GlobalConfig = &Config{}

func (config *Config) Load() (err error) {
	yamlFile, err := ioutil.ReadFile("/usr/src/app/etc/siera-kube/config.yaml")
	if err != nil {
		log.Fatalf("Error read config file: %v", err)
		return
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Error unmarshal: %v", err)
	}

	return
}
