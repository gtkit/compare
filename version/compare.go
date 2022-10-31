// @Author xiaozhaofu 2022/10/31 19:49:00
package version

import (
	"strings"
)

const (
	// 版本号大
	VersionBig = 1

	// 版本号小
	VersionSmall = -1

	// 版本号相等
	VersionEqual = 0
)

// @Title Compare
// @Description 版本号比较
// @Author xiaozhaofu 2022-10-31 19:58:52
// @Param verA
// @Param verB
// @Return int
func Compare(verA, verB string) int {

	verStrArrA := spliteStrByDot(verA)
	verStrArrB := spliteStrByDot(verB)

	lenStrA := len(verStrArrA)
	lenStrB := len(verStrArrB)

	if lenStrA != lenStrB {
		panic("版本号格式不一致!")

	}

	return compareArrVersion(verStrArrA, verStrArrB)
}

// 比较版本号字符串数组
func compareArrVersion(verA, verB []string) int {

	for i, _ := range verA {

		littleResult := compareLittleVersion(verA[i], verB[i])

		if littleResult != VersionEqual {
			return littleResult
		}
	}

	return VersionEqual
}

//
// 比较小版本号字符串
//
func compareLittleVersion(verA, verB string) int {

	bytesA := []byte(verA)
	bytesB := []byte(verB)

	lenA := len(bytesA)
	lenB := len(bytesB)
	if lenA > lenB {
		return VersionBig
	}

	if lenA < lenB {
		return VersionSmall
	}

	// 如果长度相等则按byte位进行比较

	return compareByBytes(bytesA, bytesB)
}

// 按byte位进行比较小版本号
func compareByBytes(verA, verB []byte) int {

	for index, _ := range verA {
		if verA[index] > verB[index] {
			return VersionBig
		}
		if verA[index] < verB[index] {
			return VersionSmall
		}

	}

	return VersionEqual
}

// 按“.”分割版本号为小版本号的字符串数组
func spliteStrByDot(strV string) []string {

	return strings.Split(strV, ".")
}
