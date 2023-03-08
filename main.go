package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.New().String()
	fmt.Println("Hello Go")
	fmt.Println(uuid)
}
