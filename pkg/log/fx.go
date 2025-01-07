package log

import (
	"go.uber.org/fx/fxevent"
)

type FxLogger interface {
	Logger
	LogEvent(event fxevent.Event)
}

type fxLoggerImpl struct {
	Logger
}

func NewFxLogger(logger Logger) FxLogger {
	return &fxLoggerImpl{
		Logger: logger.WithModule("fx"),
	}
}

func (l *fxLoggerImpl) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Debugf("OnStart executing: function=%s caller=%s", e.FunctionName, e.CallerName)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Errorf("OnStart failed: function=%s caller=%s error=%v", e.FunctionName, e.CallerName, e.Err)
		} else {
			l.Debugf("OnStart executed: function=%s caller=%s", e.FunctionName, e.CallerName)
		}
	case *fxevent.OnStopExecuting:
		l.Debugf("OnStop executing: function=%s caller=%s", e.FunctionName, e.CallerName)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Errorf("OnStop failed: function=%s caller=%s error=%v", e.FunctionName, e.CallerName, e.Err)
		} else {
			l.Debugf("OnStop executed: function=%s caller=%s", e.FunctionName, e.CallerName)
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Errorf("Supply failed: type=%s error=%v", e.TypeName, e.Err)
		} else {
			l.Debugf("Supplied: type=%s", e.TypeName)
		}
	case *fxevent.Provided:
		var moduleName string
		if e.ModuleName == "" {
			moduleName = "init"
		} else {
			moduleName = e.ModuleName
		}
		if e.Err != nil {
			l.Errorf("Provide failed: module=%s type=%v error=%v", moduleName, e.OutputTypeNames, e.Err)
		} else {
			l.Debugf("Provided: module=%s type=%v", moduleName, e.OutputTypeNames)
		}
	}
}
