package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ChampManZ/ExeCode/v2/internal/piston"
)

func main() {
	client := piston.NewClient(http.DefaultClient, "localhost:2000", "")

	_, err := piston.EnsurePackagesFromFile("piston-packages.txt", client)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Finished")
}
