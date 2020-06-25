package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Webhook struct {
		Enabled bool   `yaml:"enabled"`
		Url     string `yaml:"url"`
	}

	Slack struct {
		Enabled bool   `yaml:"enabled"`
		Url     string `yaml:"url"`
	}

	Telegram struct {
		Enabled bool   `yaml:"enabled"`
		Token   string `yaml:"token"`
		ChatID  string `yaml:"chatID"`
	}

	Workplace struct {
		Enabled   bool   `yaml:"enabled"`
		ThreadKey string `yaml:"thread.key"`
		Token     string `yaml:"token"`
	}

	ExcludedReasons []string `yaml:"excluded.reasons,flow"`
	IncludedReasons []string `yaml:"included.reasons,flow"`

	IncludedNamespace []string `yaml:"included.namespaces,flow"`
}

var GlobalConfig = &Config{}

func (config *Config) Load() (err error) {
	yamlFile, err := ioutil.ReadFile("/usr/src/app/etc/siera-kube-watch/config.yaml")
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
