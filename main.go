package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	err = callAPI()
	if err != nil {
		fmt.Println(err)
	}

	botCycle()
}
