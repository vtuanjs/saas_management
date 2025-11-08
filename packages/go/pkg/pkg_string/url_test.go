package pkgstring_test

import (
	"testing"

	ustring "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_string"
)

func TestIsURL(t *testing.T) {
	t.Parallel()

	t.Run("Valid URL with http", func(t *testing.T) {
		if got := ustring.IsURL("http://example.com"); !got {
			t.Errorf("IsURL(%q) = %v, want %v", "http://example.com", got, true)
		}
	})

	t.Run("Valid URL with https", func(t *testing.T) {
		if got := ustring.IsURL("https://example.com"); !got {
			t.Errorf("IsURL(%q) = %v, want %v", "https://example.com", got, true)
		}
	})

	t.Run("Invalid URL without scheme", func(t *testing.T) {
		if got := ustring.IsURL("example.com"); got {
			t.Errorf("IsURL(%q) = %v, want %v", "example.com", got, false)
		}
	})

	t.Run("Invalid URL with empty string", func(t *testing.T) {
		if got := ustring.IsURL(""); got {
			t.Errorf("IsURL(%q) = %v, want %v", "", got, false)
		}
	})

	t.Run("Invalid URL with only scheme", func(t *testing.T) {
		if got := ustring.IsURL("http://"); got {
			t.Errorf("IsURL(%q) = %v, want %v", "http://", got, false)
		}
	})
}
