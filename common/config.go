package common

import (
	yaml "gopkg.in/yaml.v2"

	"os"
)

const (
	pathToConfigFile = "config.yaml"
)

// Config структура файла config.yaml
type Config struct {
	Server struct {
		Listen string `yaml:"listen"`
		Port   int    `yaml:"port"`
		Debug  bool   `yaml:"debug"`
	} `yaml:"server"`

	MusicDir string `yaml:"musicDir"`
}

// NewConfig создает новый объект структуры Config
func NewConfig() (*Config, error) {
	var cfg = new(Config)

	file, err := os.Open(pathToConfigFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
