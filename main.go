package main

import (
	"flag"
	"fmt"

	"github.com/gofrs/uuid"
)

func main() {

	zero := flag.Bool("nil", false, "generate nil uuid")
	v1 := flag.Bool("v1", false, "generate version 4")
	v3 := flag.Bool("v3", false, "generate version 3")
	v4 := flag.Bool("v4", false, "generate version 4")
	v5 := flag.Bool("v5", false, "generate version 5")
	v6 := flag.Bool("v6", false, "generate version 6")
	v7 := flag.Bool("v7", false, "generate version 7")
	namespace := flag.String("namespace", "", "namespace used for version 3 & 4 uuid's")
	name := flag.String("name", "", "name used for version 3 & 4 uuid's")

	flag.Parse()

	var id uuid.UUID
	switch {
	case *zero:
		id = uuid.Nil
	case *v1:
		id = uuid.Must(uuid.NewV1())
	case *v3:
		ns := uuid.FromStringOrNil(*namespace)
		id = uuid.NewV3(ns, *name)
	case *v4:
		id = uuid.Must(uuid.NewV4())
	case *v5:
		ns := uuid.FromStringOrNil(*namespace)
		id = uuid.NewV5(ns, *name)
	case *v6:
		id = uuid.Must(uuid.NewV6())
	case *v7:
		// In the current RFC Draft version 4, the precision is always milliseconds
		// https://datatracker.ietf.org/doc/html/draft-peabody-dispatch-new-uuid-format-04#name-uuid-version-7
		id = uuid.Must(uuid.NewV7(uuid.MillisecondPrecision))
	default:
		id = uuid.Must(uuid.NewV4())
	}
	fmt.Println(id)
}
