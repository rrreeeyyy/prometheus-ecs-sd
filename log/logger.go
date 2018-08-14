package log

import (
	"io"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type Logger struct {
	Out log.Logger
	Err log.Logger
}

func NewLogger(out io.Writer, err io.Writer) Logger {
	outLogger := log.NewLogfmtLogger(log.NewSyncWriter(out))
	outLogger = level.NewFilter(outLogger, level.AllowNone(), level.AllowDebug(), level.AllowInfo())
	outLogger = log.With(outLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	errLogger := log.NewLogfmtLogger(log.NewSyncWriter(err))
	errLogger = level.NewFilter(errLogger, level.AllowWarn(), level.AllowError())
	errLogger = log.With(errLogger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)

	return Logger{
		Out: outLogger,
		Err: errLogger,
	}
}
