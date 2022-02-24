package main

import (
	"flag"
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {

	zero := flag.Bool("nil", false, "generate nil uuid")
	flag.Parse()

	var id uuid.UUID
	if *zero {
		id = uuid.Nil
	} else {
		id = uuid.Must(uuid.NewV4())
	}
	fmt.Println(id)
}
