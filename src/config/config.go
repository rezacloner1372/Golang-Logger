package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	...
	Logger   LoggerConfig
	...
}

// some code here...
type LoggerConfig struct {
	FilePath string
	Encoding string
	Level    string
	Logger   string
}
// some code here...
