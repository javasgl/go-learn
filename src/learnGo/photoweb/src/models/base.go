package models

import (
	"html/template"
	"net/http"
)

var (
	UPLOAD_DIR   = `./uploads`
	TEMPLATE_DIR = "./views"
	Templates    = make(map[string]*template.Template)
)

func RenderHtml(w http.ResponseWriter, temp string, locals map[string]interface{}) error {
	return Templates[temp+".html"].Execute(w, locals)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
