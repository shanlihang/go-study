package main

import (
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	slog.Info("Hello world")
	slog.Warn("Hello world")
	slog.Error("Hello world")
	slog.Debug("Hello world")

}