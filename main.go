package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("ID: ", id)
}
