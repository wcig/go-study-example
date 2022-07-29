package log

import (
	"encoding/json"
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

// 日志第二版本: 新增调用文件和行号记录
type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return fmt.Sprintf("Level(%d)", l)
	}
}

type Logger struct {
	Writer     io.Writer
	Level      Level
	PathPrefix string
}

func NewLogger(writer io.Writer, level Level, pathPrefix string) *Logger {
	logger := &Logger{
		Writer:     writer,
		Level:      level,
		PathPrefix: pathPrefix,
	}
	return logger
}

type LogMsg struct {
	Time   string
	Level  string
	Msg    string
	Caller string
}

func (l *Logger) Debug(args ...interface{}) {
	l.log(DebugLevel, "", args)
}

func (l *Logger) Info(args ...interface{}) {
	l.log(InfoLevel, "", args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, "", args)
}

func (l *Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, "", args)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.log(DebugLevel, template, args)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.log(InfoLevel, template, args)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.log(WarnLevel, template, args)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.log(ErrorLevel, template, args)
}

func (l *Logger) log(level Level, template string, args []interface{}) {
	if level < l.Level {
		return
	}

	var msg string
	if len(args) > 0 {
		if template == "" {
			msg = fmt.Sprint(args...)
		} else {
			msg = fmt.Sprintf(template, args...)
		}
	}

	logMsg := &LogMsg{
		Time:   ISO8601TimeEncoder(time.Now()),
		Level:  level.String(),
		Msg:    msg,
		Caller: l.getCaller(),
	}
	if line, err := json.Marshal(logMsg); err == nil {
		line = append(line, '\n')
		_, _ = l.Writer.Write(line)
	}
}

func (l *Logger) getCaller() string {
	_, file, line, ok := runtime.Caller(3)
	if ok {
		caller := fmt.Sprintf("%s:%d", file, line)
		if l.PathPrefix != "" {
			if start := strings.Index(caller, l.PathPrefix); start > 0 {
				return caller[start:]
			}
		}
		return caller
	}
	return ""
}

func ISO8601TimeEncoder(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z0700")
}
