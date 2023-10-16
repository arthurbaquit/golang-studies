package main

import (
	"fmt"

	"github.com/gofrs/uuid"
)

const NUMBER_OF_UUIDS = 4

func main() {
	for i := 0; i < NUMBER_OF_UUIDS; i++ {
		fmt.Println(uuid.NewV7())
	}
}
