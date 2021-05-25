package ch26_log

import (
	"os"
	"testing"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zap log sample use
// zap.SugaredLogger, zap.Logger区别: 1.性能差点,2.支持interface类型,3.zap.Logger.Sugar()可获取zap.SugaredLogger
func TestZapLogSampleUse(t *testing.T) {
	exLogger := zap.NewExample()
	devLogger, _ := zap.NewDevelopment()
	prodLogger, _ := zap.NewProduction()

	exLogger.Info("ok")
	exLogger.Sugar().Info("ok")
	devLogger.Info("ok")
	devLogger.Sugar().Info("ok")
	prodLogger.Info("ok")
	prodLogger.Sugar().Info("ok")
}

// output:
// {"level":"info","msg":"ok"}
// {"level":"info","msg":"ok"}
// 2020-12-09T17:43:06.502+0800	INFO	ch20_log/zap_log_test.go:17	ok
// 2020-12-09T17:43:06.502+0800	INFO	ch20_log/zap_log_test.go:18	ok
// {"level":"info","ts":1607506986.5025754,"caller":"ch20_log/zap_log_test.go:19","msg":"ok"}
// {"level":"info","ts":1607506986.502586,"caller":"ch20_log/zap_log_test.go:20","msg":"ok"}

// 定时zap log1 (json格式日志)
func TestZapLog1(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	enc := zapcore.NewJSONEncoder(encoderCfg)
	file, err := os.OpenFile("zap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
	if err != nil {
		t.Fatal(err)
	}
	ws := zapcore.AddSync(file)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	logger := zap.New(core).Sugar()
	logger.Info("ok") // {"level":"info","ts":1607507841.5085142,"msg":"ok"}
}

// 定时zap log2 (console格式日志)
func TestZapLog2(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	enc := zapcore.NewConsoleEncoder(encoderCfg)
	file, err := os.OpenFile("zap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
	if err != nil {
		t.Fatal(err)
	}
	ws := zapcore.AddSync(file)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	logger := zap.New(core).Sugar()
	logger.Info("ok") // 1.607507860892628e+09	info	ok
}

// 定时zap log3 (json格式日志 + 时间格式 + 大写日志级别 + 打印函数和堆栈信息)
func TestZapLog3(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder   // 时间格式
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder // 大写日志级别
	encoderCfg.TimeKey = "time"                          // 设置时间字段key
	encoderCfg.CallerKey = "call"                        // 设置调用key
	encoderCfg.FunctionKey = "callFunc"                  // 设置调用函数key
	encoderCfg.StacktraceKey = "stack"                   // 设置堆栈key
	enc := zapcore.NewJSONEncoder(encoderCfg)
	file, err := os.OpenFile("zap.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0775)
	if err != nil {
		t.Fatal(err)
	}
	ws := zapcore.AddSync(file)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	// 增加打印函数行号: zap.AddCaller()
	// 增加堆栈信息: zap.AddStacktrace(zapcore.ErrorLevel)
	// 跳过指定级别的堆栈信息: zap.AddCallerSkip(skip)
	srcLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	logger := srcLogger.Sugar()
	logger.Info("ok")
	logger.Error("err")
}

// 定制zap log4 (增加自定义参数)
func TestZapLog4(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	enc := zapcore.NewJSONEncoder(encoderCfg)
	ws := zapcore.AddSync(os.Stdout)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	fields := []zap.Field{
		zap.String("name", "service-001"),
		zap.String("version", "001"),
	}
	addOption := zap.Fields(fields...)
	zLogger := zap.New(core, addOption)
	logger := zLogger.Sugar()
	logger.Info("ok")
}

// 定制zap log5 (写入日志到多个目标)
func TestZapLog5(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	enc := zapcore.NewJSONEncoder(encoderCfg)
	ws := zapcore.NewMultiWriteSyncer(os.Stdout, os.Stderr)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	fields := []zap.Field{
		zap.String("name", "service-001"),
		zap.String("version", "001"),
	}
	addOption := zap.Fields(fields...)
	zLogger := zap.New(core, addOption)
	logger := zLogger.Sugar()
	logger.Info("ok")
}

// 使用Lumberjack日志归档
func TestZapLogLumberjack(t *testing.T) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encoderCfg)
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "zap.log", // 日志文件路径
		MaxSize:    500,       // 日志文件最大大小 (megabytes)
		MaxAge:     7,         // 最长保留时间 (days)
		MaxBackups: 30,        // 备份日志最多保留个数
		LocalTime:  false,     // 使用当前utc时间
		Compress:   false,     // 启用gzip压缩
	}
	ws := zapcore.AddSync(lumberjackLogger)
	level := zap.DebugLevel

	core := zapcore.NewCore(enc, ws, level)
	srcLogger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	logger := srcLogger.Sugar()
	logger.Info("ok") // {"level":"info","ts":1607570204.8173146,"msg":"ok","name":"service-001","version":"001"}
}

// zap log生产环境示例
func TestProdZapLog(t *testing.T) {
	logger := InitLogger("./zap.log", "debug")
	u := map[string]interface{}{
		"id":   1,
		"name": "tom",
	}
	logger.Debugf("debug log user:%+v", u)
	logger.Infof("info log user:%+v", u)
	logger.Warnf("warn log user:%+v", u)
	logger.Errorf("err log user:%+v", u)
}

func InitLogger(logPath, logLevel string) *zap.SugaredLogger {
	// 设置日志文件
	hook := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    500,
		MaxAge:     7,
		MaxBackups: 30,
		LocalTime:  false,
		Compress:   false,
	}
	w := zapcore.AddSync(hook)

	// 设置日志级别: debug->info->warn->error
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	// 设置日志编码: 时间格式
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 构建logger
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		level,
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	return logger.Sugar()
}
