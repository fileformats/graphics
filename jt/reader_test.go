package jt

import (
	"fmt"
	"testing"
)

func TestReader(t *testing.T) {
	jt, err := Load("missing")
	if jt != nil || err == nil {
		t.Fail()
	}
	jt, err = Load("./testdata/feder.jt")
	if err != nil {
		t.Fail()
	}
	fmt.Print("\n\n\n")
}
