package multiparttest

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/puripat-hugeman/go-clean-todo/todo/datamodel"
)

//TODO: Variables name!!
func PrepareMultipartRequest(inputBody datamodel.TodoRequestEntity, image string) (*http.Request, error) {

	//should mock file with os.Create()
	filePointer, err := os.Open(image)
	if err != nil {
		return nil, err
	}
	fileInfo, _ := os.Stat(image)
	// For comparison
	fileBuf := bytes.NewBuffer(nil)
	if _, err := io.Copy(fileBuf, filePointer); err != nil {
		return nil, err
	}
	defer filePointer.Close()
	// MP BODY: JSON Body + image
	body := new(bytes.Buffer)

	mw := multipart.NewWriter(body)
	// Create the form data field filePointer
	dataPart, err := mw.CreateFormField("data")
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(inputBody)
	if err != nil {
		panic("bad json")
	}
	jr := bytes.NewReader(j)
	if _, err := io.Copy(dataPart, jr); err != nil {
		return nil, err
	}

	// Create the form data field 'file'
	filePart, err := mw.CreateFormFile("file", fileInfo.Name())
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(filePart, filePointer); err != nil {
		return nil, err
	}

	mw.Close()
	// Our test HTTP request
	req := httptest.NewRequest("POST", "/todo/create", body)
	req.Header.Add("Content-Type", mw.FormDataContentType())
	return req, nil
}
