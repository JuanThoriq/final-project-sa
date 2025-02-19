package main

import (
	"log"

	"final-project-sa-be/database"
	"final-project-sa-be/routes"
)

func main() {
	// Inisialisasi koneksi database
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	// Lakukan auto migration untuk model
	database.DBAutoMigrate(db)

	// Setup router dan jalankan server
	router := routes.SetupRouter()
	router.Run(":8080") // Server berjalan di port 8080
}
