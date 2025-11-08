//go:generate sh -c "mockgen -source=$GOFILE -destination=mock/$(basename $GOFILE .go)_test.go -package=mock"
package core

import "context"

type Logger interface {
	Debug(ctx context.Context, msg string, data ...any)
	Info(ctx context.Context, msg string, data ...any)
	Warn(ctx context.Context, msg string, data ...any)
	Error(ctx context.Context, msg string, err error)
	Fatal(ctx context.Context, msg string, err error)
}
