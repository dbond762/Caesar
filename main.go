package main // import "github.com/dbond762/caesar-backend"

import (
	"encoding/json"
	"fmt"
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
	Freqs []float64 `json:"freqs"`
	Shift int       `josn:"shift"`
}

func index(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var in input
	err := decoder.Decode(&in)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	out := caesar(in)

	encoder := json.NewEncoder(w)
	err = encoder.Encode(out)
	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func main() {
	//http.HandleFunc("/", index)
	//log.Fatal(http.ListenAndServe(":8000", nil))
	fmt.Println(shift("Hello, world!", 3))
}
