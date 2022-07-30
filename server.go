package main

import (
	"fmt"
	"log"
	"net/http"
)

var data = []byte{}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	// Will write 16 megabytes of data overall.
	for i := 0; i < 16; i++ {
		num, err := w.Write(data)
		log.Printf("%d bytes written", num)
		if err != nil {
			panic("Error writing")
		}
	}
}

func main() {
	// one megabyte of data
	for i := 0; i < 64*1024; i++ {
		data = append(data, '1')
		data = append(data, '2')
		data = append(data, '3')
		data = append(data, '4')
		data = append(data, '5')
		data = append(data, '6')
		data = append(data, '7')
		data = append(data, '8')
		data = append(data, '9')
		data = append(data, 'A')
		data = append(data, 'B')
		data = append(data, 'C')
		data = append(data, 'D')
		data = append(data, 'E')
		data = append(data, 'F')
		data = append(data, '\n')
	}
	http.HandleFunc("/", getRoot)

	http.ListenAndServe("127.0.0.1:3333", nil)
}
