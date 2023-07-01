package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hafidzadha23/task-5-vix-btpns-Hafidz-Adha/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Inisialisasi koneksi database
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Inisialisasi router
	r := gin.Default()
	router.SetupRouter(r, db)

	// Menjalankan aplikasi
	r.Run(":8080")
}
