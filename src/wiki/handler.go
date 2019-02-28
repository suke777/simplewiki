package wiki

import (
	"fmt"
	"net/http"
	"text/template"
	"unicode/utf8"
)

const (
	viewHandlerURL  = "/"
	editHandlerURL  = "/edit/"
	saveHandlerURL  = "/save/"
	errorHandlerURL = "/error/"
)

func registerViewHandler() {
	http.HandleFunc(viewHandlerURL, viewHandler)
}

func registerEditHandler() {
	http.HandleFunc(editHandlerURL, editHandler)
}

func registerSaveHandler() {
	http.HandleFunc(saveHandlerURL, saveHandler)
}

func registerErrorHandler() {
	http.HandleFunc(errorHandlerURL, errorHandler)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	titleIndex := utf8.RuneCountInString(viewHandlerURL)
	title := r.URL.Path[titleIndex:]

	page, err := loadPageData(title)
	if err != nil {
		page = &PageData{Title: title}
	}
	viewPage := &PageData{page.Title, nl2Br(page.Body)}

	t, _ := template.ParseFiles("wiki/view.html.tmpl")
	t.Execute(w, viewPage)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	titleIndex := utf8.RuneCountInString(editHandlerURL)
	title := r.URL.Path[titleIndex:]

	page, err := loadPageData(title)
	if err != nil {
		page = &PageData{Title: title}
	}

	t, _ := template.ParseFiles("wiki/edit.html.tmpl")
	t.Execute(w, page)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	titleIndex := utf8.RuneCountInString(saveHandlerURL)
	title := r.URL.Path[titleIndex:]
	body := r.FormValue("body")

	err := savePageData(title, []byte(body))
	if err != nil {
		http.Redirect(w, r, errorHandlerURL+title, http.StatusFound)
	} else {
		http.Redirect(w, r, viewHandlerURL+title, http.StatusFound)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "An error has occurred. %q", r.URL.Path[7:])
}
