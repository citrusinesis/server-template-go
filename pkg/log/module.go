package log

import (
	"go.uber.org/fx"
)

func NewLogger(opts *Options) Logger {
	return newLogrusLogger(opts)
}

func DefaultOptions() *Options {
	return &Options{
		FormatterType: TextFormatter,
		FilePath:      "",
		Level:         DebugLevel,
	}
}

var Module = fx.Module("log",
	fx.Provide(
		DefaultOptions,
		NewLogger,
		NewFxLogger,
	),
)
