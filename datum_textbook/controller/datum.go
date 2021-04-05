/**
 * @Author: sxiaohao
 * @Description:
 * @File:  datum
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package controller

import (
	"context"
	"datum_textbook/dao"
	"datum_textbook/models"
	proto "datum_textbook/proto/upload"
	"datum_textbook/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DatumDetail struct {
	Datum          models.Datum `json:"datum"  `
	UniversityName string       `json:"university_name"`
	MajorName      string       `json:"major_name"`
	IsPaid         bool         `json:"is_paid"`
}

/**
 * @Author sxiaohao
 * @Description 资料详情
 * @Date 2020/10/30 下午3:18
 * @Param
 * @return Datum
 **/
func Detail(ctx *gin.Context) {

	id := ctx.Query("id")
	var datumDetail DatumDetail
	datumDetail.IsPaid = false

	if strings.TrimSpace(id) == "" {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	//获取缓存数据
	if exist := utils.GetRedisVal(ctx, "datumDetail"+id, &datumDetail); exist {
		utils.Response(ctx, http.StatusOK, datumDetail, "success！")
		return
	}

	//id查询资料并检查是否被删除
	result := dao.Db.Table("datum").Select("datum.* ,major.title AS major_name , university.name AS university_name").Joins("INNER JOIN major ON datum.major_id=major.id").Joins("INNER JOIN university ON datum.university_id=university.id").Where("datum.id=? AND datum.delete_time = 0", id).Row()

	if result.Err() != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, "服务器崩溃！！")
		panic(result.Err().Error())
	}

	//填充结构体
	err := result.Scan(&datumDetail.Datum.Id,
		&datumDetail.Datum.Uid,
		&datumDetail.Datum.MajorId,
		&datumDetail.Datum.UniversityId,
		&datumDetail.Datum.Type,
		&datumDetail.Datum.Title,
		&datumDetail.Datum.Desc,
		&datumDetail.Datum.Picture,
		&datumDetail.Datum.Price,
		&datumDetail.Datum.FileLink,
		&datumDetail.Datum.PurchaseCount,
		&datumDetail.Datum.CreateTime,
		&datumDetail.Datum.UpdateTime,
		&datumDetail.Datum.DeleteTime,
		&datumDetail.MajorName,
		&datumDetail.UniversityName)

	if err != nil {
		utils.Response(ctx, http.StatusCreated, nil, "未查询到此资料！！")
		panic(err)
	}

	//存入缓存
	utils.SetRedisVal(ctx, "datumDetail"+id, datumDetail, 3600*time.Second)

	utils.Response(ctx, http.StatusOK, datumDetail, "success！")
	return

}

/**
 * @Author sxiaohao
 * @Description 获取资料列表
 * @Date 2020/11/2 下午8:23
 * @Param
 * @return
 **/
func DatumList(ctx *gin.Context) {
	//每页数据数量
	const DataCount = 8
	var page = 1
	var datumType = 1
	var sort = 1
	var universityId = 0
	var majorId = 0
	var dataList []models.Datum
	var total int64 = 0

	page, _ = strconv.Atoi(ctx.Query("page"))
	datumType, _ = strconv.Atoi(ctx.Query("type"))
	sort, _ = strconv.Atoi(ctx.Query("sort"))
	universityId, _ = strconv.Atoi(ctx.Query("university_id"))
	majorId, _ = strconv.Atoi(ctx.Query("major_id"))

	switch sort {
	case 1: //购买量降序
		if universityId == 0 {
			//用majorId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 AND major_id = ?", datumType, majorId).Count(&total).Order("purchase_count desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)

			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}
			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break

		} else {
			//用universityId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 AND university_id = ?", datumType, universityId).Count(&total).Order("purchase_count desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}

			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break
		}
	case 2: //发布时间降序
		if universityId == 0 {
			//用majorId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 AND major_id = ?", datumType, majorId).Count(&total).Order("create_time desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}
			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break
		} else {
			//用universityId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 and university_id=?", datumType, universityId).Limit(DataCount).Count(&total).Order("create_time desc").Offset((page - 1) * DataCount).Find(&dataList)
			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}
			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break
		}
	case 3: //价格降序
		if universityId == 0 {
			//用majorId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 AND major_id = ?", datumType, majorId).Count(&total).Order("price desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}
			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break
		} else {
			//用universityId查询资料
			result := dao.Db.Table("datum").Where("type=? AND delete_time = 0 and university_id=?", datumType, universityId).Limit(DataCount).Count(&total).Order("price desc").Offset((page - 1) * DataCount).Find(&dataList)
			if result.Error != nil {
				utils.Response(ctx, http.StatusInternalServerError, nil, "数据库发生错误！！")
				panic(result.Error)
			} else if result.RowsAffected == 0 {
				utils.Response(ctx, http.StatusCreated, nil, "未查询到数据！！")
				panic(result.Error)
			}
			utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!")
			break
		}
	default:
		utils.Response(ctx, http.StatusCreated, nil, "没有此排序方式！！")
		break
	}

	return

}

