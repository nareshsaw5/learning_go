package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {

	resp, error := http.Get("http://google.com")

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Print("processed ", len(bs))
	return len(bs), nil
}
