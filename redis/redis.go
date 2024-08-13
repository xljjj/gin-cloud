package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"time"
)

var CLI *redis.Client

func InitRedis() {
	CLI = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	_, err := CLI.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
}

// GetKey 获取键对应的值
func GetKey(ctx context.Context, key string) (string, error) {
	return CLI.Get(ctx, key).Result()
}

// SetKey 设置键值对 举例：永不过期0，10秒过期10*time.Second
func SetKey(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return CLI.Set(ctx, key, value, expiration).Err()
}

// DeleteKey 删除键
func DeleteKey(ctx context.Context, key string) error {
	return CLI.Del(ctx, key).Err()
}

// LRange 获取列表中指定范围内的元素
func LRange(ctx context.Context, key string, start, stop int64) ([]string, error) {
	vals, err := CLI.LRange(ctx, key, start, stop).Result()
	if err != nil {
		return nil, err
	}
	return vals, nil
}

// LPop 从列表中弹出并返回第一个元素
func LPop(ctx context.Context, key string) (string, error) {
	val, err := CLI.LPop(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

// KeyExists 检查指定的key是否存在
func KeyExists(ctx context.Context, key string) (bool, error) {
	exists, err := CLI.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}

// KeyTTL 返回指定key的剩余过期时间
func KeyTTL(ctx context.Context, key string) (time.Duration, error) {
	ttl, err := CLI.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return ttl, nil
}

// Incr 自增指定键的值
func Incr(ctx context.Context, key string) (int64, error) {
	val, err := CLI.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// Decr 自减指定键的值
func Decr(ctx context.Context, key string) (int64, error) {
	val, err := CLI.Decr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// MSet 批量设置多个键值对
func MSet(ctx context.Context, values map[string]interface{}) error {
	_, err := CLI.MSet(ctx, values).Result()
	if err != nil {
		return err
	}
	return nil
}

// MGet 批量读取多个键的值
func MGet(ctx context.Context, keys ...string) ([]string, error) {
	vals, err := CLI.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, err
	}

	// 将结果转换为字符串切片
	results := make([]string, len(vals))
	for i, v := range vals {
		if v == nil {
			results[i] = "" // 键不存在时的默认值
		} else {
			results[i] = v.(string) // 结果是 string 类型
		}
	}
	return results, nil
}

// LPushAndTrimKey 将值推送到列表的左侧，并裁剪列表到指定的最大长度
func LPushAndTrimKey(ctx context.Context, key string, value string, maxLen int64) error {
	// 将值推送到列表的左侧
	_, err := CLI.LPush(ctx, key, value).Result()
	if err != nil {
		return err
	}

	// 裁剪列表到指定长度
	_, err = CLI.LTrim(ctx, key, 0, maxLen-1).Result()
	if err != nil {
		return err
	}

	return nil
}

// RPushAndTrimKey 将值推送到列表的右侧，并裁剪列表到指定的最大长度
func RPushAndTrimKey(ctx context.Context, key string, value string, maxLen int64) error {
	// 将值推送到列表的左侧
	_, err := CLI.RPush(ctx, key, value).Result()
	if err != nil {
		return err
	}

	// 裁剪列表到指定长度
	_, err = CLI.LTrim(ctx, key, 0, maxLen-1).Result()
	if err != nil {
		return err
	}

	return nil
}
