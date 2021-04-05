/**
 * @Author: sxiaohao
 * @Description:
 * @File:  token
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package middleware

import (
	"context"
	proto "datum_textbook/proto/token"
	"datum_textbook/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParseToUid() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			utils.Response(c, http.StatusCreated, nil, "access fail,don't find token")
			c.Abort()
			return
		}

		tokenReq := proto.ParseTokenRequest{Token: token}

		res, err := utils.GetTokenClient().ParseToken(context.Background(), &tokenReq)

		if err != nil {
			utils.Response(c, http.StatusCreated, nil, "token parse fail")
			c.Abort()
			return
		} else {
			c.Set("uid", res.Uid)
			c.Next()
		}

	}

}
