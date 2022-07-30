package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// One eigth of a megabyte
	buffer := make([]byte, 1024*128)

	resp, err := http.Get("http://127.0.0.1:3333")
	if err != nil {
		log.Printf("%s", err.Error())
		panic("Error getting data")
	}
	for {
		num, err := resp.Body.Read(buffer)
		log.Printf("%d bytes read", num)

		if err == io.EOF {
			break
		} else if err != nil {
			panic("Error reading data")
		}
		time.Sleep(1 * time.Second)
	}

}
