package models

type News struct {
	Id         string
	Title      string
	Content    string
	AllContent string
}

func NewNews(id, title, content, allContent string) *News {
	return &News{id, title, content, allContent}
}
