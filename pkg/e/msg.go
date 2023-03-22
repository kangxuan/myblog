package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	INVALID_PARAMS: "参数错误",
}

// GetMsg 获取错误信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return "非法错误"
}
