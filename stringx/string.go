package stringx

const (
	defaultLimitSuffix = "..."
)

// LimitRune 按照 rune 为单位，截取 s 字符串前 n 位
// 如果 s 比 n 短，直接返回 s
// 如果 s 比 n 长，则默认在后面补充 ...
func LimitRune(s string, n int, suffix ...string) string {
	list := []rune(s)

	if len(list) <= n {
		return s
	}

	suffixStr := defaultLimitSuffix
	if len(suffix) > 0 {
		suffixStr = suffix[0]
	}

	return string(list[0:n]) + suffixStr
}
