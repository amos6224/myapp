package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"/Users/jd/tutorial/myapp/ui/html/home.page.tmpl",
		"/Users/jd/tutorial/myapp/ui/html/base.layout.tmpl",
		"/Users/jd/tutorial/myapp/ui/html/footer.partial.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) gilbert(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Gilbert does not like to read thus he is killing me`))
}

func (app *application) createPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method Not allowed", 405)
		return
	}
	w.Write([]byte(`Create a new Post...`))
}

func (app *application) showPost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display my post with ID %d...", id)
}
