package ristretto

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	system "github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/core"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgmust "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_must"

	"github.com/dgraph-io/ristretto"
)

type ristrettoCache struct {
	cache *ristretto.Cache
}

var (
	singleton *ristrettoCache
	once      sync.Once
)

func NewRistrettoCache() system.LocalCache {
	once.Do(func() {
		cache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e7,     // number of keys to track frequency of (10M).
			MaxCost:     1 << 30, // maximum cost of cache (1GB).
			BufferItems: 64,      // number of keys per Get buffer.
		})
		if err != nil {
			pkgmust.Check("Failed to create Ristretto cache", err)
			return
		}

		singleton = &ristrettoCache{cache: cache}
	})
	return singleton
}

func (rc *ristrettoCache) Get(_ context.Context, key string) (interface{}, bool) {
	return rc.cache.Get(key)
}

func (rc *ristrettoCache) Set(_ context.Context, key string, value interface{}) bool {
	return rc.cache.Set(key, value, 1) // Cost mặc định = 1
}

func (rc *ristrettoCache) SetWithTTL(_ context.Context, key string, value interface{}) (bool, error) {
	dataJSON, err := json.Marshal(value)
	if err != nil {
		return false, pkgerr.NewInternalError(
			"ristretto.SetWithTTL",
			fmt.Sprintf("json marshal error for value %s", value),
			err,
		)
	}
	return rc.cache.SetWithTTL(key, string(dataJSON), 1, 5*time.Minute), nil
}

func (rc *ristrettoCache) Del(_ context.Context, key string) {
	rc.cache.Del(key)
}
