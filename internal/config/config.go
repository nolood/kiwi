package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Env struct {
	Token string `env:"TELEGRAM_BOT_TOKEN"`
	Key   string `env:"MEILISEARCH_API_KEY"`
}

type Meilisearch struct {
	Key     string `yaml:"-"`
	Address string `yaml:"address"`
}

type Telegram struct {
	Token string
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type Config struct {
	Env         string      `yaml:"env" env-required:"true"`
	Meilisearch Meilisearch `yaml:"meilisearch"`
	Storage     Storage     `yaml:"storage" env-required:"true"`
	Telegram    Telegram    `yaml:"-"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading .env file: %v", err)
	}

	var envCfg Env
	if err = cleanenv.ReadEnv(&envCfg); err != nil {
		panic("failed to read env: " + err.Error())
	}

	cfg.Telegram.Token = envCfg.Token
	cfg.Meilisearch.Key = envCfg.Key

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
