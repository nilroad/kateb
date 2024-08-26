package kateb

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"
)

var logger Logger

// ANSI color codes for terminal output
const (
	colorReset         = "\033[0m"
	colorBlack         = "\033[30m"
	colorRed           = "\033[31m"
	colorGreen         = "\033[32m"
	colorYellow        = "\033[33m"
	colorBlue          = "\033[34m"
	colorMagenta       = "\033[35m"
	colorCyan          = "\033[36m"
	colorWhite         = "\033[37m"
	colorBrightRed     = "\033[91m"
	colorBrightGreen   = "\033[92m"
	colorBrightYellow  = "\033[93m"
	colorBrightBlue    = "\033[94m"
	colorBrightMagenta = "\033[95m"
	colorBrightCyan    = "\033[96m"
	colorBrightWhite   = "\033[97m"
)

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
	Colorize  bool
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
func (r *Logger) setColor(level slog.Level) *Logger {

	if r.config.Colorize {
		switch level {
		case slog.LevelDebug:
			fmt.Println(colorCyan)
		case slog.LevelInfo:
			fmt.Println(colorWhite)
		case slog.LevelWarn:
			fmt.Println(colorBrightYellow)
		case slog.LevelError:
			fmt.Println(colorRed)
		default:
			fmt.Println(colorReset)
		}
	}

	return r
}
func (r *Logger) restColor() *Logger {
	fmt.Println(colorReset)

	return r
}
func (r *Logger) Error(message string, args map[string]any) {

	r.setColor(slog.LevelError).sl.Error(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
}
func (r *Logger) Fatal(message string, args map[string]any) {
	r.setColor(slog.LevelError).sl.Error(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
	os.Exit(1)
}
func (r *Logger) Info(message string, args map[string]any) {
	r.setColor(slog.LevelInfo).sl.Info(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
}
func (r *Logger) Debug(message string, args map[string]any) {
	r.setColor(slog.LevelDebug).sl.Debug(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
}
func (r *Logger) Warn(message string, args map[string]any) {
	r.setColor(slog.LevelWarn).sl.Warn(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
}
func (r *Logger) Panic(message string, args map[string]any) {
	r.setColor(slog.LevelError).sl.Error(r.config.Prefix+":"+message, contextKey, args)
	r.restColor()
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
