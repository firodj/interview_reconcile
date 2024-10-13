package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/firodj/interview_reconcile/lib"
)

type AppArg struct {
	InternalPath lib.ArgPathValue
	ExternalPath lib.ArgPathValue
	FromDate     lib.ArgDateValue
	ToDate       lib.ArgDateValue
}

func (app *AppArg) Scan() {
	flag.Var(&app.InternalPath, "i", "internal data path (eg system transactions)")
	flag.Var(&app.ExternalPath, "x", "external data path (eg bank transactions)")
	flag.Var(&app.FromDate, "f", "from date (format YYYY-MM-DD)")
	flag.Var(&app.ToDate, "t", "to date (format YYYY-MM-DD)")
	flag.Parse()

	app.CheckMissing()
	app.CheckDate()
}

func (app *AppArg) CheckMissing() {
	if app.InternalPath.String() == "" {
		println("missing internal data path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if app.ExternalPath.String() == "" {
		println("missing external data path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if app.FromDate.String() == "" {
		println("missing from date")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if app.ToDate.String() == "" {
		println("missing to date")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func (app *AppArg) CheckDate() {
	if app.FromDate.String() > app.ToDate.String() {
		println("from date should before or same with to date")
		os.Exit(1)
	}
}

func main() {
	var appArg AppArg
	appArg.Scan()

	fmt.Println("Reconcile.")
}
