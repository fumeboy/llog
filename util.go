package llog

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
)

type LogData struct {
	curTime     time.Time
	message     string
	timeStr     string
	level       LogLevel
	filename    string
	lineNo      int
	fields      []*KeyVal
}

func writeField(buffer *bytes.Buffer, field, sep string) {
	buffer.WriteString(field)
	buffer.WriteString(sep)
}

func (l *LogData) Bytes() []byte {

	var buffer bytes.Buffer
	levelStr := getLevelText(l.level)

	writeField(&buffer, l.timeStr, SpaceSep)
	writeField(&buffer, levelStr, SpaceSep)
	writeField(&buffer, l.filename, ColonSep)
	writeField(&buffer, fmt.Sprintf("%d", l.lineNo), SpaceSep)

	for _, field := range lm.globalFields {
		writeField(&buffer, fmt.Sprintf("%v=%v", field.key, field.val), SpaceSep)
	}

	if l.level == LogLevelAccess && l.fields != nil {
		for _, field := range l.fields {
			writeField(&buffer, fmt.Sprintf("%v=%v", field.key, field.val), SpaceSep)
		}
	}
	writeField(&buffer, l.message, LineSep)
	return buffer.Bytes()
}

//util.go 10
func GetLineInfo() (fileName string, lineNo int) {
	_, fileName, lineNo, _ = runtime.Caller(3)
	return
}
