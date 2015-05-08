package logger

var (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FAIL
)

// 日志等级
var levelInfo map[int]interface{} = make(map[int]interface{}, 0, 5)

func init() {
	levelInfo[DEBUG] = "DEBUG"
	levelInfo[INFO] = "INFO"
	levelInfo[WARN] = "WARN"
	levelInfo[ERROR] = "ERROR"
	levelInfo[FAIL] = "FAIL"
}
