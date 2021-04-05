/**
 * @Author: sxiaohao
 * @Description:
 * @File:  db
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package dao

import (
	"datum_textbook/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() {
	conn := config.GetString("database.local")
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db = db
}
