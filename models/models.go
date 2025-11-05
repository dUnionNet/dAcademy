package models

type ChapterData struct {
	ID    int    `yaml:"id"`
	Title string `yaml:"title"`
}
type CourseData struct {
	Name         string `yaml:"name" json:"name"`
	Description  string `yaml:"description" json:"description"`
	ChapterCount int    `yaml:"chapter_count" json:"chapter_count"`
}
