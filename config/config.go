package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

//Main application configuration
type AppConfig struct {
	IsDevelopmentConfig bool `yaml:"is_dev"` //Development environment indicator - implemented stupidly but works
	// Http server configuration block. Has basic host and port params with default timeout values
	ServerConfig struct {
		Host    string   `yaml:"host"`
		Port    string   `yaml:"port"`
		Timeout struct { // Timeout struct, which includes default timeout values
			Read  int
			Write int
		}
	} `yaml:"server"`
	// Database configuration block - used for DSN construction
	DatabaseConfig struct {
		Host            string
		Port            string
		User            string // dev only username yaml field - in production it is read from env
		Password        string // dev only password yaml field - in production it is read from env
		Database        string
		ConnectionLimit int `yaml:"connection_limit"`
	} `yaml:"database"`
	// configuration for the email sender
	EmailConfig struct {
		Host         string
		Port         int
		User         string
		Password     string
		MandatoryTLS bool `yaml:"mandatory_tls"`
	} `yaml:"email"`
}

var ErrNoConfigFile = errors.New("config file does not exist")
var ErrFileIssue = errors.New("config file can not be opened")
var ErrParseIssue = errors.New("configuration file can not be parsed")

//Reads configuration from hardcoded file or from path specified by env variable
func ReadConfig() (*AppConfig, error) {
	path, _ := filepath.Abs("./config/config.dev.yml") // make absolute for portability
	if val, exists := os.LookupEnv("PROD_CONFIG"); exists {
		path = val // if env is specified - forget about hardcoded value and use one from env instead
	}

	_, err := os.Stat(path)
	if os.IsNotExist(err) { // check that file actually exists
		return nil, ErrNoConfigFile
	}

	config := &AppConfig{}

	file, err := os.Open(path)
	if err != nil {
		return nil, ErrFileIssue
	}

	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		return nil, fmt.Errorf("%s. Error: %s", ErrParseIssue.Error(), err.Error())
	}

	return config, nil
}
