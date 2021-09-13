package main

import (
	"fmt"

	"github.com/tas1999/testswag/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

func main() {
	fmt.Println(docs.SwaggerInfo)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run("localhost:8082") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