/**
 * @Author sxiaohao
 * @Description 获取相关推荐资料列表
 * @Date 2020/11/11 下午7:32
 * @Param
 * @return
 **/
func RecommendList(ctx *gin.Context) {
	//每页数据数量
	const DataCount = 5
	var page = 0      //当前页数
	var excludeId = 0 //资料id，查询结果会排除此资料
	var majorId = 0   //专业id
	var dataList []models.Datum = nil
	var total int64 = 0

	page, _ = strconv.Atoi(ctx.Query("page"))
	excludeId, _ = strconv.Atoi(ctx.Query("exclude_id"))
	majorId, _ = strconv.Atoi(ctx.Query("major_id"))

	if page == 0 || excludeId == 0 || majorId == 0 {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	//用majorId查询资料 不包括excludeId  购买量降序然后创建时间降序
	result := dao.Db.Table("datum").Where("NOT id = ?  AND delete_time = 0 AND major_id = ?", excludeId, majorId).Count(&total).Order("purchase_count desc").Order("create_time desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
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
 * @Description 上传资料图片
 * @Date 2020/11/2 下午12:50
 * @Param
 * @return filePath
 **/
func UploadPicture(ctx *gin.Context) {
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
	request.UploadPath = "/datum/datum/picture"

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

/**
 * @Author sxiaohao
 * @Description
 * @Date 2020/11/16 下午12:31
 * @Param
 * @return
 **/
func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, err.Error())
		panic(err)
	} else if file.Size >= 41943040 {
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

	request := proto.UploadFileRequest{}

	//把二进制文件平分成若干份
	copies := utils.Round(file.Size, 4000000)

	//响应期限
	c, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Duration(5*copies)*time.Second))
	//平分二进制
	arr := utils.SplitArray(buff, copies)
	for i := 0; i < len(arr); i++ {
		request.Content = arr[i]
		request.Size = file.Size
		_, _ = utils.GetUploadClient().GetFilePath(c, &request)
	}
	request.Content = nil
	request.Size = 0
	request.FileName = newFileName
	request.FileType = fileType

	response, err := utils.GetUploadClient().GetFilePath(c, &request)
	cancel()

	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, map[string]interface{}{"path": ""}, err.Error())
		panic(err)
	} else if response.StatusCode != 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"path": ""}, response.Msg)
		panic(response.StatusCode)
	}

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"path": response.Path}, response.Msg)
	return

}

/**
 * @Author sxiaohao
 * @Description 添加资料
 * @Date 2020/11/2 下午4:27
 * @Param
 * @return
 **/
func Insert(ctx *gin.Context) {

	var majorId = 0
	var datumType = 0
	var universityId = 0
	var price = 0.00

	uid, _ := ctx.Get("uid")

	majorId, _ = strconv.Atoi(ctx.PostForm("major_id"))
	universityId, _ = strconv.Atoi(ctx.PostForm("university_id"))
	datumType, _ = strconv.Atoi(ctx.PostForm("type"))
	title := ctx.PostForm("title")
	desc := ctx.PostForm("desc")
	picture := ctx.PostForm("picture")
	price, _ = strconv.ParseFloat(ctx.PostForm("price"), 32)
	fileLink := ctx.PostForm("file_link")

	if majorId == 0 || datumType == 0 || universityId == 0 || price == 0.00 {
		utils.Response(ctx, http.StatusCreated, nil, "参数错误！！")
		panic("参数错误！！")
	}

	request := proto.CheckTextRequest{}
	request.Text = desc
	response, err := utils.GetUploadClient().CheckText(context.Background(), &request)
	if err != nil {
		utils.Response(ctx, http.StatusInternalServerError, map[string]interface{}{}, err.Error())
		panic(err)
	} else if response.StatusCode != 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{}, response.Msg)
		panic(response.Msg)
	}

	//初始化datum模型
	datum := models.Datum{
		Uid:          uid.(int),
		MajorId:      majorId,
		UniversityId: universityId,
		Type:         datumType,
		Title:        title,
		Desc:         desc,
		Picture:      picture,
		Price:        price,
		FileLink:     fileLink,
	}

	//插入资料
	result := dao.Db.Table("datum").Create(&datum)

	if result.Error != nil {
		utils.Response(ctx, http.StatusInternalServerError, nil, result.Error.Error())
		panic(result.Error)
	}
	utils.Response(ctx, http.StatusOK, map[string]interface{}{}, "success!!")
	return
}

