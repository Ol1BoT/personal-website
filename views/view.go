package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

//NewView returns a View with the necessary templates
func NewView(files ...string) *View {

	files = append(files, fetchLayouts()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
	}
}

//View is a struct containing the parsed templates
type View struct {
	Template *template.Template
}

//Render renders the webpage
func (v View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.Execute(w, data)
}

func fetchLayouts() []string {

	files, err := filepath.Glob("views/layouts/*.gohtml")
	if err != nil {
		panic(err)
	}

	return files

}
