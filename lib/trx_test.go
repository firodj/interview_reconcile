package lib

import "testing"

func TestAja(t *testing.T) {

	err := InternalReader("../fixtures/sys1.csv", "2020-05-24", "2020-05-24")

	if err != nil {
		t.Fatalf("Error %v", err)
	}

	err = ExternalReader("../fixtures/bank1.csv", "2020-05-24", "2020-05-24")

	if err != nil {
		t.Fatalf("Error %v", err)
	}
}
