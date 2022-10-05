package Golang_Web__test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func QueryParameter(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query().Get("name")
	if queryParams == "" {
		fmt.Fprint(writer, "hello no query")
		return
	}
	fmt.Fprint(writer, "hello query", queryParams)
}

func QueryMultipleParamenter(writer http.ResponseWriter, request *http.Request) {
	queryParams := request.URL.Query()
	params := queryParams["name"]
	fmt.Fprint(writer, strings.Join(params, " "))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/hello?name=melvin&name=alex", nil)
	recorder := httptest.NewRecorder()

	QueryMultipleParamenter(recorder, request)
	result := recorder.Result()
	final, _ := io.ReadAll(result.Body)
	fmt.Println(string(final))
}
