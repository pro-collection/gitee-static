package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	AccessToken string `yaml:"access_token"`
	Owner       string `yaml:"owner"`
	Repository  string `yaml:"repository"`
	Path        string `yaml:"path"`
}

var config *Config

func init() {
	// 加载配置
	err := load("src/config/config.yaml")
	if err != nil {
		fmt.Println("Failed to load configuration")
		return
	}
}

func load(path string) error {
	result, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
