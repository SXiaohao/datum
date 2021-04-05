/**
 * @Author: sxiaohao
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package main

import (
	"datum_textbook/dao"
	"datum_textbook/log"
	"datum_textbook/router"
)

func main() {
	//启动日志
	log.InitLog()

	//连接数据库
	dao.InitDB()

	//连接redis
	dao.InitRedis()

	//配置路由> _ <
	r := router.SetupRouters()

	//link start
	r.Run(":8880")

}
