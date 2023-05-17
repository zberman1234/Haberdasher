package main

import (
	"html/template"
	"net/http"

	"github.com/example/internal/haberdasherserver"
	"github.com/example/rpc/haberdasher"
)

func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)
	// http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler)

	http.ListenAndServe(":8080", twirpHandler)
	// http.ListenAndServe(":8081", nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template file
	t, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template, passing any necessary data
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
