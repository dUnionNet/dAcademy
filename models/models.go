package models

type ChapterData struct {
	ID    int    `yaml:"id" json:"id"`
	Title string `yaml:"title" json:"title"`
}
type CourseData struct {
	Name         string `yaml:"name" json:"name"`
	Description  string `yaml:"description" json:"description"`
	Slug         string `yaml:"slug" json:"slug"`
	Folder       string `yaml:"folder" json:"folder"`
	ChapterCount int    `yaml:"chapter_count" json:"chapter_count"`
}
