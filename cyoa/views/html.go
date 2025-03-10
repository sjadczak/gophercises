package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/sjadczak/gophercises/cyoa/story"
)

type HTMLTemplate struct {
	tpl *template.Template
}

func Parse(pattern string) (HTMLTemplate, error) {
	tpl := template.New(pattern)
	tpl, err := tpl.ParseFiles("templates/story.gohtml")
	if err != nil {
		return HTMLTemplate{}, fmt.Errorf("parsing html template: %w", err)
	}

	return HTMLTemplate{tpl}, nil
}

func (t HTMLTemplate) Execute(w http.ResponseWriter, r *http.Request, data story.Arc) {
	tpl, err := t.tpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v\n", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing html template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}

	io.Copy(w, &buf)
}
