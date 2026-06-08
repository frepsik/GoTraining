package main

import (
	"fmt"
	"net/http"
)

func defaultHandleFunc(w http.ResponseWriter, r *http.Request) {
	fooParam := r.URL.Query().Get("foo")
	booParam := r.URL.Query().Get("boo")

	fmt.Println("foo param:", fooParam)
	fmt.Println("boo param:", booParam)
}

func main() {
	http.HandleFunc("/default", defaultHandleFunc)
	fmt.Println("Стартуем сервер")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("fail start server: ", err)
	}
}
