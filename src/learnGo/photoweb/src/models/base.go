package models

import (
	"html/template"
	"log"
	"net/http"
	"runtime/debug"
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

func SafeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {

				RenderHtml(w, "error", map[string]interface{}{"err": err.Error(), "stack": string(debug.Stack())})

				log.Printf("Warning:panic in %v - %v ", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()
		fn(w, r)
	}
}
