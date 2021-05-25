package ch26_log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

// log使用
func TestLogUse(t *testing.T) {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("message")
	log.Fatalln("fatal message") // Fatalln调用Println后调用os.Exit(1)退出程序，所以下一步不会被执行
	log.Panicln("panic message") // Panicln调用Println后调用panic
}

// 定制log
func TestMyLog(t *testing.T) {
	var (
		traceLog   *log.Logger
		infoLog    *log.Logger
		warningLog *log.Logger
		errorLog   *log.Logger
	)

	file, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal("failed to open error log file:", err)
	}

	traceLog = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	infoLog = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warningLog = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorLog = log.New(io.MultiWriter(os.Stdout, file),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	traceLog.Println("trace msg")
	infoLog.Println("info msg")
	warningLog.Println("warning msg")
	errorLog.Println("error msg")
}

// output:
// INFO: 2020/09/27 22:00:22 log_test.go:52: info msg
// WARNING: 2020/09/27 22:00:22 log_test.go:53: warning msg
// ERROR: 2020/09/27 22:00:22 log_test.go:54: error msg
