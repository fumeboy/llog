package llog

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	outputer, err := NewFileOutputer("./logs/test.log")
	if err != nil {
		t.Errorf("init file outputer failed, err:%v", err)
		return
	}

	InitLogger(LogLevelDebug, 10000,&KeyVal{"service_name","apple"})
	AddOutputer(outputer)
	
	Debug("this is a good test")
	Trace("this is a good test")
	Info("this is a good test")
	Access("this is a good test")
	Warn("this is a good test")
	Error("this is a good test")
	Stop()
}