package imp

import (
	"context"
	"datum_common/config"
	proto "datum_common/proto/token"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenService struct {
}

func (*TokenService) GetToken(c context.Context, req *proto.GetTokenRequest, rsp *proto.GetTokenResponse) error {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": req.Uid,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	})
	token, err := at.SignedString([]byte(config.GetString("jwt.secret")))
	if err != nil {
		rsp.Msg = "fail"
		return err
	}
	rsp.Token = token
	rsp.Msg = "success"
	return nil
}

func (*TokenService) ParseToken(c context.Context, req *proto.ParseTokenRequest, rsp *proto.ParseTokenResponse) error {
	claim, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.secret")), nil
	})
	if err != nil {
		rsp.Msg = "fail"
		return err
	}
	rsp.Uid = uint32(claim.Claims.(jwt.MapClaims)["uid"].(float64))
	return nil
}
