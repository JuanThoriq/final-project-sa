package database

import (
	"log"

	"final-project-sa-be/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	// Ganti dengan detail koneksi PostgreSQL milikmu
	dsn := "host=localhost user=postgres password=Tanggallahir9 dbname=final_project_ruangguru port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}

func DBAutoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.CV{}, &models.Skill{}); err != nil {
		log.Fatal("Failed to auto migrate: ", err)
	}
}

