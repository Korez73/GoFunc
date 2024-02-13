package lib

import (
	"fmt"
	"log"
	"os"
)

func ReadFile() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	fmt.Println("Make it this far")
}
