package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	guid := uuid.New().String()
	fmt.Printf("Hello World! %s", guid)
}
