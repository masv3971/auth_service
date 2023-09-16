package configuration

import (
	"auth_service/pkg/helpers"
	"auth_service/pkg/logger"
	"auth_service/pkg/model"
	"errors"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type envVars struct {
	ConfigYAML string `envconfig:"AUTH_CONFIG_YAML" required:"true"`
}

// Parse parses config file from VC_CONFIG_YAML environment variable
func Parse(logger *logger.Log) (*model.Cfg, error) {
	logger.Info("Read environmental variable")
	env := envVars{}
	if err := envconfig.Process("", &env); err != nil {
		return nil, err
	}

	configPath := env.ConfigYAML

	cfg := &model.Cfg{}

	configFile, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(configPath)
	if err != nil {
		return nil, err
	}

	if fileInfo.IsDir() {
		return nil, errors.New("config is a folder")
	}

	if err := yaml.Unmarshal(configFile, cfg); err != nil {
		return nil, err
	}

	if err := helpers.Check(cfg, logger); err != nil {
		return nil, err
	}

	return cfg, nil
}
