// Package log is a log package used by TKE team.
package log

import (
	"context"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/nine-monsters/go-hutool/hutool-log/klog"
)

// InfoLogger represents the ability to log non-error messages, at a particular verbosity.
type InfoLogger interface {
	// Info logs a non-error message with the given key/value pairs as context.
	//
	// The msg argument should be used to add some constant description to
	// the log line.  The key/value pairs can then be used to add additional
	// variable information.  The key/value pairs should alternate string
	// keys and arbitrary values.
	Info(msg string, keysAndValues ...interface{})
	Infof(format string, args ...interface{})

	// Enabled tests whether this InfoLogger is enabled.  For example,
	// commandline flags might be used to set the logging verbosity and disable
	// some info logs.
	Enabled() bool
}

// Logger represents the ability to log messages, both errors and not.
type Logger interface {
	// InfoLogger All Loggers implement InfoLogger.  Calling InfoLogger methods directly on
	// a Logger value is equivalent to calling them on a V(0) InfoLogger.  For
	// example, logger.Info() produces the same result as logger.V(0).Info.
	InfoLogger

	// Error logs an error, with the given message and key/value pairs as context.
	// It functions similarly to calling Info with the "error" named value, but may
	// have unique behavior, and should be preferred for logging errors (see the
	// package documentations for more information).
	//
	// The msg field should be used to add context to any underlying error,
	// while the err field should be used to attach the actual error that
	// triggered this log line, if present.
	Error(err error, msg string, keysAndValues ...interface{})
	Errorf(format string, args ...interface{})

	// V returns an InfoLogger value for a specific verbosity level.  A higher
	// verbosity level means a log message is less important.  It's illegal to
	// pass a log level less than zero.
	V(level int) InfoLogger

	Write(p []byte) (n int, err error)

	// WithValues adds some key-value pairs of context to a logger.
	// See Info for documentation on how key/value pairs work.
	WithValues(keysAndValues ...interface{}) Logger

	// WithName adds a new element to the logger's name.
	// Successive calls with WithName continue to append
	// suffixes to the logger's name.  It's strongly recommended
	// that name segments contain only letters, digits, and hyphens
	// (see the package documentation for more information).
	WithName(name string) Logger

	// WithContext returns a copy of context in which the log value is set.
	WithContext(ctx context.Context) context.Context

	// Flush calls the underlying Core's Sync method, flushing any buffered
	// log entries. Applications should take care to call Sync before exiting.
	Flush()
}

var _ Logger = &zapLogger{}

// noopInfoLogger is a logr.InfoLogger that's always disabled, and does nothing.
type noopInfoLogger struct{}

func (l *noopInfoLogger) Enabled() bool                    { return false }
func (l *noopInfoLogger) Info(_ string, _ ...interface{})  {}
func (l *noopInfoLogger) Infof(_ string, _ ...interface{}) {}

var disabledInfoLogger = &noopInfoLogger{}

// NB: right now, we always use the equivalent of sugared logging.
// This is necessary, since logr doesn't define non-suggared types,
// and using zap-specific non-suggared types would make uses tied
// directly to Zap.

// infoLogger is a logr.InfoLogger that uses Zap to log at a particular
// level.  The level has already been converted to a Zap level, which
// is to say that `logrLevel = -1*zapLevel`.
type infoLogger struct {
	level zapcore.Level
	log   *zap.Logger
}

func (l *infoLogger) Enabled() bool { return true }
func (l *infoLogger) Info(msg string, keysAndValues ...interface{}) {
	if checkedEntry := l.log.Check(l.level, msg); checkedEntry != nil {
		checkedEntry.Write(handleFields(l.log, keysAndValues)...)
	}
}

func (l *infoLogger) Infof(format string, args ...interface{}) {
	l.log.Sugar().Infof(format, args...)
}

// zapLogger is a logr.Logger that uses Zap to log.
type zapLogger struct {
	// NB: this looks very similar to zap.SugaredLogger, but
	// deals with our desire to have multiple verbosity levels.
	zapLogger *zap.Logger
	infoLogger
}

