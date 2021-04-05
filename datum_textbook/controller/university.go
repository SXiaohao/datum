/**
 * @Author: sxiaohao
 * @Description:
 * @File:  university
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:48
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
 * @Description 根据专业类型获取所有学校名称
 * @Date 2020/11/2 下午8:07
 * @Param
 * @return
 **/
func ListByType(ctx *gin.Context) {
	universityType := ctx.Query("type")
	var universityList []models.University

	if strings.TrimSpace(universityType) == "" {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	//获取redis数据
	if exist := utils.GetRedisVal(ctx, "universityListByType"+universityType, &universityList); exist {
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"dataList": universityList}, "success!")
		return
	}

	//按类型查询学校
	result := dao.Db.Table("university").Where("type=?", universityType).Scan(&universityList)
	if result.Error != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, result.Error.Error())
		panic(result.Error.Error())
	} else if result.RowsAffected == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
		panic("未查询到数据！！")
	}

	//存入缓存  有效期1小时
	utils.SetRedisVal(ctx, "universityListByType"+universityType, universityList, 3600*time.Second)

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"dataList": universityList}, "success!")
	return
}
