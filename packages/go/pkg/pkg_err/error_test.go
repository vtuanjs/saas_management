package pkgerr_test

import (
	"testing"

	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
)

func TestIsValidationError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsValidationError success", func(t *testing.T) {
		err := pkgerr.NewValidationError("test_key", "test_message", nil)
		if !pkgerr.IsValidationError(err) {
			t.Errorf("Expected IsValidationError to return true")
		}
	})

	t.Run("TestIsValidationError failure", func(t *testing.T) {
		err := pkgerr.NewNotFoundError("test_key", "test_message", nil)
		if pkgerr.IsValidationError(err) {
			t.Errorf("Expected IsValidationError to return false")
		}
	})

	t.Run("TestIsValidationError with nil error", func(t *testing.T) {
		if pkgerr.IsValidationError(nil) {
			t.Errorf("Expected IsValidationError to return false for nil error")
		}
	})
}

func TestIsNotFoundError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsNotFoundError success", func(t *testing.T) {
		err := pkgerr.NewNotFoundError("test_key", "test_message", nil)
		if !pkgerr.IsNotFoundError(err) {
			t.Errorf("Expected IsNotFoundError to return true")
		}
	})

	t.Run("TestIsNotFoundError failure", func(t *testing.T) {
		err := pkgerr.NewValidationError("test_key", "test_message", nil)
		if pkgerr.IsNotFoundError(err) {
			t.Errorf("Expected IsNotFoundError to return false")
		}
	})

	t.Run("TestIsNotFoundError with nil error", func(t *testing.T) {
		if pkgerr.IsNotFoundError(nil) {
			t.Errorf("Expected IsNotFoundError to return false for nil error")
		}
	})
}

func TestIsUnauthorizedError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsUnauthorizedError success", func(t *testing.T) {
		err := pkgerr.NewUnauthorizedError("test_key", "test_message", nil)
		if !pkgerr.IsUnauthorizedError(err) {
			t.Errorf("Expected IsUnauthorizedError to return true")
		}
	})

	t.Run("TestIsUnauthorizedError failure", func(t *testing.T) {
		err := pkgerr.NewNotFoundError("test_key", "test_message", nil)
		if pkgerr.IsUnauthorizedError(err) {
			t.Errorf("Expected IsUnauthorizedError to return false")
		}
	})

	t.Run("TestIsUnauthorizedError with nil error", func(t *testing.T) {
		if pkgerr.IsUnauthorizedError(nil) {
			t.Errorf("Expected IsUnauthorizedError to return false for nil error")
		}
	})
}

func TestIsForbiddenError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsForbiddenError success", func(t *testing.T) {
		err := pkgerr.NewForbiddenError("test_key", "test_message", nil)
		if !pkgerr.IsForbiddenError(err) {
			t.Errorf("Expected IsForbiddenError to return true")
		}
	})

	t.Run("TestIsForbiddenError failure", func(t *testing.T) {
		err := pkgerr.NewUnauthorizedError("test_key", "test_message", nil)
		if pkgerr.IsForbiddenError(err) {
			t.Errorf("Expected IsForbiddenError to return false")
		}
	})

	t.Run("TestIsForbiddenError with nil error", func(t *testing.T) {
		if pkgerr.IsForbiddenError(nil) {
			t.Errorf("Expected IsForbiddenError to return false for nil error")
		}
	})
}

func TestIsDuplicateError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsDuplicateError success", func(t *testing.T) {
		err := pkgerr.NewDuplicateError("test_key", "test_message", nil)
		if !pkgerr.IsDuplicateError(err) {
			t.Errorf("Expected IsDuplicateError to return true")
		}
	})

	t.Run("TestIsDuplicateError failure", func(t *testing.T) {
		err := pkgerr.NewForbiddenError("test_key", "test_message", nil)
		if pkgerr.IsDuplicateError(err) {
			t.Errorf("Expected IsDuplicateError to return false")
		}
	})

	t.Run("TestIsDuplicateError with nil error", func(t *testing.T) {
		if pkgerr.IsDuplicateError(nil) {
			t.Errorf("Expected IsDuplicateError to return false for nil error")
		}
	})
}

func TestIsConflictError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsConflictError success", func(t *testing.T) {
		err := pkgerr.NewConflictError("test_key", "test_message", nil)
		if !pkgerr.IsConflictError(err) {
			t.Errorf("Expected IsConflictError to return true")
		}
	})

	t.Run("TestIsConflictError failure", func(t *testing.T) {
		err := pkgerr.NewDuplicateError("test_key", "test_message", nil)
		if pkgerr.IsConflictError(err) {
			t.Errorf("Expected IsConflictError to return false")
		}
	})

	t.Run("TestIsConflictError with nil error", func(t *testing.T) {
		if pkgerr.IsConflictError(nil) {
			t.Errorf("Expected IsConflictError to return false for nil error")
		}
	})
}

func TestIsTimeoutError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsTimeoutError success", func(t *testing.T) {
		err := pkgerr.NewTimeoutError("test_key", "test_message", nil)
		if !pkgerr.IsTimeoutError(err) {
			t.Errorf("Expected IsTimeoutError to return true")
		}
	})

	t.Run("TestIsTimeoutError failure", func(t *testing.T) {
		err := pkgerr.NewConflictError("test_key", "test_message", nil)
		if pkgerr.IsTimeoutError(err) {
			t.Errorf("Expected IsTimeoutError to return false")
		}
	})

	t.Run("TestIsTimeoutError with nil error", func(t *testing.T) {
		if pkgerr.IsTimeoutError(nil) {
			t.Errorf("Expected IsTimeoutError to return false for nil error")
		}
	})
}

