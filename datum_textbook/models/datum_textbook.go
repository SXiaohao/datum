/**
 * @Author: sxiaohao
 * @Description:
 * @File:  datum_textbook
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

package models

//资料表
type Datum struct {
	Id            int     `json:"id" gorm:"primaryKey"`               //'资料id',
	Uid           int     `json:"uid"`                                //'用户id',
	MajorId       int     `json:"major_id" gorm:"ForeignKey:MajorId"` //'专业id',
	UniversityId  int     `json:"university_id"`                      //`学校id`
	Type          int     `json:"type"`                               //'1是升本2是考研',
	Title         string  `json:"title"`                              //'标题',
	Desc          string  `json:"desc"`                               //'介绍',
	Picture       string  `json:"picture"`                            //'图片',
	Price         float64 `json:"price"`                              //'价钱',
	FileLink      string  `json:"file_link"`                          //'文件链接',
	PurchaseCount int     `json:"purchase_count"`                     //'购买次数',
	CreateTime    int64   `gorm:"autoCreateTime" json:"create_time"`  //'创建时间',
	UpdateTime    int64   `json:"update_time"`                        //'更新时间',
	DeleteTime    int64   `json:"delete_time"`                        //'删除时间默认为null (软删除)',
}

//教材表
type Textbook struct {
	Id            int     `json:"id" gorm:"primaryKey"` //'教材id',
	Uid           int     `json:"uid"`                  //'用户id',
	MajorId       int     `json:"major_id"`             //'专业id',
	Author        string  `json:"author"`               //'作者'
	Name          string  `json:"name"`                 //'教材名',
	Isbn          string  `json:"isbn"`                 //'ISBN码',
	Picture       string  `json:"picture"`              //'图片链接',
	PurchaseCount int     `json:"purchase_count"`       //'购买次数',
	Price         float64 `json:"price"`                //'价格',
	CreateTime    int64   `json:"create_time"`          //'创建时间',
	UpdateTime    int64   `json:"update_time"`          //'更新时间',
	DeleteTime    int64   `json:"delete_time"`          //'删除时间默认为null (软删除)',
}

//学校表
type University struct {
	Id        int    `json:"id" gorm:"primaryKey"` //'学校id',
	Name      string `json:"name"`                 //'学校名',
	Type      int    `json:"type"`                 //'1升本 2考研',
	ShortName string `json:"short_name"`
}

//专业表
type Major struct {
	Id           int    `json:"id" gorm:"primaryKey"`          //'专业id',
	UniversityId int    `json:"university_id" ForeignKey:"Id"` //'学校id',
	Title        string `json:"title"`                         //'专业名称',
}
