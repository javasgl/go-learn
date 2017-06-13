package main

import (
	"html/template"
	"io/ioutil"
	"learnGo/photoweb/src/models"
	"log"
	"net/http"
	"path"
)

func main() {
	log.Print("listen on http://127.0.0.1:8080")
	http.HandleFunc("/upload", models.Upload)
	http.HandleFunc("/view", models.View)
	http.HandleFunc("/list", models.List)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
}

func init() {
	fileInfos, err := ioutil.ReadDir(models.TEMPLATE_DIR)
	models.CheckErr(err)
	var templateName, templatePath string
	for _, fileInfo := range fileInfos {
		templateName = fileInfo.Name()
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}
		templatePath = models.TEMPLATE_DIR + "/" + templateName
		log.Println("Loading template:", templatePath)
		t := template.Must(template.ParseFiles(templatePath))
		models.Templates[templateName] = t
	}
}
