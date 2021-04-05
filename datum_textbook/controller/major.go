/**
 * @Author: sxiaohao
 * @Description:
 * @File:  major
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:47
 */

package controller

import (
	"datum_textbook/dao"
	"datum_textbook/models"
	"datum_textbook/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

/**
 * @Author sxiaohao
 * @Description 根据学校id获取所有专业
 * @Date 2020/11/2 下午8:05
 * @Param
 * @return
 **/
func ListById(ctx *gin.Context) {
	id := ctx.Query("id")
	var majorsList []models.Major
	if strings.TrimSpace(id) == "" {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}
	//获取redis数据
	if exist := utils.GetRedisVal(ctx, "majorListById"+id, &majorsList); exist {
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"dataList": majorsList}, "success!")
		return
	}

	//按学校id获取所有专业名称
	result := dao.Db.Table("major").Where("university_id=?", id).Find(&majorsList)
	if result.Error != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, result.Error.Error())
		panic(result.Error.Error())

	} else if result.RowsAffected == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
		panic(result.Error)
	}

	//存入redis 有效时间1小时
	utils.SetRedisVal(ctx, "majorListById"+id, majorsList, 3600*time.Second)

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"dataList": majorsList}, "success!")
	return
}

func ListByT() {

}
