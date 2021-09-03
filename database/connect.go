package database

import (
	"github.com/danielpumayauli/go-auth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:@/go_auth"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}

	connection.AutoMigrate((&models.User{}))
}