// handleFields converts a bunch of arbitrary key-value pairs into Zap fields.  It takes
// additional pre-converted Zap fields, for use with automatically attached fields, like
// `error`.
func handleFields(l *zap.Logger, args []interface{}, additional ...zap.Field) []zap.Field {
	// a slightly modified version of zap.SugaredLogger.sweetenFields
	if len(args) == 0 {
		// fast-return if we have no suggared fields.
		return additional
	}

	// unlike Zap, we can be pretty sure users aren't passing structured
	// fields (since logr has no concept of that), so guess that we need a
	// little less space.
	fields := make([]zap.Field, 0, len(args)/2+len(additional))
	for i := 0; i < len(args); {
		// check just in case for strongly-typed Zap fields, which is illegal (since
		// it breaks implementation agnosticism), so we can give a better error message.
		if _, ok := args[i].(zap.Field); ok {
			l.DPanic("strongly-typed Zap Field passed to logr", zap.Any("zap field", args[i]))
			break
		}

		// make sure this isn't a mismatched key
		if i == len(args)-1 {
			l.DPanic("odd number of arguments passed as key-value pairs for logging", zap.Any("ignored key", args[i]))
			break
		}

		// process a key-value pair,
		// ensuring that the key is a string
		key, val := args[i], args[i+1]
		keyStr, isString := key.(string)
		if !isString {
			// if the key isn't a string, DPanic and stop logging
			l.DPanic("non-string key argument passed to logging, ignoring all later arguments", zap.Any("invalid key", key))
			break
		}

		fields = append(fields, zap.Any(keyStr, val))
		i += 2
	}

	return append(fields, additional...)
}

var (
	logger  *zapLogger
	options *Options
)

// nolint: gochecknoinits // need to init a default logger
func init() {
	Init(NewOptions())
}

// Init initializes logger by opts which can custmoized by command arguments.
func Init(opts *Options) {
	options = opts
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: milliSecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	// when output to local path, with color is forbidden
	if opts.EnableColor {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = InfoLevel
	}

	loggerConfig := &zap.Config{
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Development:       false,
		DisableCaller:     !opts.EnableCaller,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         opts.Format,
		EncoderConfig:    encoderConfig,
		OutputPaths:      opts.OutputPaths,
		ErrorOutputPaths: opts.ErrorOutputPaths,
	}

	var err error
	l, err := loggerConfig.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	logger = &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
	klog.InitLogger(l)
	zap.RedirectStdLog(l)
}

// StdErrLogger returns logger of standard library which writes to supplied zap
// logger at error level.
func StdErrLogger() *log.Logger {
	if logger == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(logger.zapLogger, zapcore.ErrorLevel); err == nil {
		return l
	}
	return nil
}

// StdInfoLogger returns logger of standard library which writes to supplied zap
// logger at info level.
func StdInfoLogger() *log.Logger {
	if logger == nil {
		return nil
	}
	if l, err := zap.NewStdLogAt(logger.zapLogger, zapcore.InfoLevel); err == nil {
		return l
	}
	return nil
}

func (l *zapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if checkedEntry := l.zapLogger.Check(zap.ErrorLevel, msg); checkedEntry != nil {
		checkedEntry.Write(handleFields(l.zapLogger, keysAndValues, zap.Error(err))...)
	}
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.log.Sugar().Errorf(format, args...)
}

// V return a leveled InfoLogger.
func V(level int) InfoLogger { return logger.V(level) }
func (l *zapLogger) V(level int) InfoLogger {
	if level < 0 || level > 1 {
		panic("valid log level is [0, 1]")
	}
	lvl := zapcore.Level(-1 * level)
	if l.zapLogger.Core().Enabled(lvl) {
		return &infoLogger{
			level: lvl,
			log:   l.zapLogger,
		}
	}
	return disabledInfoLogger
}

func (l *zapLogger) Write(p []byte) (n int, err error) {
	l.zapLogger.Info(string(p))
	return len(p), nil
}

