package Golang_Web__test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHancler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello handler")
}

func TestName(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/hello", nil)
	recorder := httptest.NewRecorder()

	HelloHancler(recorder, request)

	result := recorder.Result()
	done, _ := io.ReadAll(result.Body)
	fmt.Println(string(done))
}
