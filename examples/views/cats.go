package views

import (
	"github.com/joncalhoun/viewcon"
	"html/template"
	"log"
	"path/filepath"
)

type CatsView struct {
	viewcon.View
	Feed viewcon.Page // A custom page for the cats view
}

var Cats CatsView

func CatsFiles() []string {
	files, err := filepath.Glob("templates/cats/includes/*.tmpl")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, viewcon.LayoutFiles()...)
	return files
}

func init() {
	indexFiles := append(CatsFiles(), "templates/cats/index.tmpl")
	Cats.Index = viewcon.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "my_layout",
	}

	feedFiles := append(CatsFiles(), "templates/cats/feed.tmpl")
	Cats.Feed = viewcon.Page{
		Template: template.Must(template.New("feed").ParseFiles(feedFiles...)),
		Layout:   "other_layout",
	}
}
