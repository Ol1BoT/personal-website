package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Ol1BoT/website/views"
)

var homeView *views.View
var newHomeView *views.View

func main() {
	homeView = views.NewView("views/home.gohtml")
	newHomeView = views.NewView("views/new.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	r.HandleFunc("/new", newHome)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}

func newHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := newHomeView.Render(w, nil); err != nil {
		panic(err)
	}
}
