package handlers

import (
	"dAcademy/models"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	courseScanMutex  sync.Mutex
	isCourseScanning bool
)

func CourseScanHandler(c *gin.Context) {
	courseScanMutex.Lock()
	if isCourseScanning {
		courseScanMutex.Unlock()
		c.JSON(http.StatusConflict, gin.H{"message": "A scan job is already in progress."})
		return
	}

	isCourseScanning = true
	courseScanMutex.Unlock()

	go func() {
		defer func() {
			courseScanMutex.Lock()
			isCourseScanning = false
			courseScanMutex.Unlock()
			log.Println("Courses scan process finished.")
		}()

		fmt.Println("Courses scan process starting.")
		if err := scanAll("./courses"); err != nil {
			log.Printf("❌ Scan failed: %v", err)
		} else {
			log.Println("✅ All courses indexed successfully.")
		}
	}()

	c.JSON(http.StatusAccepted, gin.H{"message": "Scan job has been accepted."})
}

// Scan all(courses and their chapters)
func scanAll(root string) error {
	var courses []models.CourseData
	coursesYaml := filepath.Join(root, "_courses.yaml")

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if !d.IsDir() {
			return nil
		}

		courseYaml := filepath.Join(path, "course.yaml")
		chaptersYaml := filepath.Join(path, "_chapters.yaml")

		if _, err := os.Stat(courseYaml); err == nil {
			fmt.Println("📘 構建章節索引:", path)

			chapters, err := buildChapters(path)
			if err != nil {
				return fmt.Errorf("構建章節索引失敗 %s: %v", path, err)
			}

			if len(chapters) > 0 {
				if err := saveYamlFile(chaptersYaml, chapters); err != nil {
					return fmt.Errorf("寫入 _chapters.yaml 失敗: %v", err)
				}
			}

			data, err := os.ReadFile(courseYaml)
			if err != nil {
				return fmt.Errorf("讀取 course.yaml 失敗 %s: %v", path, err)
			}

			var course models.CourseData
			if err := yaml.Unmarshal(data, &course); err != nil {
				return fmt.Errorf("解析 course.yaml 失敗 %s: %v", path, err)
			}
			course.ChapterCount = len(chapters)
			courses = append(courses, course)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("掃描課程目錄失敗: %v", err)
	}

	if err := saveYamlFile(coursesYaml, courses); err != nil {
		return fmt.Errorf("寫入 _courses.yaml 失敗: %v", err)
	}

	fmt.Println("✅ Scan finished, total: ", len(courses), " courses")
	return nil
}

// Scan chapter folders
func buildChapters(coursePath string) ([]models.ChapterData, error) {
	files, err := os.ReadDir(coursePath)
	if err != nil {
		return nil, err
	}

	var chapters []models.ChapterData
	for _, f := range files {
		if f.IsDir() {
			id, title, ok := parseChapterFolder(f.Name())
			if !ok {
				log.Fatalf("❌ 無效章節資料夾: %s", f.Name())
			}
			chapters = append(chapters, models.ChapterData{ID: id, Title: title})
		}
	}

	// Sort by ID
	sort.Slice(chapters, func(i, j int) bool { return chapters[i].ID < chapters[j].ID })
	return chapters, nil
}

// Save data as yaml
func saveYamlFile[T any](file string, data T) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("❌ Failed to marshal YAML: %v", err)
	}
	return os.WriteFile(file, bytes, 0644)
}

// Parse chapter folder name, like 1-Hello
func parseChapterFolder(name string) (int, string, bool) {
	parts := strings.SplitN(name, "-", 2)
	if len(parts) != 2 {
		return 0, "", false
	}
	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", false
	}
	return id, parts[1], true
}
