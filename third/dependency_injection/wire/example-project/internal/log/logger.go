package log

import (
	"go-app/third/dependency_injection/wire/example-project/internal/config"
	"io"
	"os"
	"path/filepath"

	"github.com/google/wire"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logFileName = "app.log"
)

var ProviderSet = wire.NewSet(NewLogger)

func NewLogger(c *config.Logger) *zap.SugaredLogger {
	var (
		w    io.Writer
		file *os.File
		err  error
	)

	if c.Path != "" {
		if err = MakeDir(c.Path); err != nil {
			panic(err)
		}
		filePath := filepath.Join(c.Path, logFileName)
		file, err = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		w = io.MultiWriter(file, os.Stdout)
	} else {
		w = os.Stdout
	}

	var logLevel zapcore.Level
	switch c.Level {
	case "debug":
		logLevel = zap.DebugLevel
	case "info":
		logLevel = zap.InfoLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	default:
		logLevel = zap.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 构建logger
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(w),
		logLevel,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger.Sugar()
}

func MakeDir(f string) error {
	if IsFileExist(f) {
		return nil
	}
	return os.MkdirAll(f, os.ModePerm)
}

func IsFileExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
