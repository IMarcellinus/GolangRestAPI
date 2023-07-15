package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeypc/go-crud/config"
	"github.com/jeypc/go-crud/controllers/productcontroller"
)

func main() {
	r := gin.Default()
	config.DBConnection()

	api := r.Group("/api")

	api.GET("/products", productcontroller.Index)
	api.GET("/products/:id", productcontroller.Show)
	api.POST("/products", productcontroller.Create)
	api.PUT("/products/:id", productcontroller.Update)
	api.DELETE("/products", productcontroller.Delete)
	
	r.Run()
}
