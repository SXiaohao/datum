package utils

/**
 * @Author: sxiaohao
 * @Description:
 * @File:  client
 * @Version: 1.0.0
 * @Date: 2020/10/28 下午9:40
 */

import (
	"datum_textbook/config"
	proto4 "datum_textbook/proto/qrcode"
	proto3 "datum_textbook/proto/sms"
	proto2 "datum_textbook/proto/token"
	proto "datum_textbook/proto/upload"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var (
	service micro.Service
)

func init() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.GetString("etcd.addr")),
	)
	service = micro.NewService(
		micro.Registry(etcdReg))
}

func GetUploadClient() proto.UploadService {

	return proto.NewUploadService(config.GetString("etcd.serviceName"), service.Client())
}

func GetTokenClient() proto2.TokenService {

	return proto2.NewTokenService(config.GetString("etcd.serviceName"), service.Client())
}
func GetSmsClient() proto3.SmsService {

	return proto3.NewSmsService(config.GetString("etcd.serviceName"), service.Client())
}

func GetQrcodeClient() proto4.QrcodeService {
	return proto4.NewQrcodeService(config.GetString("etcd.serviceName"), service.Client())
}
