package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nidhey27/go_curd_postgres/controllers"
	"github.com/nidhey27/go_curd_postgres/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  true,
			"message": "GO CURD POSTGRES API",
			"error":   "",
			"author":  "Nidhey Indurkar",
			"data":    make([]string, 0),
		})
	})

	r.POST("/post", controllers.CreatePost)
	r.GET("/post", controllers.GetPosts)
	r.GET("/post/:id", controllers.GetPost)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
