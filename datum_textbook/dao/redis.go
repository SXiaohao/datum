/**
 * @Author: sxiaohao
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2020/11/9 下午8:56
 */

package dao

import (
	"datum_textbook/config"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetString("redis.addr"),
		Password: config.GetString("redis.password"),
		DB:       0, // use default DB
	})
	_, err := RedisClient.Ping().Result()

	if err != nil {
		panic(err)
	}
	RedisClient.Del("KeywordList")
	RedisClient.ZAdd("KeywordList",
		redis.Z{Score: 0, Member: config.GetString("datum.keywordList.1")},
		redis.Z{Score: 1, Member: config.GetString("datum.keywordList.2")},
		redis.Z{Score: 2, Member: config.GetString("datum.keywordList.3")},
		redis.Z{Score: 3, Member: config.GetString("datum.keywordList.4")},
		redis.Z{Score: 4, Member: config.GetString("datum.keywordList.5")},
		redis.Z{Score: 5, Member: config.GetString("datum.keywordList.6")},
		redis.Z{Score: 6, Member: config.GetString("datum.keywordList.7")},
		redis.Z{Score: 7, Member: config.GetString("datum.keywordList.8")},
		redis.Z{Score: 8, Member: config.GetString("datum.keywordList.9")},
		redis.Z{Score: 9, Member: config.GetString("datum.keywordList.10")},
	)

}
