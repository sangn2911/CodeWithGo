package config

import (
	"fmt"
	"go-research/internal/pkg/util"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Sql SqlConfig `yaml:"dbconfig"`
}

type SqlConfig struct {
	DBtype   string `yaml:"dbtype"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func GetArgs(Args []string) (cfgType string, pkg string) {
	cfgType, pkg = "env", "mux"
	if len(Args) > 0 && len(Args) == 3 {
		fmt.Println("not enough arguments, use defaults")
		cfgType, pkg = os.Args[0], os.Args[1]
	}
	fmt.Printf("receive arguments: \n\t+ cfg_type: %v\n\t+ pkg: %v\n\n", cfgType, pkg)
	return cfgType, pkg
}

func ProvideConfig(cfg_option string) *Config {
	switch cfg_option {
	case "env":
		return provideConfigFromEnv()
	case "yml":
		return provideConfigFromYml()
	default:
		return provideConfigFromEnv()
	}
}

func provideConfigFromEnv() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		util.Logger.Warnf("Failed to load .env. Err: %s", err)
	}

	util.Logger.Info("loading configuration from environment")

	SqlConfig := SqlConfig{
		DBtype:   os.Getenv("DATABASE_TYPE"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		DbName:   os.Getenv("DATABASE_NAME"),
		User:     os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
	}

	util.Logger.Infof(
		"host=%s port=%s datebase=%s user=%s",
		SqlConfig.Host,
		SqlConfig.Port,
		SqlConfig.DbName,
		SqlConfig.User,
	)

	return &Config{
		Sql: SqlConfig,
	}
}

func provideConfigFromYml() *Config {
	var c Config
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		util.Logger.Fatalf("error read config.yml: %v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		util.Logger.Fatalf("error unmarshal config.yml: %v ", err)
	}
	return &c
}
