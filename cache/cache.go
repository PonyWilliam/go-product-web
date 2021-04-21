package cache

import (
	"encoding/json"
	"github.com/PonyWilliam/go-ProductWeb/global"
	product "github.com/PonyWilliam/go-product/proto"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/v2/util/log"
)
var products *product.Response_ProductInfos
func GetCache(key string)(*product.Response_ProductInfos,error){
	val,err := global.RedisDB.Get(key).Result()
	if err == redis.Nil || err != nil{
		return nil,err
	}else {
		if err := json.Unmarshal([]byte(val),&products);err != nil{
			return nil,err
		}
		return products,nil
	}
}
func SetCache(key string,products *product.Response_ProductInfos)error{
	content,err := json.Marshal(products)
	if err != nil{
		log.Fatal(err)
	}
	err = global.RedisDB.Set(key,content,global.Duration).Err()
	if err != nil{
		log.Fatal(err)
	}
	return nil
}