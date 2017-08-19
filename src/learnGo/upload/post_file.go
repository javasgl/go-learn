package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {

	target_url := "http://localhost:8081/upload"
	file_name := "upload.png"

	post_file(file_name, target_url)

}

func post_file(file_name, target_url string) error {

	bodyBuf := &bytes.Buffer{}

	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("upload_file", file_name)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	fh, err := os.Open(file_name)

	if err != nil {
		fmt.Println("error open file")
		return err
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(target_url, contentType, bodyBuf)
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}
