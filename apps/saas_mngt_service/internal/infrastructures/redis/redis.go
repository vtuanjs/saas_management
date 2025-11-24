package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/vtuanjs/saas_management/apps/saas_mngt_service/internal/core"
	pkgerr "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_err"
	pkgmust "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_must"
	pkgretry "github.com/vtuanjs/saas_management/packages/go/pkg/pkg_retry"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
	locker *redislock.Client
}

func NewRedisCache(cfg *core.Config) core.Cache {
	redisCfg := cfg.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port),
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})

	err := pkgretry.WithRetry(context.Background(), func() error {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err := client.Ping(ctx).Result()
		if err != nil {
			return pkgerr.NewRetryableError("ping_redis", "failed to ping redis", err)
		}
		return nil
	}, 10, 1*time.Second)

	if err != nil {
		pkgmust.Check("redis connection", err)
	}

	fmt.Printf("Connecting to Redis at %s:%d\n", redisCfg.Host, redisCfg.Port)

	return &redisCache{
		client: client,
		locker: redislock.New(client),
	}
}

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return val, nil
		}
		return val, pkgerr.NewInternalError("redis.Get", fmt.Sprintf("key %s", key), err)
	}

	return val, nil
}

func (c *redisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return pkgerr.NewInternalError("redis.Set", fmt.Sprintf("json marshal error for key %s", key), err)
	}
	if err := c.client.Set(ctx, key, b, expiration).Err(); err != nil {
		return pkgerr.NewInternalError("redis.Set", fmt.Sprintf("failed to set key %s", key), err)
	}
	return nil
}

func (c *redisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *redisCache) Incr(ctx context.Context, key string) (int64, error) {
	val, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, pkgerr.NewInternalError("redis.Incr", fmt.Sprintf("key %s", key), err)
	}
	return val, nil
}

func (c *redisCache) Decr(ctx context.Context, key string) (int64, error) {
	val, err := c.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, pkgerr.NewInternalError("redis.Decr", fmt.Sprintf("key %s", key), err)
	}
	return val, nil
}

func (c *redisCache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := c.client.Exists(ctx, key).Result()
	if err != nil {
		return false, pkgerr.NewInternalError("redis.Exists", fmt.Sprintf("key %s", key), err)
	}
	return val == 1, nil
}

func (c *redisCache) WithDistributedLock(
	ctx context.Context,
	key string,
	ttlSeconds int,
	fn func(ctx context.Context) error,
) error {
	lockTTL := time.Duration(ttlSeconds) * time.Second
	lock, err := c.locker.Obtain(ctx, key, lockTTL, nil)
	if err == redislock.ErrNotObtained {
		return pkgerr.NewRetryableError("redis.WithDistributedLock", fmt.Sprintf("could not obtain lock for key %s", key), err)
	} else if err != nil {
		return pkgerr.NewInternalError("redis.WithDistributedLock", fmt.Sprintf("failed to obtain lock for key %s", key), err)
	}
	defer func() {
		if err := lock.Release(ctx); err != nil {
			fmt.Printf("Failed to release lock for key %s: %v\n", key, err)
			return
		}
	}()

	return fn(ctx)
}
