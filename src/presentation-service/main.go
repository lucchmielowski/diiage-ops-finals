package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/get-key", string(os.Getenv("CATALOG_API_URL")))
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Handled request")

	io.WriteString(w, string(body))
}

func main() {
	if os.Getenv("CATALOG_API_URL") == "" {
		log.Fatalln("Missing CATALOG_API_URL environment variable ")
	}
	http.HandleFunc("/", Handler)
	err := http.ListenAndServe(":4444", nil)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server running on prot 4444 ...")
}
