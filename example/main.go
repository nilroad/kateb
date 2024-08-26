package main

import (
	"github.com/nilroad/kateb"
	"log/slog"
	"os"
)

func main() {

	kc := kateb.Config{
		Level:     slog.LevelDebug,
		AddSource: true,
		Prefix:    "example",
	}
	log := kateb.New(os.Stdout, kc)
	log.Info("Application started", map[string]any{"version": "0.0.0"})
	log.Warn("Warning issued", nil)
	log.Error("An error occurred", map[string]any{"error": "something went wrong"})
}
