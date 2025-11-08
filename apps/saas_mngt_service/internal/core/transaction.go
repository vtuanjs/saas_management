package core

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
	"time"

	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgmust "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_must"
)

// txKey is a custom type for context key to avoid collisions.
type (
	DBTxKey  struct{}
	DBClient interface {
		Tx(ctx context.Context) (DBTransaction, error)
	}
	DBTransaction = interface {
		Commit() error
		Rollback() error
		Client() DBClient
	}
)

// setTxToContext stores the ent transaction in the context.
func setTxToContext(ctx context.Context, tx DBTransaction) context.Context {
	return context.WithValue(ctx, DBTxKey{}, tx)
}

func execute(ctx context.Context, tx DBTransaction, fn func(newCtx context.Context) error) error {
	defer func() {
		if v := recover(); v != nil {
			err := tx.Rollback()
			if err != nil {
				pkgmust.Check("Failed to rollback transaction after panic", err)
			}
			panic(v)
		}
	}()

	if err := fn(ctx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = pkgerr.NewInternalError("execute", "Failed to rollback transaction", rerr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return pkgerr.NewInternalError("execute", "Failed to commit transaction", err)
	}

	return nil
}

// isTransientError checks if the error is a transient PostgreSQL error that can be retried.
func isTransientError(err error) bool {
	if err == nil {
		return false
	}

	var e *pkgerr.Error
	if errors.As(err, &e) {
		if e.Code == pkgerr.CodeTransactionError { // Serialization failure
			return true
		}
	}

	// General connection errors
	if strings.Contains(err.Error(), "connection refused") ||
		strings.Contains(err.Error(), "network error") ||
		strings.Contains(err.Error(), "dial tcp") ||
		strings.Contains(err.Error(), "i/o timeout") {
		return true
	}

	return false
}

// Execute a function within a database transaction with retry logic for transient PostgreSQL errors.
func WrapTransaction(ctx context.Context, client DBClient, fn func(newCtx context.Context) error) error {
	var lastErr error
	const maxRetries = 5
	const baseDelay = 100 * time.Millisecond // Base delay for exponential backoff

	txFromCtx, ok := GetDBTx(ctx)
	if ok && txFromCtx != nil {
		// If a transaction already exists in the context, use it.
		return fn(ctx)
	}

	for attempt := 0; attempt <= maxRetries; attempt++ {
		// Otherwise, create a new transaction.
		tx, err := client.Tx(ctx)
		if err != nil {
			return err
		}

		// Store the transaction in the context.
		newCtx := setTxToContext(ctx, tx)

		// Attempt the transaction
		err = execute(newCtx, tx, fn)
		if err == nil {
			return nil // Success
		}

		// Store the last error
		lastErr = err

		// Check if the error is transient and retryable
		if !isTransientError(err) {
			return err // Non-transient error, return immediately
		}

		// If this is not the last attempt, wait with exponential backoff + jitter
		if attempt < maxRetries {
			delay := baseDelay * time.Duration(1<<uint(attempt)) // Exponential backoff: baseDelay * 2^attempt
			// Generate secure random jitter up to 50ms
			jitterMax := big.NewInt(50)
			jitterRand, _ := rand.Int(rand.Reader, jitterMax)
			jitter := time.Duration(jitterRand.Int64()) * time.Millisecond
			select {
			case <-ctx.Done():
				return pkgerr.NewCancelledError("WrapTransaction", "Transaction canceled", ctx.Err())
			case <-time.After(delay + jitter):
				// Continue to the next attempt
			}
		}
	}

	// All retries failed, return the last error
	return pkgerr.NewAbortedError("WrapTransaction", "Transaction failed after retries", lastErr)
}

// GetClient retrieves the ent client from the context, or the original client if no transaction is set.
func GetDBClient(ctx context.Context, client DBClient) DBClient {
	tx, ok := ctx.Value(DBTxKey{}).(DBTransaction)
	if !ok || tx == nil {
		return client
	}
	return tx.Client()
}

// GetTx retrieves the ent transaction from the context.
func GetDBTx(ctx context.Context) (DBTransaction, bool) {
	tx, ok := ctx.Value(DBTxKey{}).(DBTransaction)
	return tx, ok
}
