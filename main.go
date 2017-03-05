package main

import (
	"fmt"

	"github.com/SuperGod/blog/model"

	"log"
)

func main() {
	users, err := model.GetUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
