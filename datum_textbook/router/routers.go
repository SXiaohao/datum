/**
 * @Author: sxiaohao
 * @Description:
 * @File:  routers
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package router

import (
	"datum_textbook/controller"
	"datum_textbook/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouters() (r *gin.Engine) {
	r = gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.DebugToFile())
	//r.Use(middleware.ParseToUid())
	datumGroup := r.Group("/datum")

	{
		datumGroup.GET("/detail", controller.Detail)
		datumGroup.GET("/list", controller.DatumList)
		datumGroup.GET("/recommend/list", controller.RecommendList)
		datumGroup.GET("/download/list", controller.DownloadList)
		datumGroup.GET("/keyword/list", controller.KeywordList)
		datumGroup.GET("/free/list", controller.FreeList)
		datumGroup.GET("/search", controller.Search)
		//middleware.ParseToUid()
		datumGroup.POST("/upload/picture", middleware.ParseToUid(), controller.UploadPicture)
		//middleware.ParseToUid()
		datumGroup.POST("/upload/file", middleware.ParseToUid(), controller.UploadFile)
		//middleware.ParseToUid()
		datumGroup.POST("/insert", middleware.ParseToUid(), controller.Insert)
		datumGroup.PUT("/keyword/list", controller.UpdateKeywordList)
	}

	majorGroup := r.Group("/major")
	{
		majorGroup.GET("/all", controller.ListById)
	}

	textbookGroup := r.Group("/textbook")
	{
		textbookGroup.GET("/detail", controller.TextbookDetail)
		textbookGroup.GET("/list", controller.TextbookList)
		textbookGroup.GET("/recommend/list", controller.TextbookRecommendList)
		textbookGroup.POST("/upload/picture", controller.UploadTextbookPicture)
	}

	universityGroup := r.Group("/university")
	{
		universityGroup.GET("/all", controller.ListByType)
	}

	commonGroup := r.Group("/common")
	{
		commonGroup.GET("/search", controller.Search)
		//middleware.ParseToUid()
		commonGroup.GET("/myPublish", controller.MyPublish)
		commonGroup.POST("/qrcode", controller.Qrcode)
	}

	return
}
