package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Print("hello!")
    r := mux.NewRouter()
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)
    http.Handle("/", r)
}
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Category: %v\n", vars["category"])
}