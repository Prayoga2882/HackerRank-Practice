package Golang_Web__test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostForm(writer http.ResponseWriter, request *http.Request) {
	//err := request.ParseForm()
	//if err != nil {
	//	panic(err)
	//}
	name := request.PostFormValue("data")
	//name := request.PostForm.Get("name")
	fmt.Fprintf(writer, "hello %s", name)
}

func TestPostForm(t *testing.T) {
	body := strings.NewReader("data=melvin Boedihartoyo")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/", body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	PostForm(recorder, request)

	result := recorder.Result()
	done, _ := io.ReadAll(result.Body)
	fmt.Println(string(done))
}
