package util

import (
	"fmt"
	"testing"
)

func TestRandUA(t *testing.T) {
	for i := 0; i < 5; i++ {
		t.Log(RandUA())
	}
}
func TestCompressStr(t *testing.T) {
	s:=CompressStr(`
	113.121.78.181		`)
	fmt.Println(s)
}