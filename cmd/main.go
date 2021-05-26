package main

import (
	"fmt"
	"os"
	"log"
)

func main () {
	log.Println("NameSorter Starting")
	arg := os.Args[3]

	if arg == "" {
		fmt.Println("Please add a file as the command line arguement")
	}

	log.Println("NameSorter Ended")
}