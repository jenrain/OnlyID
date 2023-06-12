package repository

import (
	"OnlyID/entity"
	"OnlyID/library/log"
	"OnlyID/library/tool"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func (r *Repository) SegmentsIdNext(tag string) (id *entity.Segments, err error) {
	r.db.Transaction(func(tx *gorm.DB) error {
		id = &entity.Segments{}
		if err = tx.Exec("update segments set max_id = max_id + step, update_time = ? where biz_tag = ?", tool.GetTimeUnix(), tag).Error; err != nil {
			log.GetLogger().Error("[Repository] SegmentsIdNext Update", zap.String("tag", tag), zap.Error(err))
			return err
		}
		if err = tx.Table("segments").Where("biz_tag = ?", tag).Find(id).Error; err != nil {
			log.GetLogger().Error("[Repository] SegmentsIdNext Find", zap.String("tag", tag), zap.Error(err))
			return err
		}
		return nil
	})
	return
}
