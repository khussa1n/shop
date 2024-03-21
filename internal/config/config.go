package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DB DBConfig `yaml:"db"`
}

type DBConfig struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	DBName        string `yaml:"db_name"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	MigrationPath string `yaml:"migration_path"`
}

func InitConfig(path string) (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)

	return cfg, nil
}
