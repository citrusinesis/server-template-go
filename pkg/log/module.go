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

func WithOptions(opts *Options) fx.Option {
	if opts == nil {
		opts = DefaultOptions()
	}

	return fx.Module("log",
		fx.Supply(opts),
		fx.Provide(
			NewLogger,
			NewFxLogger,
		),
	)
}

var Module = fx.Module("log",
	fx.Provide(
		DefaultOptions,
		NewLogger,
		NewFxLogger,
	),
)
