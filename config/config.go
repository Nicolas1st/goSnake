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

	KeyToActionMap map[rune]string
}

// NewConfig - parses the user-provided config file
// if not possible returns a default config
func NewConfig(filePath string) *Config {

	b, err := ioutil.ReadFile(filePath)

	// check file exists
	if err != nil {
		return &defaultConfig
	}

	// parsing main part of the config
	config := &Config{}
	_, err = toml.Decode(string(b), config)

	if err != nil {
		log.Fatal("Could not parse the config file")
		return &defaultConfig
	}

	// parsing key mappings
	keyboardKeys := &KeyboardKeys{}
	_, err = toml.Decode(string(b), keyboardKeys)

	if err != nil {
		log.Fatal("Could not parse the config file")
		return &defaultConfig
	}

	// contstructing final config
	keyToActionMap := keyboardKeys.toMap()
	config.KeyToActionMap = keyToActionMap

	return config

}
