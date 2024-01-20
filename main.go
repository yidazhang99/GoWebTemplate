package main

import (
	"fmt"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "<html><head><title>HelloWorld Example Page</title></head><body><h1>Hello, world!</h1>Welcome to the world of HTML.<p>This is one paragraph.</p><p>This is a second paragraph.</p></body></html>")
		if err !=nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})
	_ = http.ListenAndServe(":1111",nil)
}
