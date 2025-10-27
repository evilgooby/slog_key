// Package sl расширяет стандартные атрибуты slog (slog.String, slog.Int итд.)
// и предоставляет удобные конструкторы для ошибок, паник и компонентов.
package sl

import (
	"log/slog"
	"runtime/debug"
	"time"
)

const (
	keyError     = "error"
	keyPanic     = "panic"
	keyComponent = "component"
	keyDuration  = "duration"
)

// Error возвращает атрибут с текстом ошибки.
// Можно использовать в любом уровне логирования (slog.LevelDebug, slog.LevelInfo, slog.LevelWarn, slog.LevelError).
//
//	logger.Error("cannot connect", sl.Error(err))
//	logger.Info("incorrect http request", sl.Error(err))
func Error(err error) slog.Attr {
	return slog.String(keyError, err.Error())
}

// Panic возвращает группу атрибутов с сообщением паники и стеком вызова.
// Подходит для логирования в recover-блоках.
//
//	defer func() {
//		if r := recover(); r != nil {
//			logger.Error("panic recovered", sl.Panic(r))
//		}
//	}()
func Panic(recovered any) slog.Attr {
	return slog.Group(keyPanic,
		slog.Any("message", recovered),
		slog.Any("stack", string(debug.Stack())),
	)
}

// Component возвращает атрибут с именем компонента.
//
//	logger := slog.With(sl.Component("integrations.grpc.SomeService"))
//	logger.Info("start request")
func Component(component string) slog.Attr {
	return slog.String(keyComponent, component)
}

// Duration возвращает атрибут с временем выполнения.
//
//	start := time.Now()
//	end := time.Since(start)
//	logger.Info("done", sl.Duration(end))
func Duration(t time.Duration) slog.Attr {
	return slog.Duration(keyDuration, t)
}

// Since возвращает slog.Attr с длительностью,
// прошедшей с момента t0. Удобно использовать в defer для логирования времени выполнения.
//
//	start := time.Now()
//	defer logger.Info("done", sl.Since(start))
func Since(t0 time.Time) slog.Attr {
	return slog.Duration(keyDuration, time.Since(t0))
}
