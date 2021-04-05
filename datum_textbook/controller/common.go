/**
 * @Author: sxiaohao
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2020/11/28 下午4:08
 */

package controller

import (
	"context"
	"datum_textbook/dao"
	"datum_textbook/models"
	proto "datum_textbook/proto/qrcode"
	"datum_textbook/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

/**
 * @Author sxiaohao
 * @Description 搜索资料
 * @Date 2020/11/24 下午7:47
 * @Param
 * @return
 **/
func Search(ctx *gin.Context) {
	var datumList []models.Datum
	var textbookList []models.Textbook
	keyword := ctx.Query("keyword")
	if strings.TrimSpace(keyword) == "" {
		utils.Response(ctx, http.StatusCreated, nil, "参数不能为空！！")
		panic("参数不能为空！！")
	}

	datumResult := dao.Db.Table("datum").Where(" title like ? or `desc`  like ? ", "%"+keyword+"%", "%"+keyword+"%").Find(&datumList)
	textbookResult := dao.Db.Table("textbook").Where(" name like ?  ", "%"+keyword+"%").Find(&textbookList)
	if datumResult.Error != nil {
		panic(datumResult.Error.Error())
	} else if textbookResult.Error != nil {
		panic(textbookResult.Error.Error())
	}

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"datumList": datumList, "bookList": textbookList}, "success!!")
	return

}

/**
 * @Author sxiaohao
 * @Description 获取用户发布列表
 * @Date 2020/11/28 下午9:13
 * @Param
 * @return
 **/
func MyPublish(ctx *gin.Context) {
	//每页数据数量
	const DataCount = 10
	var page = 0
	var publishType = 0
	var datumList []models.Datum
	var textbookList []models.Textbook
	var total int64 = 0

	page, _ = strconv.Atoi(ctx.Query("page"))
	publishType, _ = strconv.Atoi(ctx.Query("type"))
	//uid, _ := ctx.Get("uid")
	uid := 6312

	if page == 0 || publishType == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}
	switch publishType {
	case 1:
		//查询资料
		result := dao.Db.Table("datum").Where(" delete_time = 0 AND uid = ?", uid).Count(&total).Order("create_time desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&datumList)
		if result.Error != nil {
			utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
			panic(result.Error)
		} else if result.RowsAffected == 0 {
			utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
			panic(result.Error)
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": datumList}, "success!")
		break

	case 2:
		//查询教材
		result := dao.Db.Table("textbook").Where(" delete_time = 0 AND uid = ?", uid).Count(&total).Order("create_time desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&textbookList)
		if result.Error != nil {
			utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
			panic(result.Error)
		} else if result.RowsAffected == 0 {
			utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
			panic(result.Error)
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": textbookList}, "success!")
		break
	default:
		utils.Response(ctx, http.StatusCreated, nil, "没有此查询方式！！")
		break
	}

	return
}

/**
 * @Author sxiaohao
 * @Description 获取二维码
 * @Date 2020/11/29 下午1:57
 * @Param
 * @return
 **/
func Qrcode(ctx *gin.Context) {

	content := ctx.Query("text")

	request := proto.GetQrcodeRequest{}
	request.Content = content

	response, err := utils.GetQrcodeClient().GetQrcode(context.Background(), &request)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, err.Error())
		panic(err)
	} else if response.Msg != "success" {
		utils.Response(ctx, http.StatusCreated, nil, response.Msg)
		panic(response.Msg)
	}
	utils.Response(ctx, http.StatusOK, map[string]interface{}{"qrcode": response.Path}, "success!!")
	return
}
