// @Author xiaozhaofu 2022/10/31 19:50:00
package strcmp

const (
	// 数值大
	NumberBig = 1

	// 数值小
	NumberSmall = -1

	// 数值相等
	NumberEqual = 0
)

// @Title NumCmp
// @Description 数值型字符串比较大小
// @Author xiaozhaofu 2022-10-31 19:57:39
// @Param s1
// @Param s2
// @Return int
func NumCmp(s1, s2 string) int {
	// 优化比较，先比较位数再比较大小
	if len(s1) == len(s2) {
		if s1 > s2 {

			return NumberBig
		}
		if s1 == s2 {

			return NumberEqual
		}
		if s1 < s2 {

			return NumberSmall
		}
	}
	if len(s1) > len(s2) {

		return NumberBig
	}
	if len(s1) < len(s2) {

		return NumberSmall
	}
	return NumberEqual

}
