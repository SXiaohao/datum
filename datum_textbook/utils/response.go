/**
 * @Author: sxiaohao
 * @Description:
 * @File:  response
 * @Version: 1.0.0
 * @Date: 2020/10/30 下午2:35
 */

package utils

import "github.com/gin-gonic/gin"

/**
 * @Author sxiaohao
 * @Description //
 * @Date 2020/10/30 下午3:05
 * @Param	ctx,httpStatus,data,msg
 * @return	nil
 **/
func Response(ctx *gin.Context, httpStatus int, data interface{}, msg string) {

	ctx.JSON(httpStatus, gin.H{"code": httpStatus, "data": data, "msg": msg})

}
