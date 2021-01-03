package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Ol1BoT/website/views"
)

var homeView *views.View

func main() {
	fmt.Println("Hello, world")
	homeView = views.NewView("views/home.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	Test := "Testing data sharing"

	if err := homeView.Template.Execute(w, Test); err != nil {
		panic(err)
	}
}
