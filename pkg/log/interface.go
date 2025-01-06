package log

// Fields type for structured logging
type Fields map[string]interface{}

type Logger interface {
	Info(i ...interface{})
	Warn(i ...interface{})
	Error(i ...interface{})
	Debug(i ...interface{})

	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Debugf(format string, args ...interface{})

	WithModule(module string) Logger
	WithFields(fields Fields) Logger
}

type FormatterType string

const (
	TextFormatter FormatterType = "text"
	JSONFormatter FormatterType = "json"
)

type Options struct {
	FormatterType FormatterType
	FilePath      string
	Level         Level
}