/**
 * @Author sxiaohao
 * @Description //获取资料下载榜列表
 * @Date 2020/11/17 下午12:54
 * @Param
 * @return
 **/
func DownloadList(ctx *gin.Context) {
	const DataCount = 10
	var downloadList []models.Datum
	var total int64 = 0
	var page = 0
	page, _ = strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"total": total, "dataList": nil}, "参数错误！！")
		panic("参数错误！！")
	}

	result := dao.Db.Table("datum").Where("delete_time = 0").Count(&total).Order("purchase_count desc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&downloadList)
	if result.Error != nil {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"total": total, "dataList": nil}, result.Error.Error())
		panic(result.Error.Error())
	}
	utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": downloadList}, "success!!")
	return
}

/**
 * @Author sxiaohao
 * @Description 更新关键字列表
 * @Date 2020/11/24 下午7:42
 * @Param
 * @return
 **/
func UpdateKeywordList(ctx *gin.Context) {
	var keywordList [10]string
	keywordList[0] = ctx.Query("1")
	keywordList[1] = ctx.Query("2")
	keywordList[2] = ctx.Query("3")
	keywordList[3] = ctx.Query("4")
	keywordList[4] = ctx.Query("5")
	keywordList[5] = ctx.Query("6")
	keywordList[6] = ctx.Query("7")
	keywordList[7] = ctx.Query("8")
	keywordList[8] = ctx.Query("9")
	keywordList[9] = ctx.Query("10")

	dao.RedisClient.Del("KeywordList")
	for i := 0; i < 10; i++ {
		if keywordList[i] != "" {
			dao.RedisClient.ZAdd("KeywordList", redis.Z{Score: float64(i), Member: keywordList[i]})
		}

	}
	utils.Response(ctx, http.StatusOK, nil, "success!")
}

/**
 * @Author sxiaohao
 * @Description 获取关键字列表
 * @Date 2020/11/24 下午7:42
 * @Param
 * @return
 **/
func KeywordList(ctx *gin.Context) {
	keywordCount, _ := dao.RedisClient.ZCard("KeywordList").Result()

	keywordList, err := dao.RedisClient.ZRange("KeywordList", 0, keywordCount).Result()

	if err != nil {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"result": nil}, "缓存错误！！")
		panic(err)
	} else if len(keywordList) == 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"result": nil}, "数据为空！！")
		panic("数据为空！！")
	}

	utils.Response(ctx, http.StatusOK, map[string]interface{}{"result": keywordList}, "success!")
}

/**
 * @Author sxiaohao
 * @Description 获取免费列表
 * @Date 2020/11/28 下午9:12
 * @Param
 * @return
 **/
func FreeList(ctx *gin.Context) {
	const DataCount = 8
	var dataList []models.Datum
	var total int64 = 0
	var page = 0
	page, _ = strconv.Atoi(ctx.Query("page"))
	if page == 0 {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"total": total, "dataList": nil}, "参数错误！！")
		panic("参数错误！！")
	}

	result := dao.Db.Table("datum").Where("delete_time = 0 AND price < 1.00").Count(&total).Order("purchase_count desc").Order("price asc").Limit(DataCount).Offset((page - 1) * DataCount).Find(&dataList)
	if result.Error != nil {
		utils.Response(ctx, http.StatusCreated, map[string]interface{}{"total": total, "dataList": nil}, result.Error.Error())
		panic(result.Error.Error())
	}
	utils.Response(ctx, http.StatusOK, map[string]interface{}{"total": total, "dataList": dataList}, "success!!")
	return
}
