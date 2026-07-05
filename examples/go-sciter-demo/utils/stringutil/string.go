package stringutil

import "strings"

// input: 待处理字符串 separator：分隔符
func GetLastSplitItem(input string, separator string) string {
	lines := strings.Split(input, separator)
	size := len(lines)
	if size == 0 {
		return ""
	}
	return lines[size-1]
}

func GetSplitItemN(input string, separator string, n int) string {
	if n < 0 {
		return ""
	}
	lines := strings.Split(input, separator)
	size := len(lines)
	if size == 0 || n > size {
		return ""
	}
	return lines[size-1]
}

func Trim(input string) string {
	return strings.Trim(input, " ")
}
