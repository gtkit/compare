// @Author xiaozhaofu 2022/10/31 19:51:00
package test

import (
	"testing"

	"gitlab.superjq.com/go-tools/compare/strcmp"
)

func TestStrCompare(t *testing.T) {
	a := "19"
	b := "20"
	r := strcmp.NumCmp(a, b)
	t.Log(r)
}
