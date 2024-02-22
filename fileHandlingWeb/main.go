package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(writer http.ResponseWriter, reader *http.Request) {
	uploadFile, header, err := reader.FormFile("upload_file")
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, err)
		return
	}
	defer uploadFile.Close()

	dirname := "./uploads"
	os.Mkdir(dirname, 077)
	filePath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(writer, err)
		return
	}

	io.Copy(file, uploadFile)
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, filePath)
}

func main() {
	http.HandleFunc("/uploads", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}
