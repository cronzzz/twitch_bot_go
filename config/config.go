package config

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
)

var configFileName string

type Config struct {
	General   general
	Blacklist []string
}

type general struct {
	Channel         string
	Token           string
	Username        string
	TimeoutDuration int
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		instance = &Config{}
	}

	return instance
}

func (c *Config) Init() {
	configFileName = "config.json"
	if _, err := os.Stat(configFileName); errors.Is(err, os.ErrNotExist) {
		c.Save()
	} else {
		c.Load()
	}
}

func (c *Config) Load() {
	content, err := ioutil.ReadFile(configFileName)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, c)

	if err != nil {
		panic(err)
	}
}

func (c *Config) Save() {
	content, _ := json.Marshal(c)
	_ = ioutil.WriteFile(configFileName, content, fs.ModePerm)
}
