package imp

import (
	"context"
	proto "datum_common/proto/upload"
	"datum_common/upload"
	"io/ioutil"
)

type UploadService struct {
}

func (*UploadService) GetPicturePath(ctx context.Context, request *proto.UploadPictureRequest, response *proto.UploadPictureResponse) error {
	fileName := request.GetFileName()
	fileType := request.GetFileType()
	size := request.GetSize()
	content := request.GetContent()
	uploadPath := request.GetUploadPath()

	//检测图片
	if statusCode, msg := upload.CheckPicture(size, content, fileType); statusCode != 0 {
		response.Path = ""
		response.Msg = msg
		response.StatusCode = statusCode
		return nil
	}

	//将文件保存至本项目指定目录中
	_ = ioutil.WriteFile("./upload/picture/"+fileName+fileType, content, 0666)

	//上传又拍云
	dst, result := upload.StartUpload(uploadPath+"/"+fileName+fileType, "./upload/picture/"+fileName+fileType)

	response.Path = dst
	response.Msg = result
	response.StatusCode = 0
	return nil
}

var content []byte

func (*UploadService) GetFilePath(ctx context.Context, request *proto.UploadFileRequest, response *proto.UploadFileResponse) error {

	if request.GetContent() != nil && request.GetSize() != 0 {
		content = append(content, request.Content...)
		return nil
	}

	if content == nil {
		response.Path = ""
		response.Msg = "未获取到文件内容！"
		return nil
	}

	fileName := request.GetFileName()
	fileType := request.GetFileType()

	//将文件保存至本项目指定目录中
	_ = ioutil.WriteFile("./upload/file/"+fileName+fileType, content, 0666)
	content = nil
	//上传又拍云
	dst, result := upload.StartUpload("/datum/datum/"+fileName+fileType, "./upload/file/"+fileName+fileType)

	response.Path = dst
	response.Msg = result
	return nil

}

func (*UploadService) CheckText(ctx context.Context, request *proto.CheckTextRequest, response *proto.CheckTextResponse) error {
	text := request.Text
	if text == "" {
		response.Msg = "未接收到文本内容！"
		response.StatusCode = 0
		return nil
	}
	//检测文本
	statusCode, msg := upload.CheckText(text)
	response.Msg = msg
	response.StatusCode = statusCode

	return nil
}
