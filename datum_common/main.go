package main

import (
	"datum_common/config"
	qrcodeService "datum_common/proto/qrcode"
	smsService "datum_common/proto/sms"
	tokenService "datum_common/proto/token"
	uploadService "datum_common/proto/upload"
	qrcodeImp "datum_common/qrcode/imp"
	smsImp "datum_common/sms/imp"
	tokenImp "datum_common/token/imp"
	uploadImp "datum_common/upload/imp"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	//注册consul
	etcdRegistry := etcdv3.NewRegistry(
		registry.Addrs(config.GetString("etcd.addr")),
	)
	service := micro.NewService(
		micro.Name("datum_common"),
		micro.Registry(etcdRegistry),
	)
	service.Init()
	//注册token服务
	_ = tokenService.RegisterTokenServiceHandler(service.Server(), new(tokenImp.TokenService))
	//注册上传服务
	_ = uploadService.RegisterUploadHandler(service.Server(), new(uploadImp.UploadService))
	//注册短信服务
	_ = smsService.RegisterSmsServiceHandler(service.Server(), new(smsImp.SmsService))
	_ = qrcodeService.RegisterQrcodeHandler(service.Server(), new(qrcodeImp.QrcodeService))

	_ = service.Run()
}
