package config

import (
	"errors"
	"os"
)

var ErrUnknownEnvironment = errors.New("unknown environment")

const (
	OSEnvEnvironment = "ENVIRONMENT"
)

//go:generate enumer -type=Environment -json -text -yaml -sql -transform=snake
type Environment int

const (
	Production Environment = iota
	Develop
	Local
)

func Path(env Environment, configPaths map[Environment]string) (string, error) {
	path, exists := configPaths[env]
	if !exists {
		return "", ErrUnknownEnvironment
	}

	return path, nil
}

func CurrentEnvironment() (Environment, error) {
	osEnv := os.Getenv("ENVIRONMENT")
	if osEnv == "" {
		return Production, nil
	}

	return EnvironmentString(osEnv)
}
