package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
)

//go:embed static
var staticFiles embed.FS

//go:embed templates
var templates embed.FS

var (
	t      *template.Template
	dblMap = make(map[int]int)
	dblArr []int
)

func main() {
	parseTemplates()
	handleStaticFiles()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/success", successHandler)

	serve()
}

func parseTemplates() {
	parsedTemplates, err := template.ParseFS(templates, "templates/*")
	if err != nil {
		log.Fatal(err)
	}

	t = parsedTemplates
}

func serve() {
	log.Println("👂 listening on port 3000")
	log.Panic(http.ListenAndServe(":3000", nil))
}
