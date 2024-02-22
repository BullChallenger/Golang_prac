package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestuploadFile(t *testing.T) {
	assert := assert.New(t)
	path := "C:\\Users\\am23a\\Downloads\\image (1).png"
	file, _ := os.Open(path)

	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multipart, err := writer.CreateFormFile("upload_file", filepath.Base(path))

	assert.NoError(err)
	io.Copy(multipart, file)
	writer.Close()

	response := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/uploads", buf)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	uploadHandler(response, request)
	assert.Equal(http.StatusOK, response.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, errInStat := os.Stat(uploadFilePath)
	assert.NoError(errInStat)

	uploadFile, _ := os.Open(uploadFilePath)
	origin, _ := os.Open(path)
	defer uploadFile.Close()
	defer origin.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	origin.Read(originData)

	assert.Equal(uploadData, originData)
}
