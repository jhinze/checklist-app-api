package main

import (
	"checklist/database"
	"checklist/models"
	"checklist/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	var err error
	database.DB, err = gorm.Open(database.Dialect(), database.Args())
	if err != nil {
		log.Fatalln(err)
	}
	defer database.DB.Close()

	database.DB.LogMode(database.LogMode())
	database.DB.AutoMigrate(&models.Item{})
	database.DB.AutoMigrate(&models.List{})

	r := routes.SetupRouter()

	runErr := r.Run()
	if runErr != nil {
		log.Fatalln(runErr)
	}
}
