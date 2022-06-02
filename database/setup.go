package database

import (
	"gocompany/models"
	"log"

	// sqlite "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

type Parameters struct {
	ConnectionString       string
	ConnectionType         string
	ConnectionDatabaseFile string `default:"company.db"`
}

func ConnectDatabase(p Parameters) {
	Instance, err = gorm.Open(mysql.Open(p.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("database : failed to connect database!")
	}
}

func Migrate() {
	Instance.AutoMigrate(&models.Company{})
	Instance.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed")
}
