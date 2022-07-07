package main

import (
	"errors"
	"fmt"
	"log"

	"alex/test/internal/service"
)

func main() {
	fmt.Println("Hello, World! - client")

	svc := service.NewService1()
	err := svc.Exec()
	if err != nil {
		if errors.Is(err, service.ErrMyErrorBlue) {
			log.Fatalf("IT'S a blue error: %v", err)
		}

		log.Fatalf("it's another error: %v", err)
	}
}
