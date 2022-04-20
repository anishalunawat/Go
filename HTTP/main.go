package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	/*	bs := make([]byte, 99999)
		resp.Body.Read(bs)
		fmt.Println(string(bs))
	*/

	lw := logWriter{}
	//io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)

}

func (lr logWriter) Write(bt []byte) (int, error) {
	fmt.Println(string(bt))
	fmt.Println("Processed this much bytes:", len(bt))
	return len(bt), nil
}