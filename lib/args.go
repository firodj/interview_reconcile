package lib

import (
	"fmt"
	"os"
	"time"
)

type ArgPathValue struct {
	val string
}

type ArgDateValue struct {
	val *time.Time
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

func (v ArgDateValue) String() string {
	if v.val != nil {
		return v.val.Format("2006-01-02")
	}
	return ""
}

func (v *ArgDateValue) Set(s string) error {
	if t, err := time.Parse("2006-01-02", s); err != nil {
		return fmt.Errorf("Date %s invalid should be YYYY-MM-DD", s)
	} else {
		v.val = new(time.Time)
		*v.val = t
	}

	return nil
}
