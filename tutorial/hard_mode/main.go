package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1] // access the first arguments during call

	/**
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
	*/
	file, error := os.Open(fileName)
	if error != nil {
		fmt.Print(error)
	}
	/**
	bs := make([]byte, 99999)
	n, eerror := file.Read(bs)
	if eerror != nil {
		fmt.Print(eerror)
	}
	fmt.Printf("Read upto %d character ", n)
	*/
	io.Copy(os.Stdout, file)

}
