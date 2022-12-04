package main

import (
	"log"

	"github.com/ChampManZ/ExeCode/v2/internal/auth"
)

func main() {
	env, _ := auth.GetEnv()

	err := env.InitKeys()
	if err != nil {
		log.Fatalln(err)
	}

}
