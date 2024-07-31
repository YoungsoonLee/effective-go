package main

import (
	"fmt"
	"net/http"
)

func crashHander(w http.ResponseWriter, _ *http.Request) {
	go func() {
		panic("boom")
	}()
	fmt.Fprintf(w, "ok")
}

func main() {

}
