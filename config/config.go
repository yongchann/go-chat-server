package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

var config Config

func init() {
	path, err := filepath.Abs("./config/config.yml")
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(file, &config); err != nil {
		panic(err)
	}
	log.Println("config loaded")
}

type Config struct {
	Redis struct {
		Address  string `yaml:"address"`
		Passwrod string `yaml:"password"`
	} `yaml:"redis"`
}

func Get() Config {
	return config
}
