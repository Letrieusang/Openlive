package utils

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

var StatusLocked = "process_locked"

type OptimisticLock struct {
	Client   *redis.Client
	Duration time.Duration
	Prefix   string
}

func (o OptimisticLock) Lock(key string, expired ...time.Duration) error {
	duration := o.Duration
	if len(expired) > 0 {
		duration = expired[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	allowed, err := o.Client.SetNX(ctx, o.Prefix+key, "ok", duration).Result()
	if err != nil {
		return err
	}
	if !allowed {
		return errors.New(StatusLocked)
	}
	return nil
}

func (o OptimisticLock) Release(key string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := o.Client.Del(ctx, o.Prefix+key).Err()
	return err
}
