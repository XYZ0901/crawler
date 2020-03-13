package view

import (
	"crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	template_ *template.Template
}

func CreateSearchResultView(filename string) SearchResultView {
	return SearchResultView{template_: template.Must(template.ParseFiles(filename))}
}

func (s SearchResultView) Render(w io.Writer, data model.SearchResult) error {
	return s.template_.Execute(w, data)
}
