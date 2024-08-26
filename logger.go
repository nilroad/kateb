package kateb

import (
	"io"
	"log/slog"
	"os"
	"time"
)

var logger Logger

func init() {
	logger.sl = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == "msg" {
				return slog.Attr{
					Key:   "message",
					Value: a.Value,
				}
			}
			if a.Key == "time" {
				return slog.Attr{
					Key:   "time",
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}

			return a
		},
	}))
}

// Set you can change the default logger
func Set(l Logger) {
	logger = l
}

const contextKey = "context"

type Config struct {
	Level     slog.Level
	AddSource bool
	Prefix    string
}

type Logger struct {
	sl     *slog.Logger
	config Config
}

func New(writer io.Writer, config Config) *Logger {
	return &Logger{
		sl: slog.New(slog.NewJSONHandler(writer, &slog.HandlerOptions{
			AddSource: config.AddSource,
			Level:     config.Level,
			ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
				if a.Key == "msg" {
					return slog.Attr{
						Key:   "message",
						Value: a.Value,
					}
				}
				if a.Key == "time" {
					return slog.Attr{
						Key:   "time",
						Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
					}
				}

				return a
			},
		})),
		config: config,
	}
}

func (r *Logger) Error(message string, args map[string]any) {
	r.sl.Error(r.config.Prefix+":"+message, contextKey, args)
}
func (r *Logger) Fatal(message string, args map[string]any) {
	r.sl.Error(r.config.Prefix+":"+message, contextKey, args)
	os.Exit(1)
}
func (r *Logger) Info(message string, args map[string]any) {
	r.sl.Info(r.config.Prefix+":"+message, contextKey, args)
}
func (r *Logger) Debug(message string, args map[string]any) {
	r.sl.Debug(r.config.Prefix+":"+message, contextKey, args)
}
func (r *Logger) Warn(message string, args map[string]any) {
	r.sl.Warn(r.config.Prefix+":"+message, contextKey, args)
}
func (r *Logger) Panic(message string, args map[string]any) {
	r.sl.Error(r.config.Prefix+":"+message, contextKey, args)
	panic(message)
}

func Error(message string, args map[string]any) {
	logger.sl.Error(message, contextKey, args)
}
func Fatal(message string, args map[string]any) {
	logger.sl.Error(message, contextKey, args)
	os.Exit(1)
}
func Info(message string, args map[string]any) {
	logger.sl.Info(message, contextKey, args)
}
func Debug(message string, args map[string]any) {
	logger.sl.Debug(message, contextKey, args)
}
func Warn(message string, args map[string]any) {
	logger.sl.Warn(message, contextKey, args)
}
func Panic(message string, args map[string]any) {
	logger.sl.Error(message, contextKey, args)
	panic(message)
}
