package main // import "github.com/dbond762/caesar-backend"

import (
	"encoding/json"
	"log"
	"net/http"
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
	if r.Method != http.MethodPost {
		log.Println("Index method not allowed: ", r.Method)
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var in input
	err := decoder.Decode(&in)
	if err != nil {
		log.Println("Index json decode error: ", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	out := caesar(in)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(out)
	if err != nil {
		log.Println("Index json encode error: ", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Println("Server run on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
