package main

import (
	"math"
	"net/http"
	"sort"
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

	_ = t.ExecuteTemplate(w, "index.html", nil)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		_ = t.ExecuteTemplate(w, "new.html", nil)

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
		if abs(num) > math.MaxInt/2 {
			showError(w, http.StatusBadRequest, "Number too big to double.")

			return
		}

		dbl := num * 2

		// insert if new
		if _, ok := dblMap[num]; !ok {
			dblMap[num] = dbl

			dblArr = append(dblArr, num)

			sort.Ints(dblArr)
		}

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

	dbls := [][2]string{}

	for _, num := range dblArr {
		dbl, ok := dblMap[num]

		dblStr := "pending"
		if ok {
			dblStr = strconv.Itoa(dbl)
		}

		dbls = append(dbls, [2]string{strconv.Itoa(num), dblStr})
	}

	_ = t.ExecuteTemplate(w, "list.html", struct{ Dbls [][2]string }{dbls})
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

	_ = t.ExecuteTemplate(w, "success.html", nil)
}

func abs(num int) int {
	if num < 0 {
		num *= -1
	}

	return num
}

func showError(w http.ResponseWriter, code int, message string) {
	_ = t.ExecuteTemplate(w, "error.html", struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	})
}
