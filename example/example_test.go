package example

import (
	"github.com/fumeboy/llog"
	"testing"
)

func TestFileLogger(t *testing.T) {
	defer llog.Close()
	L.StartWithLevel(LogLevelAccess)
	L.Info("info")
	L.Debug("debug")
	L.Access("access")
	L.Trace("trace")
	L.Warn("warn")
	L.Error("error")
}
