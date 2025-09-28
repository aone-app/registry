package internal

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Registry struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRegistry() *Registry {
	redisHost := os.Getenv("REDIS_HOST") // Redis 服务地址
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: os.Getenv("REDIS_PWD"),
		DB:       0,
	})
	return &Registry{
		rdb: rdb,
		ctx: context.Background(),
	}
}

// Register 注册一个服务节点
func (r *Registry) Register(service, addr string, ttlSeconds int) error {
	key := fmt.Sprintf("services:%s", service)
	// 使用 set 存储节点列表
	if err := r.rdb.SAdd(r.ctx, key, addr).Err(); err != nil {
		return err
	}
	// 设置过期时间（如果需要心跳机制，可以定期续期）
	if ttlSeconds > 0 {
		r.rdb.Expire(r.ctx, key, time.Duration(ttlSeconds)*time.Second)
	}
	return nil
}

// Unregister 注销服务节点
func (r *Registry) Unregister(service, addr string) error {
	key := fmt.Sprintf("services:%s", service)
	return r.rdb.SRem(r.ctx, key, addr).Err()
}

// GetNodes 获取服务下的所有节点
func (r *Registry) GetNodes(service string) ([]string, error) {
	key := fmt.Sprintf("services:%s", service)
	return r.rdb.SMembers(r.ctx, key).Result()
}
