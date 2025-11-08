package pkgchunk_test

import (
	"reflect"
	"testing"

	pkgchunk "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_chunk"
)

func TestSplitChunk(t *testing.T) {
	t.Parallel()

	t.Run("should split into chunks of size 2", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		chunks := pkgchunk.Split(items, 2)
		if len(chunks) != 3 {
			t.Errorf("expected 3 chunks, got %d", len(chunks))
		}
		expected := [][]int{{1, 2}, {3, 4}, {5}}
		for i := range expected {
			if !reflect.DeepEqual(chunks[i], expected[i]) {
				t.Errorf("expected chunk %d to be %v, got %v", i, expected[i], chunks[i])
			}
		}
	})

	t.Run("should return one chunk when limit is greater than length", func(t *testing.T) {
		items := []int{1, 2}
		chunks := pkgchunk.Split(items, 5)
		if len(chunks) != 1 {
			t.Errorf("expected 1 chunk, got %d", len(chunks))
		}
		if !reflect.DeepEqual(chunks[0], items) {
			t.Errorf("expected chunk to be %v, got %v", items, chunks[0])
		}
	})

	t.Run("should return one chunk when limit is zero", func(t *testing.T) {
		items := []int{1, 2}
		chunks := pkgchunk.Split(items, 0)
		if len(chunks) != 1 {
			t.Errorf("expected 1 chunk, got %d", len(chunks))
		}
		if !reflect.DeepEqual(chunks[0], items) {
			t.Errorf("expected chunk to be %v, got %v", items, chunks[0])
		}
	})

	t.Run("should return 0 chunk when empty", func(t *testing.T) {
		items := []int{}
		chunks := pkgchunk.Split(items, 500)
		if len(chunks) != 0 {
			t.Errorf("expected 0 chunks, got %d", len(chunks))
		}
	})
}
