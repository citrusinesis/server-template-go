package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

type logrusLogger struct {
	logger *logrus.Logger
	module string
	fields logrus.Fields
}

func toLogrusLevel(level Level) logrus.Level {
	switch level {
	case DebugLevel:
		return logrus.DebugLevel
	case InfoLevel:
		return logrus.InfoLevel
	case WarnLevel:
		return logrus.WarnLevel
	case ErrorLevel:
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

func newLogrusLogger(opts *Options) Logger {
	logger := logrus.New()
	logger.SetLevel(toLogrusLevel(opts.Level))

	switch opts.FormatterType {
	case JSONFormatter:
		logger.SetFormatter(&logrus.JSONFormatter{})
	default:
		logger.SetFormatter(&logrus.TextFormatter{})
	}

	if opts.FilePath != "" {
		file, err := os.OpenFile(opts.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			logger.SetOutput(file)
		}
	}

	return &logrusLogger{
		logger: logger,
		module: "",
		fields: logrus.Fields{},
	}
}

func (l *logrusLogger) WithModule(module string) Logger {
	return &logrusLogger{
		logger: l.logger,
		module: module,
		fields: l.fields,
	}
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogger{
		logger: l.logger,
		module: l.module,
		fields: logrus.Fields(fields),
	}
}

func (l *logrusLogger) getFields() logrus.Fields {
	fields := logrus.Fields{}
	// Merge existing fields
	for k, v := range l.fields {
		fields[k] = v
	}
	// Add module if set
	if l.module != "" {
		fields["module"] = l.module
	}
	return fields
}

func (l *logrusLogger) Info(i ...interface{}) {
	l.logger.WithFields(l.getFields()).Info(i...)
}

func (l *logrusLogger) Warn(i ...interface{}) {
	l.logger.WithFields(l.getFields()).Warn(i...)
}

func (l *logrusLogger) Error(i ...interface{}) {
	l.logger.WithFields(l.getFields()).Error(i...)
}

func (l *logrusLogger) Debug(i ...interface{}) {
	l.logger.WithFields(l.getFields()).Debug(i...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.WithFields(l.getFields()).Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.WithFields(l.getFields()).Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.WithFields(l.getFields()).Errorf(format, args...)
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.WithFields(l.getFields()).Debugf(format, args...)
}
