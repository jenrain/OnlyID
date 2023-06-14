package grpc

import (
	"context"
	"errors"
	onlyIdSrv "github.com/jenrain/OnlyID/api"
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Server) GetId(ctx context.Context, in *onlyIdSrv.ReqId) (*onlyIdSrv.ResId, error) {
	if in.GetBizTag() == "" {
		return nil, errors.New("biz_tag is empty")
	}
	id, err := s.srv.GetId(in.GetBizTag())
	if err != nil {
		log.GetLogger().Error("get id failed", zap.Error(err))
		if err.Error() == "biz_tag already exist" {
			return &onlyIdSrv.ResId{Id: -1, Message: err.Error()}, nil
		}
		return &onlyIdSrv.ResId{Id: -1, Message: "service is unavailable"}, nil
	}
	return &onlyIdSrv.ResId{Id: id}, nil
}

func (s Server) GetSnowFlakeId(ctx context.Context, empty *emptypb.Empty) (*onlyIdSrv.ResId, error) {
	return &onlyIdSrv.ResId{
		Id: s.srv.SnowFlakeGetId(),
	}, nil
}
