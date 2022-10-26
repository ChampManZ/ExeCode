package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// pistonClient := piston.New(http.DefaultClient, "127.0.0.1:2000", "")
	// runtimes, err := pistonClient.GetRuntimes()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	req, err := http.NewRequest("GET", "http://127.0.0.1:2000/api/v2/runtimes", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	runtimes := string(b)
	fmt.Println(runtimes)
}
