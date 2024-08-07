package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Service struct {
	Name          string `yaml:"name"`
	ListenPort    int    `yaml:"listen_port"`
	ContainerName string `yaml:"container_name"`
	ContainerPort int    `yaml:"container_port"`
}

type Config struct {
	Services []Service `yaml:"services"`
}

func ParseYAML(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config, nil
}
