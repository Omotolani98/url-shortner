package controllers

import (
	"database/sql"
	"github.com/Omotolani98/url-shortner/internal/db"
	"github.com/Omotolani98/url-shortner/internal/url"
	"html/template"
	"net/http"
	"strings"
)

func Shorten(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		originalURL := r.FormValue("url")
		if originalURL == "" {
			http.Error(w, "URL not provided", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(originalURL, "http://") || !strings.HasPrefix(originalURL, "https://") {
			originalURL = "https://" + originalURL
		}

		shortURL := url.Shorten(originalURL)

		if err := db.StoreURL(lite, shortURL, originalURL); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		data := map[string]string{
			"ShortURL": shortURL,
		}

		t, err := template.ParseFiles("internal/views/shorten.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err = t.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func Redirect(lite *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortURL := r.URL.Path[1:]
		if shortURL == "" {
			http.Error(w, "URL not provided", http.StatusBadRequest)
			return
		}

		origUrl, err := db.GetOriginalURL(lite, shortURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		http.Redirect(w, r, origUrl, http.StatusPermanentRedirect)
	}
}