// WithValues creates a child logger and adds adds Zap fields to it.
func WithValues(keysAndValues ...interface{}) Logger { return logger.WithValues(keysAndValues...) }
func (l *zapLogger) WithValues(keysAndValues ...interface{}) Logger {
	newLogger := l.zapLogger.With(handleFields(l.zapLogger, keysAndValues)...)
	return NewLogger(newLogger)
}

// WithName adds a new path segment to the logger's name. Segments are joined by
// periods. By default, Loggers are unnamed.
func WithName(s string) Logger { return logger.WithName(s) }
func (l *zapLogger) WithName(name string) Logger {
	newLogger := l.zapLogger.Named(name)
	return NewLogger(newLogger)
}

// Flush calls the underlying Core's Sync method, flushing any buffered
// log entries. Applications should take care to call Sync before exiting.
func Flush() { logger.Flush() }
func (l *zapLogger) Flush() {
	_ = l.zapLogger.Sync()
}

// NewLogger creates a new logr.Logger using the given Zap Logger to log.
func NewLogger(l *zap.Logger) Logger {
	return &zapLogger{
		zapLogger: l,
		infoLogger: infoLogger{
			log:   l,
			level: zap.InfoLevel,
		},
	}
}

// ZapLogger used for other log wrapper such as klog.
func ZapLogger() *zap.Logger {
	return logger.zapLogger
}

// CheckIntLevel used for other log wrapper such as klog which return if logging a
// message at the specified level is enabled.
func CheckIntLevel(level int32) bool {
	var lvl zapcore.Level
	if level < 5 {
		lvl = zapcore.InfoLevel
	} else {
		lvl = zapcore.DebugLevel
	}
	checkEntry := logger.zapLogger.Check(lvl, "")
	return checkEntry != nil
}

// Debug method output debug level log.
func Debug(msg string, fields ...Field) {
	logger.zapLogger.Debug(msg, fields...)
}

// Debugf method output debug level log.
func Debugf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Debugf(format, v...)
}

// Debugw method output debug level log.
func Debugw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Debugw(msg, keysAndValues...)
}

// Info method output info level log.
func Info(msg string, fields ...Field) {
	logger.zapLogger.Info(msg, fields...)
}

// Infof method output info level log.
func Infof(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Infof(format, v...)
}

// Infow method output info level log.
func Infow(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Infow(msg, keysAndValues...)
}

// Warn method output warning level log.
func Warn(msg string, fields ...Field) {
	logger.zapLogger.Warn(msg, fields...)
}

// Warnf method output warning level log.
func Warnf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Warnf(format, v...)
}

// Warnw method output warning level log.
func Warnw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Warnw(msg, keysAndValues...)
}

// Error method output error level log.
func Error(msg string, fields ...Field) {
	logger.zapLogger.Error(msg, fields...)
}

// Errorf method output error level log.
func Errorf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Errorf(format, v...)
}

// Errorw method output error level log.
func Errorw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Errorw(msg, keysAndValues...)
}

// Panic method output panic level log and shutdown application.
func Panic(msg string, fields ...Field) {
	logger.zapLogger.Panic(msg, fields...)
}

// Panicf method output panic level log and shutdown application.
func Panicf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Panicf(format, v...)
}

// Panicw method output panic level log.
func Panicw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Panicw(msg, keysAndValues...)
}

// Fatal method output fatal level log.
func Fatal(msg string, fields ...Field) {
	logger.zapLogger.Fatal(msg, fields...)
}

// Fatalf method output fatal level log.
func Fatalf(format string, v ...interface{}) {
	logger.zapLogger.Sugar().Fatalf(format, v...)
}

// Fatalw method output Fatalw level log.
func Fatalw(msg string, keysAndValues ...interface{}) {
	logger.zapLogger.Sugar().Fatalw(msg, keysAndValues...)
}

func GetOptions() *Options {
	return options
}

func GetLogger() *zapLogger {
	return logger
}
