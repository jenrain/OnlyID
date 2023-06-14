package repository

import (
	"github.com/jenrain/OnlyID/entity"
	"github.com/jenrain/OnlyID/library/log"
	"go.uber.org/zap"
)

func (r *Repository) SegmentsGetAll() (res []entity.Segments, err error) {
	if err = r.db.Table("segments").Find(&res).Error; err != nil {
		log.GetLogger().Error("[Repository] SegmentGetAll Find", zap.Error(err))
	}
	return
}
