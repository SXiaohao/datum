package upload

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"datum_common/config"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/upyun/go-sdk/upyun"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//上传又拍云
func StartUpload(path string, localPath string) (string, string) {
	//初始化 UpYun
	up := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   config.GetString("upYun.bucket"),
		Operator: config.GetString("upYun.operator"),
		Password: config.GetString("upYun.password"),
	})

	// 上传文件
	result := up.Put(&upyun.PutObjectConfig{
		Path:      path,
		LocalPath: localPath,
	})

	//上传又拍云后删除本地文件
	go os.Remove(localPath)

	if result != nil {
		return "", result.Error()
	} else {
		return config.GetString("upYun.addr") + path, "success"
	}

}

//图片识别(图片内容)
func CheckPicture(size int64, imgContent []byte, imgType string) (int64, string) {
	key := config.GetString("upYun.check.key")
	secret := config.GetString("upYun.check.secret")
	host := config.GetString("upYun.check.host")
	url := config.GetString("upYun.check.check_picture_url")
	method := config.GetString("upYun.check.method")
	date := makeRFC1123Date(time.Now())

	sign := sign(key, secret, method, url, date, "", "")

	req, _ := http.NewRequest(method, "http://"+host+url, bytes.NewBuffer(imgContent))
	req.Header.Add("Host", host)
	req.Header.Add("Date", date)
	req.Header.Add("Authorization", sign)
	req.Header.Add("Content-Length", strconv.FormatInt(size, 10))
	req.Header.Add("Content-Type", "image/"+imgType[1:])

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	var result map[string]map[string]interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	pornLabel := int64(result["porn"]["label"].(float64))
	politicalLabel := int64(result["political"]["label"].(float64))

	if politicalLabel != 100 || pornLabel != 0 {
		statusCode := politicalLabel*10 + pornLabel
		msg := "上传失败，图片涉及违规内容！"
		return statusCode, msg
	}
	return 0, ""
}

/**
 * @Author sxiaohao
 * @Description
 * @Date 2020/11/15 下午6:45
 * @Param text string
 * @return int64, string
 **/
func CheckText(text string) (int64, string) {
	type Data struct {
		Text string `json:"text"`
	}

	key := config.GetString("upYun.check.key")
	secret := config.GetString("upYun.check.secret")
	host := config.GetString("upYun.check.host")
	url := config.GetString("upYun.check.check_text_url")
	method := config.GetString("upYun.check.method")
	date := makeRFC1123Date(time.Now())
	sign := sign(key, secret, method, url, date, "", "")

	data := Data{Text: text}
	buff, _ := json.Marshal(data)

	req, _ := http.NewRequest(method, "http://"+host+url, bytes.NewBuffer(buff))
	req.Header.Add("Host", host)
	req.Header.Add("Date", date)
	req.Header.Add("Authorization", sign)
	req.Header.Add("Content-Length", strconv.Itoa(len(text)))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	var result map[string]map[string]interface{}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return 3, "检测接口出错！"
	}
	_ = json.Unmarshal(body, &result)

	spamLabel := int64(result["spam"]["label"].(float64))

	switch spamLabel {
	case 0:
		return 0, ""
	case 1:
		msg := "涉及违规内容！"
		return 1, msg
	case 2:
		return 2, "疑似有违规内容！"
	default:
		return 3, "检测接口出错！"
	}
}

func md5Str(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
func makeRFC1123Date(d time.Time) string {
	utc := d.UTC().Format(time.RFC1123)
	return strings.Replace(utc, "UTC", "GMT", 0)
}
func base64ToStr(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

/**
 * @Author sxiaohao
 * @Description 生成签名
 * @Date 2020/11/14 下午3:38
 * @Param key, secret, method, url, date, policy, md5
 * @return string
 **/
func sign(key, secret, method, url, date, policy, md5 string) string {
	const name = "UPYUN "
	mac := hmac.New(sha1.New, []byte(secret))
	var signArr []string
	for _, v := range []string{method, url, date, policy, md5} {
		if v != "" {
			signArr = append(signArr, v)
		}
	}
	value := strings.Join(signArr, "&")
	//fmt.Println(value)
	mac.Write([]byte(value))
	signStr := base64ToStr(mac.Sum(nil))
	return "UPYUN " + key + ":" + signStr
}
