package util

// Utf8Substring 字符串切割
func Utf8Substring(s string, start, end int) string {
	// 将字符串转换为[]rune，以处理UTF-8编码的字符
	runes := []rune(s)

	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}

	return string(runes[start:end])
}
