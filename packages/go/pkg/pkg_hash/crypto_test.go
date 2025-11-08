package pkghash_test

import (
	"testing"

	pkghash "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_hash"
)

func TestGetHash(t *testing.T) {
	t.Parallel()
	key := "testKey"
	expectedHash := "15291f67d99ea7bc578c3544dadfbb991e66fa69cb36ff70fe30e798e111ff5f"
	hash := pkghash.GetHash(key)
	if hash != expectedHash {
		t.Errorf("Expected hash %s, got %s", expectedHash, hash)
	}
}
