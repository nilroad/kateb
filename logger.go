package kateb

import (
	"io"
	"log/slog"
	"os"
	"time"
)

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
