package config

import (
	"fmt"
	"io/fs"

	"github.com/aliy-turkois/cleanenv"
)

func ReadFS[T any](fs fs.FS, path string) (*T, error) {
	var cfg T
	if err := cleanenv.ReadConfigFS(fs, path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

type ReaderFunc[T any] func() (cfg *T, source string, err error)

func Reader[T any](
	configPathsFS fs.FS,
	configPaths map[Environment]string,
) ReaderFunc[T] {
	return func() (
		cfg *T,
		source string,
		err error,
	) {
		env, err := CurrentEnvironment()
		if err != nil {
			return nil, "", err
		}

		path, err := Path(env, configPaths)
		if err != nil {
			return nil, "", err
		}

		cfg, err = ReadFS[T](configPathsFS, path)
		if err != nil {
			return
		}

		source = fmt.Sprintf("internal config '%s' specified by environment '%s'", path, env)
		return
	}
}
