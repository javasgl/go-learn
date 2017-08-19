package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/upload", upload)

	log.Println("listen on http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.Method)
	if r.Method == "GET" {
		tpl, _ := template.ParseFiles("upload.html")
		tpl.Execute(w, "")
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("upload_file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload.png", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}
