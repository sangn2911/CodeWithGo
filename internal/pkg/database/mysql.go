package database

import (
	"fmt"
	"go-research/internal/pkg/config"
	"go-research/internal/pkg/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideMySQLConfig(cfg config.SqlConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("Err: %v\n", err)
		return nil, err
	}
	util.Logger.Infof(fmt.Sprintf("connect db with host=%s user=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.DbName, cfg.Port))
	return db, nil
}
