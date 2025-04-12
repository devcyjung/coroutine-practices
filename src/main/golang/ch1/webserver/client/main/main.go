package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = 10
	transport.MaxConnsPerHost = 100
	transport.MaxIdleConnsPerHost = 100
	client := &http.Client{
		Transport: transport,
	}
	request, err := http.NewRequest(
		"GET",
		"http://localhost:8080/path",
		strings.NewReader("message1"),
	)
	if err != nil {
		fmt.Printf("error creating request %+v\n", err)
	}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("error doing request %+v %+v\n", response, err)
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			fmt.Printf("error closing response body %+v %+v\n", response, err)
		}
	}()
	fmt.Printf("response %+v\n", response)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("error reading response body, %+v %+v\n", response.Body, err)
	}
	fmt.Printf("response body: %s\n", body)
}
