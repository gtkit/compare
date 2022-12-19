// @Author xiaozhaofu 2022/10/31 19:51:00
package test

import (
	"testing"

	"gitlab.superjq.com/go-tools/compare/strcmp"
)

func TestStrCompare(t *testing.T) {
	a := "10.0"
	b := "6.1"
	r := strcmp.NumCmp(a, b)
	t.Log(r)
}
