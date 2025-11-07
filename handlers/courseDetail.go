package handlers

import (
	"dAcademy/models"
	"dAcademy/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func CourseDetailHandler(c *gin.Context) {
	slug := c.Param("slug") // /api/course/detail/:slug
	if slug == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "slug is required"})
		return
	}

	var courses []models.CourseData
	if err := utils.ReadYAML("./courses/_courses.yaml", &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Read and find course by slug
	for _, course := range courses {
		if course.Slug == slug {
			var chapters []models.ChapterData
			chaptersFile := filepath.Join("./courses", course.Folder, "_chapters.yaml")

			if err := utils.ReadYAML(chaptersFile, &chapters); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message":  "course found",
				"course":   course,
				"chapters": chapters,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
}
