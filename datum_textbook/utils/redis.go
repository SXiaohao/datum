/**
 * @Author: sxiaohao
 * @Description:
 * @File:  redis
 * @Version: 1.0.0
 * @Date: 2020/11/11 下午3:09
 */

package utils

import (
	"datum_textbook/dao"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Author sxiaohao
 * @Description  获取缓存数据
 * @Date 2020/11/11 下午4:10
 * @Param ctx,key,val
 * @return  bool
 **/
func GetRedisVal(ctx *gin.Context, key string, val interface{}) bool {
	//获取redis数据
	data, _ := dao.RedisClient.Get(key).Result()

	//redis数据是否存在
	if data != "" {
		//json转对象
		if err := json.Unmarshal([]byte(data), val); err != nil {
			Response(ctx, http.StatusInternalServerError, nil, "json转换数据发生错误！")
			panic(err)
		}
		return true
	}
	return false

}

/**
 * @Author sxiaohao
 * @Description	 创建缓存数据
 * @Date 2020/11/11 下午6:52
 * @Param ctx,key,data,expiration
 * @return
 **/
func SetRedisVal(ctx *gin.Context, key string, data interface{}, expiration time.Duration) {
	//对象转json
	if val, err := json.Marshal(data); err != nil {
		Response(ctx, http.StatusInternalServerError, nil, "数据转换json发生错误")
		panic(err)

	} else {
		//存入redis
		_, err := dao.RedisClient.Set(key, val, expiration).Result()
		if err != nil {
			Response(ctx, http.StatusInternalServerError, nil, "缓存发生错误！！")
			panic(err)
		}
	}
}
