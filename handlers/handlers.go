package handlers

import (
	"ascii-art-web-export-file/printascii"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	var err error
	tpl, err = template.ParseFiles("ui/htmlfiles/index.html", "ui/htmlfiles/error.html", "ui/static/style.css")
	if err != nil {
		log.Fatal(err)
	}
}

type ErrorBody struct {
	Status  int
	Message string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		headError(w, http.StatusNotFound)
		return
	}
	if err := tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		headError(w, http.StatusInternalServerError)
		return
	}
}

func ProcessorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		headError(w, http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path != "/ascii-art" {
		headError(w, http.StatusNotFound)
		return
	}
	fname := r.FormValue("string")
	f := r.FormValue("font")
	color := r.FormValue("color")
	ascii, err := printascii.AsciiWeb(fname, f)

	if err == printascii.ErrFont || err == printascii.ErrNonAscii || err == printascii.ErrString {
		headError(w, http.StatusBadRequest)
		return
	} else if err == printascii.ErrTxtFile || err == printascii.ErrRead {
		headError(w, http.StatusInternalServerError)
		return
	}
	if !(color == "white" || color == "black" || color == "red" || color == "pink" || color == "blue") {
		headError(w, http.StatusBadRequest)
		return
	}

	submission := r.FormValue("download")
	if submission == "Download" {
		w.Header().Set("Content-Disposition", "attachment; filename=data.txt")
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Content-Length", strconv.Itoa(len(ascii)))
		w.Write([]byte(ascii))
		return
	}
	d := struct {
		AsciiPrint string
		AsciiColor string
	}{
		AsciiPrint: ascii,
		AsciiColor: color,
	}
	if err = tpl.ExecuteTemplate(w, "index.html", d); err != nil {
		headError(w, http.StatusInternalServerError)
		return
	}
}

func headError(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	eh := setError(status)
	tpl.ExecuteTemplate(w, "error.html", eh)
}

func setError(status int) *ErrorBody {
	return &ErrorBody{
		Status:  status,
		Message: http.StatusText(status),
	}
}
