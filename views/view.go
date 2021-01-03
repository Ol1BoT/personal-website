package views

import "html/template"

//NewView returns a View with the necessary templates
func NewView(files ...string) *View {

	files = append(files, "views/layouts/mainLayout.gohtml")
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
