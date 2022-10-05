package Golang_Web__test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "Application/json")
	fmt.Fprint(writer, "OK")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000", nil)
	request.Header.Add("Content-Type", "Application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	result := recorder.Result()
	done, _ := io.ReadAll(result.Body)
	fmt.Println(string(done))

	fmt.Println(result.Header.Get("Content-Type"))
}
