package main // import "github.com/dbond762/caesar-backend"

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
