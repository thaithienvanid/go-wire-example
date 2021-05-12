package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"go-wire-example/internal/services/rest_api"
)

// Config struct for webapp config
type Config struct {
	Environment string           `yaml:"environment"`
	RestAPI     *rest_api.Config `yaml:"rest_api"`
}

// LoadConfig ...
func LoadConfig(configPath string) (*Config, error) {
	config := &Config{}

	tmpl, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	tmpl = []byte(os.ExpandEnv(string(tmpl)))

	err = yaml.Unmarshal(tmpl, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// ValidateConfigPath ...
func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}

// ParseFlags ...
func ParseFlags() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()

	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}

	return configPath, nil
}

func ProvideConfig() *Config {
	configPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
		return nil
	}
	config, err := LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return config
}

func ProvideRestAPIConfig(config *Config) *rest_api.Config {
	return config.RestAPI
}
