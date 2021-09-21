package log

import (
	"fmt"
	"io"
	"os"
)

type Config struct {
	Color     bool
	FatalFunc func()
	Out       io.Writer
}

type logger struct {
	cfg *Config
}

type Logger interface {
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Errorf(msg string, args ...interface{})
	Fatalf(msg string, args ...interface{})
	Logf(level Level, msg string, args ...interface{})
}

func SetDefault(l Logger) {
	defaultLogger = l
}

func New(cfg *Config) Logger {
	return &logger{cfg: cfg}
}

var defaultLogger Logger

var DefaultConfig = Config{
	Color: true,
	FatalFunc: func() {
		os.Exit(1)
	},
	Out: os.Stdout,
}

var ErrUnableToParseUnknownLevel = fmt.Errorf("tried to parse unknown level")

func init() {
	defaultLogger = &logger{
		&DefaultConfig,
	}
}

func (l *logger) format(lvl Level, f string, args ...interface{}) string {
	msg := fmt.Sprintf(f, args...)
	prefix := fmt.Sprintf("[%s]", lvl.String())

	if l.cfg.Color {
		prefix = lvl.Colorize()
	}

	return fmt.Sprintf("%s\t  %s\n", prefix, msg)
}

// Logf logs a message via the default logger at the provided Level using the fmt.Sprintf formatter
func Logf(lvl Level, f string, args ...interface{}) {
	defaultLogger.Logf(lvl, f, args...)
}

// Logf logs a message at the provided Level using the fmt.Sprintf formatter
func (l *logger) Logf(lvl Level, f string, args ...interface{}) {
	msg := l.format(lvl, f, args...)
	fmt.Fprintf(l.cfg.Out, msg)
}

// Infof logs a message via the default logger at Info Level using the fmt.Sprintf formatter
func Infof(f string, args ...interface{}) {
	defaultLogger.Infof(f, args...)
}

// Infof logs a message at Info Level using the fmt.Sprintf formatter
func (l *logger) Infof(f string, args ...interface{}) {
	l.Logf(Info, f, args...)
}

// Warnf logs a message via the default logger at Info Level using the fmt.Sprintf formatter
func Warnf(f string, args ...interface{}) {
	defaultLogger.Warnf(f, args...)
}

// Warnf logs a message at Info Level using the fmt.Sprintf formatter
func (l *logger) Warnf(f string, args ...interface{}) {
	l.Logf(Warn, f, args...)
}

// Debugf logs a message via the default logger at Info Level using the fmt.Sprintf formatter
func Debugf(f string, args ...interface{}) {
	defaultLogger.Debugf(f, args...)
}

// Debugf logs a message at Info Level using the fmt.Sprintf formatter
func (l *logger) Debugf(f string, args ...interface{}) {
	l.Logf(Debug, f, args...)
}

// Errorf logs a message via the default logger at Info Level using the fmt.Sprintf formatter
func Errorf(f string, args ...interface{}) {
	defaultLogger.Errorf(f, args...)
}

// Errorf logs a message at Info Level using the fmt.Sprintf formatter
func (l *logger) Errorf(f string, args ...interface{}) {
	l.Logf(Error, f, args...)
}

// Fatalf logs a message via the default logger at Info Level using the fmt.Sprintf formatter
func Fatalf(f string, args ...interface{}) {
	defaultLogger.Fatalf(f, args...)
}

// Fatalf logs a message at Info Level using the fmt.Sprintf formatter
func (l *logger) Fatalf(f string, args ...interface{}) {
	l.Logf(FATAL, f, args...)
	l.cfg.FatalFunc()
}
