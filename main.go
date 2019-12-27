package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	// parse our multipart form 10 << 20 specs a max fs of 10 MB
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")

	if err != nil {
		fmt.Println("Error retrieving the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("File: %+v\n", handler.Filename)
	fmt.Printf("Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temp file within a temp-images dir
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tempFile.Close()

	// read all bytes

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {

	fmt.Println("hello world")
	setupRoutes()

}
