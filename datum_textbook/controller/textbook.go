/**
 * @Author: sxiaohao
 * @Description:
 * @File:  textbook
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:48
 */

package controller

import (
	"context"
	"datum_textbook/dao"
	"datum_textbook/models"
	proto "datum_textbook/proto/upload"
	"datum_textbook/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author sxiaohao
 * @Description 教材详情
 * @Date 2020/11/25 下午9:24
 * @Param
 * @return
 **/
func TextbookDetail(ctx *gin.Context) {

	id := ctx.Query("id")
	var textbookDetail models.Textbook

	if strings.TrimSpace(id) == "" {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	//获取缓存数据
	if exist := utils.GetRedisVal(ctx, "textbookDetail"+id, &textbookDetail); exist {
		utils.Response(ctx, http.StatusOK, textbookDetail, "success！")
		return
	}

	//id查询资料并检查是否被删除
	result := dao.Db.Table("textbook").Where("id=? AND delete_time = 0", id).Find(&textbookDetail)

	if result.Error != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, "服务器崩溃！！")
		panic(result.Error.Error())
	} else if result.RowsAffected == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "未查询到此数据!")
		panic(result.Error.Error())
	}

	//存入缓存
	utils.SetRedisVal(ctx, "textbookDetail"+id, textbookDetail, 3600*time.Second)

	utils.Response(ctx, http.StatusOK, textbookDetail, "success！")
	return
}

/**
 * @Author sxiaohao
 * @Description 获取教材列表
 * @Date 2020/11/25 下午9:18
 * @Param
 * @return
 **/
func TextbookList(ctx *gin.Context) {
	//每页数据数量
	const DataCount = 8
	var page = 1
	var majorId int
	var dataList []models.Textbook
	var total int64 = 0

	page, _ = strconv.Atoi(ctx.Query("page"))
	majorId, _ = strconv.Atoi(ctx.Query("major_id"))

	if majorId == 0 {
		//不限制专业资料
		result := dao.Db.Table("textbook").Where(" delete_time = 0 ").Count(&total).Order("purchase_count desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
		if result.Error != nil {
			utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
			panic(result.Error)
		} else if result.RowsAffected == 0 {
			utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
			panic(result.Error)
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
		return
	} else {
		//用majorId查询资料
		result := dao.Db.Table("textbook").Where("major_id = ? AND delete_time = 0 ", majorId).Count(&total).Order("purchase_count desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)

		if result.Error != nil {
			utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
			panic(result.Error)
		} else if result.RowsAffected == 0 {
			utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
			panic(result.Error)
		}
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
		return
	}

}

/**
 * @Author sxiaohao
 * @Description 获取相关推荐教材列表
 * @Date 2020/11/28 下午3:57
 * @Param
 * @return
 **/
func TextbookRecommendList(ctx *gin.Context) {
	//每页数据数量
	const DataCount = 5
	var page = 0      //当前页数
	var excludeId = 0 //教材id，查询结果会排除此资料
	var majorId = 0   //专业id
	var dataList []models.Textbook = nil
	var total int64 = 0

	page, _ = strconv.Atoi(ctx.Query("page"))
	excludeId, _ = strconv.Atoi(ctx.Query("exclude_id"))
	majorId, _ = strconv.Atoi(ctx.Query("major_id"))

	if page == 0 || excludeId == 0 || majorId == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	//用majorId查询教材 不包括excludeId  购买量降序然后创建时间降序
	result := dao.Db.Table("textbook").Where("NOT id = ?  AND delete_time = 0 AND major_id = ?", excludeId, majorId).Count(&total).Order("purchase_count desc").Order("create_time desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
	if result.Error != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
		panic(result.Error)
	} else if result.RowsAffected == 0 {
		utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "暂无相关推荐！！")
		panic(result.Error)
	}

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")

}

/**
 * @Author sxiaohao
 * @Description 上传教材图片
 * @Date 2020/11/28 下午4:06
 * @Param
 * @return
 **/
func UploadTextbookPicture(ctx *gin.Context) {
	file, err := ctx.FormFile("picture")

	if err != nil {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, err.Error())
		panic(err)
	} else if file.Size >= 4194304 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, "图片内容过大！")
		panic(err)
	}

	newFileName, fileType, exist := utils.GetFileNameType(file.Filename)
	if !exist {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, "上传文件类型不正确！")
		panic("上传文件类型不正确！" + fileType)
	}

	newFile, _ := file.Open()
	//转换二进制存储到buff
	buff := make([]byte, file.Size)
	_, err = newFile.Read(buff)
	if err != nil {
		panic(err)
	}

	request := proto.UploadPictureRequest{}
	//文件名
	request.FileName = newFileName
	//大小
	request.Size = file.Size
	//文件二进制
	request.Content = buff
	//文件类型
	request.FileType = fileType
	//上传路径
	request.UploadPath = "/datum/book/picture"

	c, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Second))

	response, err := utils.GetUploadClient().GetPicturePath(c, &request)

	cancel()
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, map[string]interface{}{"path": ""}, err.Error())
		panic(err)
	}
	if response.StatusCode != 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, response.Msg)
		panic(response.StatusCode)
	}
	utils.Response(ctx, http.StatusOK, map[string]interface{}{"path": response.Path}, response.Msg)
	return
}
