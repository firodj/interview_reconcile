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
	var fromDate lib.ArgDateValue
	var toDate lib.ArgDateValue

	flag.Var(&internalPath, "i", "internal data path (eg system transactions)")
	flag.Var(&externalPath, "x", "external data path (eg bank transactions)")
	flag.Var(&fromDate, "f", "from date (format YYYY-MM-DD)")
	flag.Var(&toDate, "t", "to date (format YYYY-MM-DD)")
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

	if fromDate.String() == "" {
		println("missing from date")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if toDate.String() == "" {
		println("missing to date")
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Reconcile.")
}
