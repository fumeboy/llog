package example

import (
	"fmt"
	"github.com/fumeboy/llog"
	"path"
	"runtime"
	"time"
)

type exampleLog struct {
	level LogLevel
	consolePrinter *llog.Printer
	filePrinter *llog.Printer
}

var L = &exampleLog{}

func (l *exampleLog) StartWithLevel(level LogLevel){
	l.level = level
	l.filePrinter,_  = llog.Use(&PrinterFile{})
	l.consolePrinter,_  = llog.Use(&PrinterConsole{})
}

func (l *exampleLog) write(msg string, color Color) {
	now := time.Now().Format("2006-01-02 15:04:05.999")
	fileName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	text := fmt.Sprintf("%s %s %d %s\n", now, fileName, lineNo, msg)

	l.filePrinter.Write([]byte(text))
	l.consolePrinter.Write([]byte(color.contain(text)))
}

func (l *exampleLog) Info(msg string) {
	if l.level <= LogLevelInfo {
		l.write(msg, Green)
	}
}

func (l *exampleLog) Debug(msg string) {
	if l.level <= LogLevelDebug {
		l.write(msg, Magenta)
	}
}

func (l *exampleLog) Access(msg string) {
	if l.level <= LogLevelAccess {
		l.write(msg, Blue)
	}
}

func (l *exampleLog) Warn(msg string) {
	if l.level <= LogLevelWarn {
		l.write(msg, Cyan)
	}
}

func (l *exampleLog) Error(msg string) {
	if l.level <= LogLevelError {
		l.write(msg, Red)
	}
}

func (l *exampleLog) Trace(msg string) {
	if l.level <= LogLevelTrace {
		l.write(msg, Yellow)
	}
}

func GetLineInfo() (fileName string, lineNo int) {
	_, fileName, lineNo, _ = runtime.Caller(3)
	return
}