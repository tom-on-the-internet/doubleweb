package main

import (
	"net/http"
	"strconv"
)

// handleStaticFiles takes embed files
// and serves them up "as is".
func handleStaticFiles() {
	staticFS := http.FS(staticFiles)
	fs := http.FileServer(staticFS)
	http.Handle("/static/", fs)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			r.Method+" is not allowed for this endpoint.",
			http.StatusMethodNotAllowed,
		)

		return
	}

	_ = templates.ExecuteTemplate(w, "index.html", nil)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_ = templates.ExecuteTemplate(w, "new.html", nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			showError(w, http.StatusBadRequest, "Could not parse input.")

			return
		}

		// start here
		numStr := r.PostForm.Get("number")

		num, err := strconv.Atoi(numStr)
		if err != nil {
			showError(w, http.StatusBadRequest, "Invalid number.")

			return
		}

		// check for overflow
		if !dblList.canHandle(num) {
			showError(w, http.StatusBadRequest, "Number too big to double.")

			return
		}

		dblList.add(num)

		http.Redirect(w, r, "/success", http.StatusSeeOther)

	default:
		http.Error(
			w,
			r.Method+" is not allowed for this endpoint.",
			http.StatusMethodNotAllowed,
		)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			r.Method+" is not allowed for this endpoint.",
			http.StatusMethodNotAllowed,
		)

		return
	}

	_ = templates.ExecuteTemplate(w, "list.html", struct{ Dbls [][2]string }{dblList.list()})
}

func successHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(
			w,
			r.Method+" is not allowed for this endpoint.",
			http.StatusMethodNotAllowed,
		)

		return
	}

	_ = templates.ExecuteTemplate(w, "success.html", nil)
}

func showError(w http.ResponseWriter, code int, message string) {
	_ = templates.ExecuteTemplate(w, "error.html", struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	})
}
