package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		log.Printf("[PRIVATE] Hello %v!!", args[1])
		return
	}
	log.Println("[PRIVATE] Hello world!!")
}
