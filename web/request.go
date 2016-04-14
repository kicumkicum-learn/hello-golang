package web

import (
	"fmt"
	"net/http"
)

type Request struct {
	Type string
}

func (server *Request) Get(url string) (*http.Response, error) {
	result, err := http.Get(url)
	if err == nil {
		fmt.Println("Error", err)
		return result, err
	}
	return result, err
}
