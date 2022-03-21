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
var templateFS embed.FS

var (
	templates *template.Template
	dblList   *doubleList
)

func main() {
	seed()
	parseTemplates()
	handleStaticFiles()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/success", successHandler)

	serve()
}

func parseTemplates() {
	parsedTemplates, err := template.ParseFS(templateFS, "templates/*")
	if err != nil {
		log.Fatal(err)
	}

	templates = parsedTemplates
}

func seed() {
	dblList = newDoubleList()
	dblList.add(3)
	dblList.add(8)
}

func serve() {
	log.Println("👂 listening on port 3000")
	log.Panic(http.ListenAndServe(":3000", nil))
}
