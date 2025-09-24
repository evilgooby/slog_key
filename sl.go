package slog_key

import (
	"log/slog"
	"runtime/debug"
	"time"
)

const (
	errorKey     = "error"
	componentKey = "component"
	methodKey    = "method"
	elapsedKey   = "elapsed" // duration since
	panicKey     = "panic"
	stackKey     = "stack"
	sourceKey    = "source"
	valueKey     = "value"
)

// Error returns an error attribute; nil is ignored by handlers (empty key).
func Error(err error) slog.Attr {
	if err == nil {
		return slog.Attr{} // no-op, slog handler skips empty key
	}
	return slog.Any(errorKey, err)
}

// Component returns component attr, e.g. service/module name.
func Component(component string) slog.Attr {
	return slog.String(componentKey, component)
}

// Method returns method/function/operation attr.
func Method(method string) slog.Attr {
	return slog.String(methodKey, method)
}

// Elapsed logs duration since 'since' as a proper time.Duration value.
// Use for timings/latency: slog.Duration("elapsed", time.Since(since)).
func Elapsed(since time.Time) slog.Attr {
	return slog.Duration(elapsedKey, time.Since(since))
}

// Panic groups panic value + stack trace for easier filtering in JSON logs.
func Panic(p any) slog.Attr {
	return slog.Group(panicKey,
		slog.Any(valueKey, p),
		slog.String(stackKey, string(debug.Stack())),
	)
}

// Source groups component & method neatly under one key.
func Source(component, method string) slog.Attr {
	return slog.Group(sourceKey,
		slog.String(componentKey, component),
		slog.String(methodKey, method),
	)
}

// Field is a tiny helper to create arbitrary fields without importing slog in callers.
func Field(key string, value any) slog.Attr {
	return slog.Any(key, value)
}
