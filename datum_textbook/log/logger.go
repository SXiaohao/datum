/**
 * @Author: sxiaohao
 * @Description:
 * @File:  logger
 * @Version: 1.0.0
 * @Date: 2020/11/1 下午8:37
 */

package log

import (
	"datum_textbook/config"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

var Logger *log.Logger

/**
 * @Author sxiaohao
 * @Description
 * @Date 2020/11/1 下午10:39
 * @Param
 * @return
 **/
func InitLog() {

	//日志文件
	fileName := path.Join(config.GetString("log.logFilePath"), config.GetString("log.logFileName"))

	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	Logger = log.New()
	//设置输出
	Logger.Out = src

	//设置日志级别
	Logger.SetLevel(log.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Logger.SetReportCaller(true)

	//json 格式
	//Logger.SetFormatter(&log.JSONFormatter{})
}
