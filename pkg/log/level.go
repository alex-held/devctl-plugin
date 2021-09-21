package log

import (
	"fmt"
)

const (
	DebugColor = "\033[1;30m[\033[1;32m%s\033[1;30m]\033[0m"
	InfoColor  = "\033[1;30m[\033[1;37m%s\033[1;30m]\033[0m"
	WarnColor  = "\033[1;30m[\033[1;33m%s\033[1;30m]\033[0m"
	ErrorColor = "\033[1;30m[\033[1;31m%s\033[1;30m]\033[0m"
	FatalColor = "\033[1;30m[\033[1;35m%s\033[1;30m]\033[0m"
)

var (
	DebugPrefix = fmt.Sprintf(DebugColor, "DEBUG")
	InfoPrefix  = fmt.Sprintf(InfoColor, "INFO")
	WarnPrefix  = fmt.Sprintf(WarnColor, "WARN")
	ErrorPrefix = fmt.Sprintf(ErrorColor, "ERROR")
	FatalPrefix = fmt.Sprintf(FatalColor, "FATAL")
)

type Level string

const (
	Info  Level = "INFO"
	Debug Level = "DEBUG"
	Warn  Level = "WARN"
	Error Level = "ERROR"
	FATAL Level = "FATAL"
)

func ParseLevel(lvlStr string) (lvl Level, err error) {
	switch lvlStr {
	case Debug.String():
		return Debug, nil
	case Info.String():
		return Info, nil
	case Warn.String():
		return Warn, nil
	case Error.String():
		return Error, nil
	case FATAL.String():
		return FATAL, nil
	default:
		return Info, ErrUnableToParseUnknownLevel
	}
}

func (l Level) String() string {
	return string(l)
}

func (l Level) Colorize() string {
	switch l {
	case Debug:
		return DebugPrefix
	case Warn:
		return WarnPrefix
	case Error:
		return ErrorPrefix
	case FATAL:
		return FatalPrefix
	case Info:
		fallthrough
	default:
		return InfoPrefix
	}
}
