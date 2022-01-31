package config

import (
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ScoreToWin       int `toml:"ScoreToWin"`
	TimePerFrameInMs int

	Board *BoardConfig `toml:"board"`
	Snake *SnakeConfig `toml:"snake"`
	Food  *FoodConfig  `toml:"food"`
}

// NewConfig - parses the user-provided config file
// if not possible returns a default config
func NewConfig(filePath string) *Config {

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return &defaultConfig
	}

	config := &Config{}
	_, err = toml.Decode(string(b), config)
	if err != nil {
		log.Fatal("Could not parse the config file")
		return &defaultConfig
	}

	return config

}
