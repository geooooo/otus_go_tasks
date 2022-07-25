package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type rawConfig struct {
	HttpListen string `yaml:"http_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`
}

func (rConfig *rawConfig) toConfig() (*Config, error) {
	config := Config{}

	config.PathToLogFile = rConfig.LogFile

	logLevel, error := LogLevelFromString(rConfig.LogLevel)
	if error != nil {
		return nil, error
	}
	config.LogLevel = logLevel

	ip, port, error := getIpAndPort(rConfig.HttpListen)
	if error != nil {
		return nil, error
	}
	config.Ip = ip
	config.Port = port

	return &config, nil
}

func getIpAndPort(value string) (string, int, error) {
	values := strings.Split(value, ":")
	if len(values) != 2 {
		return "", 0, fmt.Errorf("invalid ip and port: '%s'", value)
	}

	ip := values[0]

	port, error := strconv.Atoi(values[1])
	if error != nil {
		return "", 0, fmt.Errorf("invalid port: '%s'", value)
	}

	return ip, port, nil
}

type Config struct {
	Ip            string
	Port          int
	PathToLogFile string
	LogLevel      LogLevel
}

var defaultConfig = Config{
	Ip:            "127.0.0.1",
	Port:          80,
	PathToLogFile: "log",
	LogLevel:      Debug,
}

func InitConfig(pathToConfig string) (*Config, error) {
	if pathToConfig == "" {
		return &defaultConfig, nil
	}

	rawConfig, error := readConfig(pathToConfig)
	if error != nil {
		return nil, error
	}

	config, error := rawConfig.toConfig()
	if error != nil {
		return nil, error
	}

	return config, nil
}

func readConfig(pathToConfig string) (*rawConfig, error) {
	configData, error := ioutil.ReadFile(pathToConfig)
	if error != nil {
		return nil, fmt.Errorf("cannot open config file: '%s'", pathToConfig)
	}

	var rConfig rawConfig
	if error := yaml.Unmarshal(configData, &rConfig); error != nil {
		return nil, errors.New("cannot read config file")
	}

	return &rConfig, nil
}
