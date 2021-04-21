package global

import (
	"github.com/go-redis/redis"
	"time"
)
var(RedisDB *redis.Client)
const Duration = time.Minute * 5
func SetupRedisDb() error{
	RedisDB = redis.NewClient(&redis.Options{
		Addr: "106.13.132.160:6379",
		Password: nil,
		DB: 0,
	})
	_,err := RedisDB.Ping().Result()
	if err != nil{
		return err
	}
	return nil
}