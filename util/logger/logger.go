package logger

type ILogger interface {
	Debug()
	Info()
	Warn()
	Error()
	Fail()
}

type ILoggerFactory interface {
	logger()
}

// 日志工厂
type LoggerFactory struct {
}

func (this *LoggerFactory) logger() *Logger {
	return &Logger{}
}

func NewLogger() {
	return loggerFactory.logger()
}

func (l *LoggerFactory) logger() {
	return &Logger{}
}

// 日志对象
type Logger struct {
}

func (this *Logger) Debug() {

}

func (this *Logger) Info() {

}

func (this *Logger) Warn() {

}

func (this *Logger) Error() {

}

func (this *Logger) Fail() {

}

// 单例
var loggerFactory LoggerFactory = &LoggerFactory{}
