package main

import (
	"dAcademy/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/course/list", func(c *gin.Context) {
			handlers.CourseListHandler(c)
		})
		api.GET("/course/scan", func(c *gin.Context) {
			handlers.CourseScanHandler(c)
		})
		api.GET("/course/:slug", func(c *gin.Context) {
			handlers.CourseDetailHandler(c)
		})

	}

	if err := r.Run(":9090"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
