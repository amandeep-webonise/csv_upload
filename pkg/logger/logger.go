package logger

import (
	"path/filepath"
	"runtime"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Ilogger interface implements logger methods
type Ilogger interface {
	Initialise()
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
}

// RealLogger contains uber zap logger
type RealLogger struct {
	log *zap.SugaredLogger
}

// Initialise initializes logger with configurations
func (al *RealLogger) Initialise() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableCaller = true
	config.DisableStacktrace = true

	logger, err := config.Build()
	if err != nil {
		panic(err)
	}
	al.log = logger.Sugar()
}

// Info logs a message at level Info on the standard logger
func (al *RealLogger) Info(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Info(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Info(args)
	}
}

// Debug logs a message at level Debug on the standard logger
func (al *RealLogger) Debug(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Debug(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Debug(args)
	}
}

// Error logs a message at level Error on the standard logger with format specified
func (al *RealLogger) Error(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Error(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Error(args)
	}
}

// Panic logs a message at level Error on the standard logger with format specified
func (al *RealLogger) Panic(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Panic(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Panic(args)
	}
}
