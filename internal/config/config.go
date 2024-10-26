package config

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var once sync.Once

type (
	AppInfo struct {
		Name      string `mapstructure:"name" validate:"required"`
		Version   string `mapstructure:"version" validate:"required"`
		Env       string `mapstructure:"environtment" validate:"required"`
		SecretKey string `mapstructure:"secretkey" validate:"required"`
	}
	Server struct {
		Port         int           `mapstructure:"port" validate:"required"`
		AllowOrigins []string      `mapstructure:"allowOrigins" validate:"required"`
		Timeout      time.Duration `mapstructure:"timeout" validate:"required"`
	}
	Database struct {
		TestDBPath    string `mapstructure:"test_db_path" validate:"required"`
		ProdDBPath    string `mapstructure:"prod_db_path" validate:"required"`
		DevDBPath     string `mapstructure:"dev_db_path" validate:"required"`
		CurrentDBPath string `mapstructure:"-"` // Used to store the active database path based on environment

	}
	Config struct {
		Server   *Server   `mapstructure:"server" validate:"required"`
		AppInfo  *AppInfo  `mapstructure:"appinfo" validate:"required"`
		Database *Database `mapstructure:"database" validate:"required"`
	}
)

var (
	configInstance *Config
	configError    error
)

func loadConfig() (*Config, error) {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			configError = fmt.Errorf("error reading config file: %w", err)
			return
		}

		// Unmarshal config into configInstance
		if err := viper.Unmarshal(&configInstance); err != nil {
			configError = fmt.Errorf("unable to decode into struct: %w", err)
			return
		}

		// Validate config values
		validate := validator.New()
		if err := validate.Struct(configInstance); err != nil {
			configError = err
			return
		}

		// Set CurrentDBPath based on environment
		switch configInstance.AppInfo.Env {
		case "production":
			configInstance.Database.CurrentDBPath = configInstance.Database.ProdDBPath
		case "test":
			configInstance.Database.CurrentDBPath = configInstance.Database.TestDBPath
		case "development":
			configInstance.Database.CurrentDBPath = configInstance.Database.DevDBPath
		default:
			configError = fmt.Errorf("invalid environment: %s", configInstance.AppInfo.Env)
		}
	})

	return configInstance, configError
}

func MustLoadConfig() *Config {
	cfg, err := loadConfig()
	if err != nil {
		panic(err)
	}

	return cfg
}
