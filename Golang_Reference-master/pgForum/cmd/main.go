package main

import (
	"net/http"

	"github.com/tknott95/cms"
)

func main() {
	http.HandleFunc("/", cms.ServeIndex)
	http.HandleFunc("/new/", cms.HandleNew)
	http.HandleFunc("/page/", cms.ServePage)
	http.HandleFunc("/post/", cms.ServePost)
	http.ListenAndServe(":8080", nil)
	// p := &cms.Page{
	// 	Title:   "Hello, world!",
	// 	Content: "this is the body of our webpage",
	// }
}
