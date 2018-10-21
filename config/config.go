package config

import "github.com/BurntSushi/toml"

type config struct {
	ListenAddr string            `toml:"listen_addr"`
	UserFiles  map[string]string `toml:"user_files"`
}

func LoadConfig(path string) (*config, error) {
	var c config
	defaultConfig(&c)

	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func defaultConfig(config *config) {
	config.ListenAddr = "0.0.0.0:8888"
}
