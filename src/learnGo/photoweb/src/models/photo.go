package models

import (
	"io/ioutil"
	"net/http"
	"os"
)

func View(res http.ResponseWriter, req *http.Request) {
	imageId := req.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if !IsExists(imagePath) {
		http.NotFound(res, req)
		return
	}
	res.Header().Set("Content-type", "image")
	http.ServeFile(res, req, imagePath)
}
func IsExists(imagePath string) bool {
	_, err := os.Stat(imagePath)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}

func List(w http.ResponseWriter, r *http.Request) {
	fileInfos, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfos {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	err = RenderHtml(w, "list", locals)
	CheckErr(err)
}
