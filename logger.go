package llog

import (
	"fmt"
	"path"
	"time"
)

func (this *CTX) Debug(format string, args ...interface{}) {
	this.writeLog(LogLevelDebug, format, args...)
}

func (this *CTX) Trace(format string, args ...interface{}) {
	this.writeLog(LogLevelTrace, format, args...)
}

func (this *CTX) Access(format string, args ...interface{}) {
	this.writeLog(LogLevelAccess, format, args...)
}

func (this *CTX) Info(format string, args ...interface{}) {
	this.writeLog(LogLevelInfo, format, args...)
}

func (this *CTX) Warn(format string, args ...interface{}) {
	this.writeLog(LogLevelWarn, format, args...)
}

func (this *CTX) Error(format string, args ...interface{}) {
	this.writeLog(LogLevelError, format, args...)
}

func (ctx *CTX) writeLog(level LogLevel, format string, args ...interface{}) {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")

	fileName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		message:  msg,
		curTime:  now,
		timeStr:  nowStr,
		level:    level,
		filename: fileName,
		lineNo:   lineNo,
		fields:   ctx.Fields,
	}

	select {
	case lm.logDataChan <- logData:
	default:
		return
	}
}

func Debug(format string, args ...interface{}) {
	writeLog(LogLevelDebug, format, args...)
}

func Trace(format string, args ...interface{}) {
	writeLog(LogLevelTrace, format, args...)
}

func Access(format string, args ...interface{}) {
	writeLog(LogLevelAccess, format, args...)
}

func Info(format string, args ...interface{}) {
	writeLog(LogLevelInfo, format, args...)
}

func Warn(format string, args ...interface{}) {
	writeLog(LogLevelWarn, format, args...)
}

func Error(format string, args ...interface{}) {
	writeLog(LogLevelError, format, args...)
}

func writeLog(level LogLevel, format string, args ...interface{}) {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")

	fileName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		message:  msg,
		curTime:  now,
		timeStr:  nowStr,
		level:    level,
		filename: fileName,
		lineNo:   lineNo,
	}

	select {
	case lm.logDataChan <- logData:
	default:
		return
	}
}
