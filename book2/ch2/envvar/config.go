package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func LoadConfig(path, envPrefix string, config any) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return errors.Wrap(err, "error loading config file from file")
		}
	}

	err := envconfig.Process(envPrefix, config)
	return errors.Wrap(err, "error loading config from environment")
}

func LoadFile(path string, config any) error {
	configFile, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "failed to open config file")
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}
