package handler

import (
	"net/http"
	"text/template"
)

func (h *Handler) VariationsPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"pkg/app/admin/api/web/templates/base.html",
		"pkg/app/admin/api/web/templates/variations.html",
	}

    ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		ApiBaseUrl string
        ActivePage string
	}{
		ApiBaseUrl: h.ApiBaseUrl,
        ActivePage: "variation",
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
