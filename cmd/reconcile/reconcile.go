package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/firodj/interview_reconcile/lib"
)

func main() {
	var internalPath lib.ArgPathValue
	var externalPath lib.ArgPathValue

	flag.Var(&internalPath, "i", "internal data path (eg system transactions)")
	flag.Var(&externalPath, "x", "external data path (eg bank transactions)")
	flag.Parse()

	if internalPath.String() == "" {
		println("missing internal data path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if externalPath.String() == "" {
		println("missing external data path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Reconcile.")
}
