package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const CONFIG_FILE_NAME = "config.toml"

type (
	Config struct {
		Host             string
		Port             int
		PostgresHost     string
		PostgresPort     int
		PostgresUser     string
		PostgresPassword string
		PostgresDB       string
	}
)

func Configuration() (Config, error) {
	if _, err := os.Stat(CONFIG_FILE_NAME); err != nil {
		return Config{}, err
	}

	var conf Config
	_, err := toml.DecodeFile(CONFIG_FILE_NAME, conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return Config{}, err
	}

	return conf, nil
}
