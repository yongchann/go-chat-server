package redis

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	redisClient *Client
	once        sync.Once

	ErrRecordNotFound = errors.New("record not found")
)

type Client struct {
	ctx context.Context
	rc  *redis.Client
}

func NewClient(addr, password string) (*Client, error) {
	once.Do(func() {
		redisClient = &Client{
			ctx: context.Background(),
			rc: redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: password,
				DB:       0,
				PoolSize: 30,
			}),
		}

		if err := redisClient.rc.Ping(redisClient.ctx).Err(); err != nil {
			panic("Unable to connect to redis " + err.Error())
		}
	})

	return redisClient, nil
}

func (c *Client) GetKey(key string, dest interface{}) error {
	val, err := c.rc.Get(c.ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrRecordNotFound
		} else {
			return err
		}
	}

	if err = json.Unmarshal([]byte(val), &dest); err != nil {
		return err
	}

	return nil
}

func (c *Client) SetKey(key string, value interface{}, exp time.Duration) error {
	cacheEntry, err := json.Marshal(value)
	if err != nil {
		return err
	}

	if err := c.rc.Set(c.ctx, key, cacheEntry, exp).Err(); err != nil {
		return err
	}

	return nil
}
