package main // import "github.com/dbond762/caesar-backend"

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type input struct {
	Text   string `json:"text"`
	Shift  int    `json:"shift"`
	Encode bool   `json:"encode"`
}

type output struct {
	Text  string    `json:"text"`
	Shift int       `josn:"shift"`
	Freqs []float64 `json:"freqs"`
}

func index(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var in input
	err := decoder.Decode(&in)
	if err != nil {
		log.Println("Index, json decode error: ", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	out := caesar(in)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(out)
	if err != nil {
		log.Println("Index, json encode error: ", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func main() {
	r := chi.NewRouter()

	CORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(CORS.Handler)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))
	r.Use(middleware.Logger)

	r.Post("/", index)

	log.Printf("Server run on http://localhost:8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
