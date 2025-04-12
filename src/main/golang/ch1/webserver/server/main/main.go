package main

import (
	"examples/ch1/animated_gif"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count uint
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/gif", gifHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	_, err := fmt.Fprintf(writer, "URL Path: %q\n", request.URL.String())
	if err != nil {
		log.Printf("error writing response: %+v\n", err)
	}
	buf, err := io.ReadAll(request.Body)
	if err != nil {
		log.Printf("error reading request body: %+v\n", err)
	}
	_, err = fmt.Fprintf(writer, "%d visitors before you\n", count)
	count++
	if err != nil {
		log.Printf("error writing count response: %+v\n", err)
	}
	_, err = fmt.Fprintf(writer, "Echo: %q\n", buf)
	if err != nil {
		log.Printf("error writing response: %+v\n", err)
	}
}

func gifHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		return
	}
	err := animated_gif.DrawLissajousFigure(
		writer,
		animated_gif.LissajousFigure{},
	)
	if err != nil {
		log.Printf("error making gif: %+v\n", err)
	}
}
