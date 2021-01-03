package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Ol1BoT/website/views"
)

var homeView *views.View

func main() {
	homeView = views.NewView("views/home.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	fmt.Println("Listening on port 3000")
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if err := homeView.Render(w, nil); err != nil {
		panic(err)
	}
}
