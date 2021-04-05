package redislocker

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func Lock(ctx context.Context, scripter redis.Scripter, key string, value interface{}, expire int64) error {
	script := redis.NewScript(`
		if redis.call("SETNX", KEYS[1], ARGV[1]) == 1 then
			return redis.call("expire", KEYS[1], ARGV[2])
		else
			return 0
		end
	`)

	result, err := script.Run(ctx, scripter, []string{key}, value, expire).Result()
	if err != nil {
		return err
	}

	if result.(int64) != 1 {
		return errors.New(fmt.Sprintf("lock failed, result: %d", result))
	}

	return nil
}

func Unlock(ctx context.Context, scripter redis.Scripter, key string) error {
	script := redis.NewScript(`return redis.call('del', KEYS[1]) ~= false`)

	_, err := script.Run(ctx, scripter, []string{key}).Result()
	if err != nil {
		return err
	}

	return nil
}

// key 和 value 同时相等的时候，更新过期时间
func Refresh(ctx context.Context, scripter redis.Scripter, key string, value interface{}, expire int64) error {
	script := redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("expire", KEYS[1], ARGV[2]) else return 0 end`)

	r, err := script.Run(ctx, scripter, []string{key}, value, expire).Result()
	if err != nil {
		return err
	}
	if r.(int64) != 1 {
		return errors.New(fmt.Sprintf("refresh failed, key: %s, value: %s, expire: %d", key, value, expire))
	}

	return nil
}
