/**
 * @Author: sxiaohao
 * @Description:
 * @File:  upload
 * @Version: 1.0.0
 * @Date: 2020/11/16 下午9:15
 */

package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author sxiaohao
 * @Description //获取文件新文件名和文件类型并判断文件类型是否符合上传
 * @Date 2020/11/16 下午9:23
 * @Param fileStr string
 * @return filename string, filetype string
 **/
func GetFileNameType(fileStr string) (string, string, bool) {
	var name = []string{".zip", ".rar", ".docx", ".pdf", ".doc", ".png", ".jpeg", ".jpg"}
	var exist = false
	var fileName string
	var fileType string
	tracer := strings.LastIndex(fileStr, ".")

	fileType = fileStr[tracer:]

	for _, v := range name {
		if fileType == v {
			exist = true
		}
	}

	if exist {
		//生成18位随机数并生成本地路径文件名
		rand.Seed(time.Now().UnixNano())
		fileName = strconv.Itoa(rand.Int())
		return fileName, fileType, exist
	}

	return fileName, fileType, exist
}

/**
 * @Author sxiaohao
 * @Description //数组平分
 * @Date 2020/11/17 下午12:22
 * @Param arr []uint8, num int64
 * @return [][]uint8
 **/
func SplitArray(arr []uint8, num int64) [][]uint8 {

	max := int64(len(arr))
	if max < num {
		return nil
	}
	var divideArray = make([][]uint8, 0)
	quantity := max / num
	var end int64
	for i := int64(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			divideArray = append(divideArray, arr[i-1+end:qu])
		} else {
			divideArray = append(divideArray, arr[i-1+end:])
		}
		end = qu - i
	}
	return divideArray
}

/**
 * @Author sxiaohao
 * @Description 取整
 * @Date 2020/11/16 下午2:02
 * @Param divisor ,dividend int64
 * @return int64
 **/
func Round(divisor, dividend int64) int64 {
	if dividend == 0 {
		return divisor / dividend
	}
	if divisor%dividend != 0 {
		return divisor/dividend + 1
	} else {
		return divisor / dividend
	}

}
