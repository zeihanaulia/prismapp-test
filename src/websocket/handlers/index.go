package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// templ represents a single template
type TemplateHandler struct {

	// sync.Once menjamin jika fungsi yang akan dilewati hanya diesekusi sekali saja
	// tidak peduli seberapa sering goroutines memanggil method ServerHTTP
	Once     sync.Once
	Filename string
	Templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.Once.Do(func() {
		t.Templ = template.Must(template.ParseFiles(filepath.Join("../../views", t.Filename)))
	})
	t.Templ.Execute(w, nil)
}
