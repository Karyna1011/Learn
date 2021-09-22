package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// запрос в формате http
	response, request := http.Get("https://ru.wikipedia.org/wiki/%D0%90%D1%80%D1%82-%D1%80%D0%BE%D0%BA")
	if request != nil {
		log.Fatal("failed to read file:",request) //Fatal is equivalent to Print() followed by a call to os.Exit(1)
	}
	defer response.Body.Close()  //закрываем файл defer до выхода из функции мейн

	// копируем инфо в нормальный вывод
	n, request := io.Copy(os.Stdout, response.Body)
	if request != nil {
		log.Fatal("failed to copy file",request)
	}

	fmt.Println("number of bytes:", n)
}