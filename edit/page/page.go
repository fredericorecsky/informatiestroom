package page

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/fredericorecsky/informatiestroom/auth"
	"github.com/fredericorecsky/informatiestroom/documents/page"
)

// TODO lib
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404 not found")
	}
	if status == http.StatusUnauthorized {
		fmt.Fprint(w, "401 not authorized")
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if auth.Logged(w, r) != nil {
		errorHandler(w, r, http.StatusUnauthorized)
		return
	}

	address := r.URL.Path[len("/edit/"):]

	fmt.Println("Stay clean")

	page := &Document.Page{Filename: address}

	if r.Method == "POST" {

		page.Body = []byte(r.FormValue("content"))
		err := page.Save()
		if err == nil {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			errorHandler(w, r, http.StatusBadRequest)
			return
		}
	} else {
		err := page.Load()

		fmt.Println(address)
		fmt.Println(err)
		if err != nil {
			errorHandler(w, r, http.StatusNotFound)
			return
		}
		t, err := template.ParseFiles("./Templates/page/edit.t")
		err = t.Execute(w, page)
	}
}

func init() {
	http.HandleFunc("/edit/", editHandler)
}
