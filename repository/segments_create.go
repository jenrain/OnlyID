package repository

import (
	"OnlyID/entity"
	"OnlyID/library/log"
	"OnlyID/library/tool"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// SegmentsCreate 创建新号段
func (r *Repository) SegmentsCreate(s *entity.Segments) (data *entity.Segments, err error) {
	data = new(entity.Segments)
	if err = r.db.Table("segments").Where("biz_tag = ?", s.BizTag).Find(data).Error; err != nil {
		// 如果biz_tag原来不存在，就创建一个
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.CreateTime = tool.GetTimeUnix()
			s.UpdateTime = tool.GetTimeUnix()
			if err = r.db.Table("segments").Create(s).Error; err != nil {
				log.GetLogger().Error("[SegmentsCreate] Create", zap.Any("req", s), zap.Error(err))
				err = errors.New("create biz_tag failed")
				return
			}
			data = s
			return
		} else {
			log.GetLogger().Error("[SegmentsCreate] Find", zap.Any("req", s), zap.Error(err))
			err = errors.New("find biz_tag failed")
			return
		}
	}
	// 已经存在这个biz_type的号段
	log.GetLogger().Error("[SegmentsCreate] Already Exist", zap.Any("req", s), zap.Error(err))
	err = errors.New("biz_tag already exist")
	return
}
