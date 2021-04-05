package imp

import (
	"context"
	proto "datum_common/proto/sms"
	"datum_common/sms/lib"
)

type SmsService struct {

}

func (*SmsService) SendRegisterCode(c context.Context, req *proto.SendRegisterCodeRequest,rsp *proto.SendRegisterCodeResponse) error  {
	err := lib.SendRegisterCode(req.Code,req.PhoneNumber)
	if err != nil {
		rsp.Message = err.Error()
		return err
	}
	rsp.Message = "success"
	return nil
}

func (*SmsService) SendRestPwdCode(c context.Context, req *proto.SendRestPwdCodeRequest, rsp *proto.SendRestPwdCodeResponse) error  {
	err := lib.SendRestPwdCode(req.Code,req.PhoneNumber)
	if err != nil {
		rsp.Message = err.Error()
		return err
	}
	rsp.Message = "success"
	return nil
}