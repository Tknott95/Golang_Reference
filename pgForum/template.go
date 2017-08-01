package cms

import (
	"html/template"
	"time"
)

// Capitol allows to be exported
var Tmpl = template.Must(template.ParseGlob("../templates/*"))

type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []*Post
}

type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

type Comment struct {
	ID            int
	Author        string
	Comment       string
	DatePublished time.Time
}
