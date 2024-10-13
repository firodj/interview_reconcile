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
