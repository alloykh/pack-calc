package main

import (
	"embed"
	"fmt"
	"github.com/alloykh/pack-calc/internal/packs"
	"github.com/alloykh/pack-calc/pkg/config"
	"strconv"

	"github.com/alloykh/pack-calc/pkg/log"
)

var (
	//go:embed environments
	configPathsFS embed.FS

	configPaths = map[config.Environment]string{
		config.Local:      "environments/local.yaml",
		config.Production: "environments/master.yaml",
		config.Develop:    "environments/develop.yaml",
	}

	readConfig = config.Reader[Config](configPathsFS, configPaths)
)

type Config struct {
	Server struct {
		Port            int `yaml:"port" env:"PORT" env-default:"4000" env-description:"HTTP port on which to run the server"`
		HealthCheckPort int `yaml:"health_check_port" env-default:"8080" env-description:"HTTP port on which to run the health check server"`
	} `yaml:"server"`

	Log struct {
		Verbosity log.Verbosity `yaml:"verbosity" env:"LOG_VERBOSITY" env-default:"0" env-description:"Logging verbosity level, 0-4"`
	} `yaml:"log"`

	Packs packs.Config `yaml:",inline"`
}

func (c Config) ServerAddr() string {
	return fmt.Sprintf(":%s", strconv.Itoa(c.Server.Port))
}

func (c Config) HealthCheckAddr() string {
	return fmt.Sprintf(":%s", strconv.Itoa(c.Server.HealthCheckPort))
}
