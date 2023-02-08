package database

import (
	"fmt"
	"go-research/internal/pkg/config"
	"go-research/internal/pkg/util"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvidePostgreSQLConfig(cfg config.SqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.DbName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	util.Logger.Infof(fmt.Sprintf("connect db with host=%s user=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.DbName, cfg.Port))
	return db, nil
}
