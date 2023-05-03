package database

import (
	"fmt"
	"waysgallery/models"
	"waysgallery/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Order{},
		&models.Art{},
		&models.Photo{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
