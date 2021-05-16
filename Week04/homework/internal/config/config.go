package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server *Server `yaml:"server"`
	Data   *Data   `yaml:"data"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type Data struct {
	Database *Database `yaml:"database"`
}

type Database struct {
	Driver string `yaml:"driver"`
	Source string `yaml:"source"`
}

func Load(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
