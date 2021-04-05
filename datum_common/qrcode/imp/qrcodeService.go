/**
 * @Author: sxiaohao
 * @Description:
 * @File:  qrcodeService
 * @Version: 1.0.0
 * @Date: 2020/11/28 下午10:17
 */

package imp

import (
	"context"
	proto "datum_common/proto/qrcode"
	"datum_common/upload"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"strconv"
	"time"
)

type QrcodeService struct {
}

func (qr *QrcodeService) GetQrcode(ctx context.Context, req *proto.GetQrcodeRequest, rep *proto.GetQrcodeResponse) error {

	content := req.GetContent()

	//生成18位随机数并生成本地路径文件名
	rand.Seed(time.Now().UnixNano())
	fileName := strconv.Itoa(rand.Int())

	_ = qrcode.WriteFile(content, qrcode.Medium, 256, "./qrcode/picture/"+fileName+".png")

	upload.StartUpload("", "./qrcode/picture/"+fileName+".png")

	//上传又拍云
	dst, result := upload.StartUpload("/datum/"+fileName+".png", "./qrcode/picture/"+fileName+".png")

	rep.Msg = result
	rep.Path = dst

	return nil
}
