package repository

import (
	"errors"
	"fmt"
	"github.com/jenrain/OnlyID/entity"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"go.uber.org/zap"
)

// SegmentsCreate 创建新号段
func (r *Repository) SegmentsCreate(s *entity.Segments) (data *entity.Segments, err error) {
	data = new(entity.Segments)
	var count int64
	if err = r.db.Table("segments").Where("biz_tag = ?", s.BizTag).Count(&count).Error; err != nil {
		log.GetLogger().Error("[SegmentsCreate] Find", zap.Any("req", s), zap.Error(err))
		err = errors.New("find biz_tag failed")
		return
	}
	if count > 0 {
		// 已经存在这个biz_type的号段
		log.GetLogger().Error("[SegmentsCreate] Already Exist", zap.Any("req", s), zap.Error(err))
		fmt.Println("data: ", data)
		err = errors.New("biz_tag already exist")
		return
	}
	// 如果biz_tag原来不存在，就创建一个
	s.CreateTime = tool.GetTimeUnix()
	s.UpdateTime = tool.GetTimeUnix()
	if err = r.db.Table("segments").Create(s).Error; err != nil {
		log.GetLogger().Error("[SegmentsCreate] Create", zap.Any("req", s), zap.Error(err))
		err = errors.New("create biz_tag failed")
		return
	}
	data = s
	return
}
