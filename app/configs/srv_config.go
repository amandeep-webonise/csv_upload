package configs

import (
	"github.com/webonise/csv_upload/pkg/envprovider"
)

// ServerConfig contains configurations for server
type ServerConfig struct {
	Environment    string `env:"ENV,required"`
	Port           string `env:"PORT,required"`
	DBUsername     string `env:"DB_USERNAME,required"`
	DBPassword     string `env:"DB_PASSWORD,required"`
	DBHost         string `env:"DB_HOST,required"`
	DBName         string `env:"DB_NAME,required"`
	DBConnParams   string `env:"DB_PARAMS,required"`
	DBPORT         string `env:"DB_PORT,required"`
	AirbrakeServer string `env:"AIRBRAKE_ENDPOINT"`
	AirbrakeAPIKey string `env:"AIRBRAKE_API_KEY"`
}

// InitialiseServerCfg initialise server configurations
func (cfg *ServerConfig) InitialiseServerCfg(envp *envprovider.EnvConfigProvider) {
	if err := envp.Parse(cfg); err != nil {
		panic(err)
	}
}
