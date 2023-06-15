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
	//log.GetLogger().Info("received a request from the client || GetId", zap.Any("biz_tag", in.BizTag))
	if s.srv.R == nil {
		return &onlyIdSrv.ResId{Id: -1, Message: "service is unavailable"}, nil
	}
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
	//log.GetLogger().Info("received a request from the client || GetSnowFlakeId")
	if s.srv.SnowFlake == nil {
		return &onlyIdSrv.ResId{Id: -1, Message: "service is unavailable"}, nil
	}
	return &onlyIdSrv.ResId{
		Id: s.srv.SnowFlakeGetId(),
	}, nil
}

func (s Server) GetRedisId(ctx context.Context, in *onlyIdSrv.ReqId) (*onlyIdSrv.ResId, error) {
	//log.GetLogger().Info("received a request from the client || GetRedisId", zap.Any("biz_tag", in.BizTag))
	if s.srv.Cache == nil {
		return &onlyIdSrv.ResId{Id: -1, Message: "service is unavailable"}, nil
	}
	if in.GetBizTag() == "" {
		return nil, errors.New("biz_tag is empty")
	}
	id, err := s.srv.RedisGetId(in.GetBizTag())
	if err != nil {
		log.GetLogger().Error("get id failed", zap.Error(err))
		return &onlyIdSrv.ResId{Id: -1, Message: "service is unavailable"}, nil
	}
	return &onlyIdSrv.ResId{Id: id}, nil
}
