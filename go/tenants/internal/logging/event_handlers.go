package logging

import (
	"app/internal/ddd"
	"context"

	"github.com/rs/zerolog"
)

type EventHandlers[T ddd.Event] struct {
	ddd.EventHandler[T]
	label  string
	logger zerolog.Logger
}

var _ ddd.EventHandler[ddd.Event] = (*EventHandlers[ddd.Event])(nil)

func LogEventHandlerAccess[T ddd.Event](handlers ddd.EventHandler[T], label string, logger zerolog.Logger) EventHandlers[T] {
	return EventHandlers[T]{
		EventHandler: handlers,
		label:        label,
		logger:       logger,
	}
}

func (h EventHandlers[T]) HandleEvent(ctx context.Context, event T) (err error) {
	h.logger.Info().Msgf("--> Stores.%s.On(%s)", h.label, event.EventName())
	defer func() { h.logger.Info().Err(err).Msgf("<-- Stores.%s.On(%s)", h.label, event.EventName()) }()
	return h.EventHandler.HandleEvent(ctx, event)
}
