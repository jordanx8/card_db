package main

import (
	"fmt"

	m "github.com/jordanx8/card_db/mongodb"
)

func main() {
	client, err := m.GetMongoClient()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(client)
}
