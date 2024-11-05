package config

import (
	"github.com/MichaelSitanggang/MiniProjectGo/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/karbon_db"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.User{}, &entities.Aktivitas{}, &entities.Input_aktivitas{})
	return db
}
