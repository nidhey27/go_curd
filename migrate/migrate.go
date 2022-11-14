package main

import (
	"github.com/nidhey27/go_curd_postgres/initializers"
	"github.com/nidhey27/go_curd_postgres/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Posts{})
}
