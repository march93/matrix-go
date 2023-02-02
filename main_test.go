package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	// Create the form data to pass in our request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	// Serve up our request
	res, err := http.Post(ts.URL+"/echo", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to echo: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusOK, res.StatusCode)
	}

	builder := new(strings.Builder)
	io.Copy(builder, res.Body)

	// Trim the response because we are passing a new line after the result
	if strings.Compare(strings.Trim(builder.String(), "\r\n"), "1,2,3\n4,5,6\n7,8,9") != 0 {
		t.Errorf("expected content of \n%s\ngot\n%s\n", "1,2,3\n4,5,6\n7,8,9", builder.String())
	}
}

func TestInvert(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/invert", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to invert: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusOK, res.StatusCode)
	}

	builder := new(strings.Builder)
	io.Copy(builder, res.Body)

	// Trim the response because we are passing a new line after the result
	if strings.Compare(strings.Trim(builder.String(), "\r\n"), "1,4,7\n2,5,8\n3,6,9") != 0 {
		t.Errorf("expected content of %s got %s", "1,4,7\n2,5,8\n3,6,9", builder.String())
	}
}

func TestFlatten(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/flatten", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to flatten: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusOK, res.StatusCode)
	}

	builder := new(strings.Builder)
	io.Copy(builder, res.Body)

	// Trim the response because we are passing a new line after the result
	if strings.Compare(strings.Trim(builder.String(), "\r\n"), "1,2,3,4,5,6,7,8,9") != 0 {
		t.Errorf("expected content of %s got %s", "1,2,3,4,5,6,7,8,9", builder.String())
	}
}

func TestSum(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/sum", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to sum: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusOK, res.StatusCode)
	}

	builder := new(strings.Builder)
	io.Copy(builder, res.Body)

	// Trim the response because we are passing a new line after the result
	if strings.Compare(strings.Trim(builder.String(), "\r\n"), "45") != 0 {
		t.Errorf("expected content of %s got %s", "45", builder.String())
	}
}

func TestMultiply(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/multiply", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to multiply: %s", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusOK, res.StatusCode)
	}

	builder := new(strings.Builder)
	io.Copy(builder, res.Body)

	// Trim the response because we are passing a new line after the result
	if strings.Compare(strings.Trim(builder.String(), "\r\n"), "362880") != 0 {
		t.Errorf("expected content of %s got %s", "362880", builder.String())
	}
}

func TestInvalidData(t *testing.T) {
	fileName := "matrix-invalid.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/echo", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to echo: %s", err)
	}

	if res.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusInternalServerError, res.StatusCode)
	}
}

func TestEmpty(t *testing.T) {
	fileName := "matrix-empty.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/echo", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to echo: %s", err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusBadRequest, res.StatusCode)
	}
}

func TestBadUrl(t *testing.T) {
	fileName := "matrix.csv"
	filePath := path.Join("./", fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, err := http.Post(ts.URL+"/bad", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatalf("error making POST request to bad: %s", err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("expected HTTP status of %d, got %d", http.StatusNotFound, res.StatusCode)
	}
}
