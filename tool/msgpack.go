package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/vmihailenco/msgpack"

	"github.com/GustavoCielo/url-shortener/shortener"
)

func main() {
	address := fmt.Sprintf("localhost%s", httpPort())
	redirect := shortener.Redirect{}
	redirect.URL = "https://github.com/GustavoCielo/url-shortener"

	body, err := msgpack.Marshal(&redirect)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(address, "application/x-msgpack", bytes.NewBuffer(body))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	msgpack.Unmarshal(body, &redirect)

	log.Printf("%v\n", redirect)
}

// httpPort returns the port to listen
func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}
