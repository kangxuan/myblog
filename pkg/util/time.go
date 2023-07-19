package util

import "time"

// TimeToString 时间戳转换成字符串
func TimeToString(unix int) string {
	tmr := time.Unix(int64(unix), 0)
	return tmr.Format("2006-01-02 03:04:05")
}
