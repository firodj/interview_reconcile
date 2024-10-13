package lib

import (
	"testing"
)

func TestArgPathValueSet(t *testing.T) {
	a := ArgPathValue{}
	a_lib := "../lib"

	err := a.Set(a_lib)
	if err != nil {
		t.Fatalf("ArgPathValue.Set %s should exist, %v", a_lib, err)
	}

	if a.String() != a_lib {
		t.Fatalf("ArgPathValue.String mismatch %s.", a.String())
	}
}

func TestArgDateValueSet(t *testing.T) {
	a := ArgDateValue{}
	a_date := "2024-04-01"

	err := a.Set(a_date)
	if err != nil {
		t.Fatalf("ArgDateValue.Set %s should success, %v", a_date, err)
	}

	if a.String() != a_date {
		t.Fatalf("ArgDateValue.String mismatch %s.", a.String())
	}
}
