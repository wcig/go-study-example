package log

import (
	"os"
	"testing"
)

func TestLogNoPathPrefix(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel, "")
	logger.Debug("debug log: ", 111)
	logger.Info("info log: ", 222)
	logger.Warn("warn log: ", 333)
	logger.Error("error log: ", 444)
	// Output:
	// {"Time":"2022-07-29T16:11:44.565+0800","Level":"debug","Msg":"debug log: 111","Caller":"/Users/yangbo/Documents/myproject/go-study-example/study/log/second/log_test.go:10"}
	// {"Time":"2022-07-29T16:11:44.565+0800","Level":"info","Msg":"info log: 222","Caller":"/Users/yangbo/Documents/myproject/go-study-example/study/log/second/log_test.go:11"}
	// {"Time":"2022-07-29T16:11:44.565+0800","Level":"warn","Msg":"warn log: 333","Caller":"/Users/yangbo/Documents/myproject/go-study-example/study/log/second/log_test.go:12"}
	// {"Time":"2022-07-29T16:11:44.565+0800","Level":"error","Msg":"error log: 444","Caller":"/Users/yangbo/Documents/myproject/go-study-example/study/log/second/log_test.go:13"}
}

func TestLogWithPathPrefix(t *testing.T) {
	logger := NewLogger(os.Stdout, DebugLevel, "go-study-example")
	logger.Debug("debug log: ", 111)
	logger.Info("info log: ", 222)
	logger.Warn("warn log: ", 333)
	logger.Error("error log: ", 444)
	// Output:
	// {"Time":"2022-07-29T16:18:21.684+0800","Level":"debug","Msg":"debug log: 111","Caller":"go-study-example/study/log/second/log_test.go:23"}
	// {"Time":"2022-07-29T16:18:21.684+0800","Level":"info","Msg":"info log: 222","Caller":"go-study-example/study/log/second/log_test.go:24"}
	// {"Time":"2022-07-29T16:18:21.684+0800","Level":"warn","Msg":"warn log: 333","Caller":"go-study-example/study/log/second/log_test.go:25"}
	// {"Time":"2022-07-29T16:18:21.684+0800","Level":"error","Msg":"error log: 444","Caller":"go-study-example/study/log/second/log_test.go:26"}
}
