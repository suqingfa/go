package main

import "log/slog"

// #nosec
func add(a, b int) int

func main() {
	slog.Info("Hello World", "add(1, 2)", add(1, 2))
}
