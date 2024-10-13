package lib

import (
	"fmt"
	"os"
)

type ArgPathValue struct {
	val string
}

func (v ArgPathValue) String() string {
	return v.val
}

func (v *ArgPathValue) Set(s string) error {
	if stat, err := os.Stat(s); os.IsNotExist(err) || !stat.IsDir() {
		return fmt.Errorf("Path %s not exists", s)
	} else {
		v.val = s
	}

	return nil
}
