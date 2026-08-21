package config

import (
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/spf13/viper"
)

// AppConfig contains all app configuration
type AppConfig struct {
	App struct {
		// Name is the name of the application
		Name string `mapstructure:"name"`
		// ShutdownTimeout is the duration to wait before forcefully shutting down the application
		ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
		// Env is the environment where the application is running
		Env string `mapstructure:"env"`
		// RestAddr is the address for the REST server to listen on
		RestAddr string `mapstructure:"rest_addr"`
		// MetricsAddr is the address for the metrics server to listen on
		MetricsAddr string `mapstructure:"metrics_addr"`
	} `mapstructure:"app"`

	// Logger contains the logger configuration
	Logger struct {
		Level   string `mapstructure:"level"`
		AsJSON  bool   `mapstructure:"as_json"`
		NoColor bool   `mapstructure:"no_color"`
	} `mapstructure:"logger"`

	// CORS contains the CORS configuration
	CORS struct {
		// Enabled determines if CORS is enabled
		Enabled bool `mapstructure:"enabled"`
		// AllowedOrigins is a list of allowed origins
		AllowedOrigins []string `mapstructure:"allowed_origins"`
		// AllowedMethods is a list of allowed HTTP methods
		AllowedMethods []string `mapstructure:"allowed_methods"`
		// AllowedHeaders is a list of allowed headers
		AllowedHeaders []string `mapstructure:"allowed_headers"`
	} `mapstructure:"cors"`

	// Database contains the database connection information
	Database struct {
		// Connection is the connection string to connect to the database
		Connection string `mapstructure:"connection"`
	} `mapstructure:"database"`
}

// LoadConfig initialize the configuration and env vars from the config files
func LoadConfig(path string) (*AppConfig, error) {
	// export env vars from .env file
	err := godotenv.Load(filepath.Join(path, ".env"))
	if err != nil {
		log.Println("Cannot load .env file, skipping...")
	}

	vip := viper.New()

	// Read config file
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")
	vip.AddConfigPath(path)
	err = vip.ReadInConfig()
	if err != nil {
		log.Fatal("Cannot read config.yaml", err)
	}

	// Read all .env vars and set on the OS
	vip.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	vip.AutomaticEnv()

	// BindEnv Binds keys on AppConfig to the .env vars
	// (this allows setting the .env value which name matches with an AppConfig element)
	_ = vip.BindEnv("Database.Connection")

	// Load config and env vars on a AppConfig instance
	cfg := &AppConfig{}
	err = vip.Unmarshal(cfg)
	if err != nil {
		log.Fatal("Cannot load config file invalid config format")
	}

	return cfg, nil
}
