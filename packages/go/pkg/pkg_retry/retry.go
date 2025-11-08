package pkgretry

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
)

func Backoff(attempt int, baseDelay time.Duration) time.Duration {
	// Calculate exponential delay: baseDelay * 2^(attempt-1)
	delay := baseDelay * time.Duration(1<<uint(attempt-1))

	// Add jitter: up to 10% of the delay
	jitter := time.Duration(rand.Int63n(int64(delay) / 10))
	return delay + jitter
}

func WithRetry(ctx context.Context, operation func() error, maxRetries int, baseDelay time.Duration) error {
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		err = operation()
		if err == nil {
			return nil
		}
		if pkgerr.IsRetryableError(err) {
			log.Printf("Attempt %d/%d: Retryable error, retrying: %v", attempt, maxRetries, err)
			time.Sleep(Backoff(attempt, baseDelay))
			continue
		}
		return pkgerr.NewInternalError("WithRetry", fmt.Sprintf("non-retryable error after %d attempts", attempt), err)
	}
	return pkgerr.NewInternalError("WithRetry", fmt.Sprintf("failed after %d attempts", maxRetries), err)
}
