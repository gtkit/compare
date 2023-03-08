// @Author xiaozhaofu 2022/10/31 19:51:00
package test

import (
	"testing"

	"github.com/xzf-tools/compare/version"
)

func TestVersion(t *testing.T) {
	versionA := "3.0.16"
	versionB := "3.0.15.6"
	t.Log("比较结果: ", version.Compare(versionA, versionB))

}
