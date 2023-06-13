package grpc

import (
	onlyIdSrv "OnlyID/api"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s Server) GetId(ctx context.Context, in *onlyIdSrv.ReqId) (*onlyIdSrv.ResId, error) {
	if in.GetBizTag() == "" {
		return nil, errors.New("biz_tag is empty")
	}
	id, err := s.srv.GetId(in.GetBizTag())
	if err != nil {
		return nil, err
	}
	return &onlyIdSrv.ResId{Id: id}, nil
}

func (s Server) GetSnowFlakeId(ctx context.Context, empty *emptypb.Empty) (*onlyIdSrv.ResId, error) {
	return &onlyIdSrv.ResId{
		Id: s.srv.SnowFlakeGetId(),
	}, nil
}
