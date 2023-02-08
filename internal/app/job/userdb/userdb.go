package main

import (
	"errors"
	"fmt"
	"go-research/internal/pkg/config"
	"go-research/internal/pkg/database"
	"go-research/internal/pkg/util"
	"go-research/internal/user/domain"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var userModel = new(domain.User)

func main() {
	cfg_option, _ := config.GetArgs(os.Args)
	util.InitializeLogger()
	cfg := config.ProvideConfig(cfg_option)
	db, err := database.ProvideDBConfig(cfg.Sql)
	if err != err {
		fmt.Printf("cannot connect db: %v\n", err)
		return
	}
	if !(db.Migrator().HasTable(userModel)) {
		if err := db.AutoMigrate(userModel); err != nil {
			fmt.Printf("cannot migrate user model: %v\n", err)
			return
		} else {
			AddSeedData(db)
		}
	} else {
		AddSeedData(db)
	}

}

func AddSeedData(db *gorm.DB) {
	userList := []domain.User{
		{Id: uuid.New(), Name: "User", Fullname: "User1", Age: 35},
		{Id: uuid.New(), Name: "User", Fullname: "User2", Age: 35},
		{Id: uuid.New(), Name: "User", Fullname: "User3", Age: 35},
	}
	if err := db.First(userModel).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Model(userModel).Create(&userList).Error; err != nil {
			fmt.Printf("cannot add user data: %v\n", err)
		}
	}
}
