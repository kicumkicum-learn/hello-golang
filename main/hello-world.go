package main

import (
	"fmt"
	"../web"
)


func main() {
	request := web.Request{}
	result, err := request.Get("http://ya.ru")
	fmt.Println("hello world", result, err)
}
