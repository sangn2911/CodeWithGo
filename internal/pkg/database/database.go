package database

import (
	"go-research/internal/pkg/config"

	"gorm.io/gorm"
)

func ProvideDBConfig(cfg config.SqlConfig) (*gorm.DB, error) {
	switch cfg.DBtype {
	case "postgreSQL":
		return ProvidePostgreSQLConfig(cfg)
	case "mySQL":
		return ProvideMySQLConfig(cfg)
	default:
		return ProvidePostgreSQLConfig(cfg)
	}
}
