package slogdiscard

import (
	"context"
	"log/slog"
)

func NewDiscardLogger() *slog.Logger {
	return slog.New(NewDiscardHandler())
}

type DiscardHandler struct{}

func NewDiscardHandler() *DiscardHandler {
	return &DiscardHandler{}
}

func (h *DiscardHandler) Handle(_ context.Context, _ slog.Record) error {
	// logging is ignored
	return nil
}

func (h *DiscardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	// returns the same handler since there are no attributes to save
	return h
}

func (h *DiscardHandler) WithGroup(_ string) slog.Handler {
	// returns the same handler since there are no group to save
	return h
}

func (h *DiscardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// always returns false because the logging is ignored
	return false
}
