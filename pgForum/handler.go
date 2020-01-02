package cms

import (
	"net/http"
	"strings"
	"time"
)

func ServePage(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimLeft(req.URL.Path, "/page/")

	if path == "" {
		http.NotFound(w, req)
		return
	}

	p := &Page{
		Title:   strings.ToTitle(path),
		Content: "here is my page",
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func ServePost(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimLeft(req.URL.Path, "/post/")

	if path == "" {
		http.NotFound(w, req)
		return
	}

	p := &Post{
		Title:   strings.ToTitle(path),
		Content: "Here is my page",
		Comments: []*Comment{
			&Comment{
				Author:        "Trevor Knott",
				Comment:       "Looks great!",
				DatePublished: time.Now(),
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "post", p)

}

func ServeIndex(w http.ResponseWriter, req *http.Request) {
	p := &Page{
		Title:   "TK Go CMS",
		Content: "Welcome to the home page!",
		Posts: []*Post{
			&Post{
				Title:         "Hello, Earth",
				Content:       "Hello, Thank you for being a G.",
				DatePublished: time.Now(),
			},
			&Post{
				Title:         "A Post w/ Comments",
				Content:       "This post will attract comments of course because it is phenomenal.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Trevor Knott",
						Comment:       "What a terrible short post ehh -____-",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p)
}

func HandleNew(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)

	case "POST":
		title := req.FormValue("title")
		content := req.FormValue("content")
		contentType := req.FormValue("content-type")
		req.ParseForm()

		if contentType == "page" {
			p := &Page{
				Title:   title,
				Content: content,
			}
			_, err := CreatePage(p)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			Tmpl.ExecuteTemplate(w, "page", p)
			return
		}
	default:
		http.Error(w, "method not supported: "+req.Method, http.StatusMethodNotAllowed)
	}
}

// TK
