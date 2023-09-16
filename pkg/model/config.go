package model

import "github.com/gin-gonic/gin"

// APIServer holds the api server configuration
type APIServer struct {
	Addr      string       `yaml:"addr" validate:"required"`
	BasicAuth gin.Accounts `yaml:"basic_auth" validate:"required"`
}

// Log holds the log configuration
type Log struct {
	Level      string `yaml:"level"`
	FolderPath string `yaml:"folder_path"`
}

// Cfg is the main configuration structure for this application
type Cfg struct {
	APIServer  APIServer `yaml:"api_server" validate:"required"`
	Log        Log       `yaml:"log" validate:"required"`
	Production bool      `yaml:"production"`
}
