package pkgstring_test

import (
	"testing"

	pkgstring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

func TestOrString(t *testing.T) {
	t.Parallel()

	t.Run("All empty strings", func(t *testing.T) {
		result := pkgstring.OrString("", "", "")
		expected := ""
		if result != expected {
			t.Errorf("OrString(%q, %q, %q) = %q, want %q", "", "", "", result, expected)
		}
	})

	t.Run("First non-empty string", func(t *testing.T) {
		result := pkgstring.OrString("first", "", "")
		expected := "first"
		if result != expected {
			t.Errorf("OrString(%q, %q, %q) = %q, want %q", "first", "", "", result, expected)
		}
	})

	t.Run("Last non-empty string", func(t *testing.T) {
		result := pkgstring.OrString("", "", "last")
		expected := "last"
		if result != expected {
			t.Errorf("OrString(%q, %q, %q) = %q, want %q", "", "", "last", result, expected)
		}
	})

	t.Run("Multiple non-empty strings", func(t *testing.T) {
		result := pkgstring.OrString("", "middle", "last")
		expected := "middle"
		if result != expected {
			t.Errorf("OrString(%q, %q, %q) = %q, want %q", "", "middle", "last", result, expected)
		}
	})
}
