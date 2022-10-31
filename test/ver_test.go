// @Author xiaozhaofu 2022/10/31 19:51:00
package test

import (
	"testing"

	"gitlab.superjq.com/go-tools/compare/version"
)

func TestVersion(t *testing.T) {
	versionA := "3.0.15.6"
	versionB := "3.0.15.5"
	t.Log(version.Compare(versionA, versionB))

}
