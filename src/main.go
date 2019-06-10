package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("template", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "Application port")
	flag.Parse()
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/room", r)
	// starts the room
	go r.run()
	// start the web server
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
