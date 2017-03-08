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
	jt, err = Load("./testdata/test.jt")
	if err != nil {
		t.Fail()
	}
	data, err := jt.ToJSON(LOD0)
	fmt.Println("\n", string(data), err)
	fmt.Print("\n\n\n")
}
