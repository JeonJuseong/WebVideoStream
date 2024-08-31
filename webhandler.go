package main

import (
	"html/template"
	"net/http"
	"strings"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := "./websource/index.html"
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

// HTML에서 요청하는 css등의 websource serve하는 함수.
func serveWebsource(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/websource/")
	http.ServeFile(w, r, "./websource/"+path)
}

func serveVideo(w http.ResponseWriter, r *http.Request, filePath string) {
	path := strings.TrimPrefix(r.URL.Path, "/video/")
	//fmt.Println("path: ", path)
	http.ServeFile(w, r, filePath+path)
}
