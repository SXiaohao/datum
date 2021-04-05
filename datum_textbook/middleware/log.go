/**
 * @Author: sxiaohao
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2020/11/2 下午12:17
 */

package middleware

import (
	"datum_textbook/log"
	"github.com/gin-gonic/gin"
	"net/url"
)

/**
 * @Author sxiaohao
 * @Description 将错误信息格式化写入日志文件
 * @Date 2020/11/2 下午12:33
 * @Param
 * @return
 **/
func DebugToFile() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {

				// 请求方式
				reqMethod := c.Request.Method

				// 请求路由
				reqUri := c.Request.RequestURI

				// 请求IP
				clientIP := c.ClientIP()

				// post 参数
				postData, _ := url.QueryUnescape(c.Request.PostForm.Encode())

				//用户设备
				userAgent := c.Request.UserAgent()

				log.Logger.Debugf("| %s | %15s | %s | %s | postData:%s | errMsg:%s |",
					reqMethod,
					reqUri,
					clientIP,
					userAgent,
					postData,
					err)
			}
		}()
		c.Next()
	}
}