func TestIsInternalError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsInternalError success", func(t *testing.T) {
		err := pkgerr.NewInternalError("test_key", "test_message", nil)
		if !pkgerr.IsInternalError(err) {
			t.Errorf("Expected IsInternalError to return true")
		}
	})

	t.Run("TestIsInternalError failure", func(t *testing.T) {
		err := pkgerr.NewTimeoutError("test_key", "test_message", nil)
		if pkgerr.IsInternalError(err) {
			t.Errorf("Expected IsInternalError to return false")
		}
	})

	t.Run("TestIsInternalError with nil error", func(t *testing.T) {
		if pkgerr.IsInternalError(nil) {
			t.Errorf("Expected IsInternalError to return false for nil error")
		}
	})
}

func TestIsTransactionError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsTransactionError success", func(t *testing.T) {
		err := pkgerr.NewTransactionError("test_key", "test_message", nil)
		if !pkgerr.IsTransactionError(err) {
			t.Errorf("Expected IsTransactionError to return true")
		}
	})

	t.Run("TestIsTransactionError failure", func(t *testing.T) {
		err := pkgerr.NewInternalError("test_key", "test_message", nil)
		if pkgerr.IsTransactionError(err) {
			t.Errorf("Expected IsTransactionError to return false")
		}
	})

	t.Run("TestIsTransactionError with nil error", func(t *testing.T) {
		if pkgerr.IsTransactionError(nil) {
			t.Errorf("Expected IsTransactionError to return false for nil error")
		}
	})
}

func TestIsCancelledError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsCancelledError success", func(t *testing.T) {
		err := pkgerr.NewCancelledError("test_key", "test_message", nil)
		if !pkgerr.IsCancelledError(err) {
			t.Errorf("Expected IsCancelledError to return true")
		}
	})

	t.Run("TestIsCancelledError failure", func(t *testing.T) {
		err := pkgerr.NewTransactionError("test_key", "test_message", nil)
		if pkgerr.IsCancelledError(err) {
			t.Errorf("Expected IsCancelledError to return false")
		}
	})

	t.Run("TestIsCancelledError with nil error", func(t *testing.T) {
		if pkgerr.IsCancelledError(nil) {
			t.Errorf("Expected IsCancelledError to return false for nil error")
		}
	})
}

func TestIsAbortedError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsAbortedError success", func(t *testing.T) {
		err := pkgerr.NewAbortedError("test_key", "test_message", nil)
		if !pkgerr.IsAbortedError(err) {
			t.Errorf("Expected IsAbortedError to return true")
		}
	})

	t.Run("TestIsAbortedError failure", func(t *testing.T) {
		err := pkgerr.NewCancelledError("test_key", "test_message", nil)
		if pkgerr.IsAbortedError(err) {
			t.Errorf("Expected IsAbortedError to return false")
		}
	})

	t.Run("TestIsAbortedError with nil error", func(t *testing.T) {
		if pkgerr.IsAbortedError(nil) {
			t.Errorf("Expected IsAbortedError to return false for nil error")
		}
	})
}

func TestIsRetryableError(t *testing.T) {
	t.Parallel()
	t.Run("TestIsRetryableError success", func(t *testing.T) {
		err := pkgerr.NewRetryableError("test_key", "test_message", nil)
		if !pkgerr.IsRetryableError(err) {
			t.Errorf("Expected IsRetryableError to return true")
		}
	})

	t.Run("TestIsRetryableError failure", func(t *testing.T) {
		err := pkgerr.NewAbortedError("test_key", "test_message", nil)
		if pkgerr.IsRetryableError(err) {
			t.Errorf("Expected IsRetryableError to return false")
		}
	})

	t.Run("TestIsRetryableError with nil error", func(t *testing.T) {
		if pkgerr.IsRetryableError(nil) {
			t.Errorf("Expected IsRetryableError to return false for nil error")
		}
	})
}

func TestIsCode(t *testing.T) {
	t.Parallel()
	t.Run("TestIsCode success", func(t *testing.T) {
		err := pkgerr.NewInternalError("test_key", "test_message", nil)
		if !pkgerr.IsCode(err, pkgerr.CodeInternalError) {
			t.Errorf("Expected IsCode to return true for CodeInternalError")
		}
	})

	t.Run("TestIsCode failure", func(t *testing.T) {
		err := pkgerr.NewValidationError("test_key", "test_message", nil)
		if pkgerr.IsCode(err, pkgerr.CodeInternalError) {
			t.Errorf("Expected IsCode to return false for CodeInternalError")
		}
	})

	t.Run("TestIsCode with nil error", func(t *testing.T) {
		if pkgerr.IsCode(nil, pkgerr.CodeInternalError) {
			t.Errorf("Expected IsCode to return false for nil error")
		}
	})
}

func TestIsKey(t *testing.T) {
	t.Parallel()
	t.Run("TestIsKey success", func(t *testing.T) {
		err := pkgerr.NewInternalError("test_key", "test_message", nil)
		if !pkgerr.IsKey(err, "test_key") {
			t.Errorf("Expected IsKey to return true for key 'test_key'")
		}
	})

	t.Run("TestIsKey failure", func(t *testing.T) {
		err := pkgerr.NewInternalError("another_key", "test_message", nil)
		if pkgerr.IsKey(err, "test_key") {
			t.Errorf("Expected IsKey to return false for key 'test_key'")
		}
	})

	t.Run("TestIsKey with nil error", func(t *testing.T) {
		if pkgerr.IsKey(nil, "test_key") {
			t.Errorf("Expected IsKey to return false for nil error")
		}
	})
}
