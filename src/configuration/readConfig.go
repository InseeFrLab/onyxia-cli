package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/kelseyhightower/envconfig"
)

var Configuration Config

func ReadConfig() {
	readFile(&Configuration)
	readEnv(&Configuration)
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config) {
	f, err := os.Open("config.yml")
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
}
