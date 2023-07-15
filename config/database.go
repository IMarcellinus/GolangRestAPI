package config

import (
	"fmt"
	"log"

	"github.com/jeypc/go-crud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	// Open database connection
	// The database is called test
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"))

	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Success Connect Database")

	// Auto Migrate
	database.AutoMigrate(&models.Product{})
	fmt.Println("Success Migrate Database")

	// CRUD
	// Create
	// product := models.Product{}
	// product.NamaProduct = "baju"
	// product.Deskripsi = "baju tidur"

	// err = database.Create(&product).Error
	// if err != nil {
	// 	fmt.Println("============================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("============================")
	// }

	DB = database
}
