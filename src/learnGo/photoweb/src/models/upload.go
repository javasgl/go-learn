package models

import (
	"io"
	"net/http"
	"os"
)

func Upload(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		RenderHtml(res, "upload", nil)
		return
	}
	if req.Method == "POST" {
		f, h, err := req.FormFile("image")
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		http.Redirect(res, req, "/view?id="+filename, http.StatusFound)
	}
}
