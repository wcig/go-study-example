package log

import (
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	logger := NewLogger(os.Stdout, InfoLevel)
	logger.Debug("debug log: ", 111)
	logger.Debugf("debug log: %d", 222)
	logger.Info("info log: ", 333)
	logger.Infof("info log: %d", 444)
	logger.Warn("info log: ", 555)
	logger.Error("info log: ", 666)
	// Output:
	// {"Time":"2022-07-29T15:54:20.263+0800","Level":"info","Msg":"info log: 333"}
	// {"Time":"2022-07-29T15:54:20.263+0800","Level":"info","Msg":"info log: 444"}
	// {"Time":"2022-07-29T15:54:20.263+0800","Level":"warn","Msg":"info log: 555"}
	// {"Time":"2022-07-29T15:54:20.263+0800","Level":"error","Msg":"info log: 666"}
}
