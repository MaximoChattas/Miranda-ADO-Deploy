package db

import (
	"os"
	"project/client"
	"project/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {

	dsn := os.Getenv("DBCONNSTRING")

	log.Info(dsn)

	Db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// Add all clients here
	client.Db = Db

}

func StartDbEngine() {
	// Migrate all model classes
	Db.AutoMigrate(&model.Hotel{})
	Db.AutoMigrate(&model.Reservation{})
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Amenity{})
	Db.AutoMigrate(&model.Image{})

	log.Info("Finishing Migration Database Tables")
}
