package logger

import (
	"fmt"
	"os"
	"path"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerFactory interface {
	NewLogger() (*zap.Logger, func(), error)
}

type loggerFactory struct {
	logPath string
}

func NewLoggerFactory(logPath string) LoggerFactory {
	return &loggerFactory{
		logPath: logPath,
	}
}

func (lf *loggerFactory) NewLogger() (*zap.Logger, func(), error) {
	now := time.Now()
	logfile := path.Join(lf.logPath, fmt.Sprintf("%s.log", now.Format("2006-01-02")))

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, err
	}

	pe := zap.NewProductionEncoderConfig()

	fileEncoder := zapcore.NewJSONEncoder(pe)
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, zapcore.AddSync(file), highPriority),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), highPriority),
	)

	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.WarnLevel))
	close := func() {
		log.Sync()
		file.Close()
	}

	return log, close, nil
}

func LoggerInit() {
	// สร้าง log factory
	logFactory := NewLoggerFactory("./logger")

	logger, loggerClose, err := logFactory.NewLogger()
	if err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}
	defer loggerClose()

	// Error log จะถูกบันทึกลงไฟล์ และออกไปที่ console
	logger.Error("Some error",
		zap.Error(fmt.Errorf("error")),
		zap.String("key", "value"))

	// Info log จะไม่ถูกบันทึกและแสดงผลใด ๆ
	logger.Info("Some data",
		zap.Error(fmt.Errorf("data")),
		zap.String("key", "value"))
}
