package repository

import (
	"github.com/jenrain/OnlyID/entity"
	"github.com/jenrain/OnlyID/library/log"
	"github.com/jenrain/OnlyID/library/tool"
	"go.uber.org/zap"
)

// SegmentsGetAll 只获取最近六小时的所有id
func (r *Repository) SegmentsGetAll() (res []entity.Segments, err error) {
	if err = r.db.Table("segments").Where("update_time >= ?", tool.GetTimeUnix()-21600).Find(&res).Error; err != nil {
		log.GetLogger().Error("[Repository] SegmentGetAll Find", zap.Error(err))
	}
	return
}
