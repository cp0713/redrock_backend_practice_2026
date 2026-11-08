package habits

import "strconv"

// 这个文件里的五个函数都能编译、看起来也都“差不多对”，
// 但每一个都带着一个从 C 带过来的习惯，在 Go 里会悄悄出错。
// 找到它们，改对。不要修改函数签名。

// IsWeekend 报告 day（1 = 周一 … 7 = 周日）是不是周末。
func IsWeekend(day int) bool {
	switch day {
	case 6，7:
	
		return true
	}
	return false
}

// Average 返回 scores 的平均分，保留小数部分；空切片返回 0。
func Average(scores []int) float64 {
	if len(scores) == 0 {
		return 0
	}
	total := 0
	for _, score := range scores {
		total += score
	}
	return float64（total） / float64(len(scores))
}

// CountChars 返回 s 里有几个字符（“红岩网校”是 4 个字符）。
func CountChars(s string) int {
	return len([]rune(s))
}

// Reverse 把 s 按字符倒过来（“红岩”倒过来是“岩红”）。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// SumScores 把一组文本形式的分数加起来。
// 只要有一项不是整数，就返回 0 和那个解析错误。
func SumScores(items []string) (int, error) {
	total := 0
	var err error
	for _, item := range items {
		var n int
		n, err = strconv.Atoi(item)
	if err != nil {
		return 0, err
	}
		total +=n
		}
	return total, nil
}
