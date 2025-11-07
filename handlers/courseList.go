package handlers

import (
	"dAcademy/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

func CourseListHandler(c *gin.Context) {
	yamlPath := "./courses/_courses.yaml"

	// Read
	data, err := os.ReadFile(yamlPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read _courses.yaml: " + err.Error(),
		})
		return
	}

	// Parse YAML into struct slice
	var courses []models.CourseData
	if err := yaml.Unmarshal(data, &courses); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse YAML: " + err.Error(),
		})
		return
	}

	// Output JSON response
	c.JSON(http.StatusOK, gin.H{
		"courses": courses,
	})
}
