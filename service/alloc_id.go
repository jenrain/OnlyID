package service

import (
	"github.com/jenrain/OnlyID/entity"
)

func (s *Service) GetId(tag string) (id int64, err error) {
	// 获取id的时候要加锁
	s.alloc.Mu.Lock()
	defer s.alloc.Mu.Unlock()

	val, ok := s.alloc.BizTagMap[tag]
	// 内存中没有这种biz_tag的号段
	if !ok {
		if err = s.CreateTag(&entity.Segments{
			BizTag: tag,
			MaxId:  1,
			Step:   1000,
		}); err != nil {
			return 0, err
		}
		val, _ = s.alloc.BizTagMap[tag]
	}
	return val.GetId(s)
}

// CreateTag 创建号段
func (s *Service) CreateTag(e *entity.Segments) error {
	data, err := s.r.SegmentsCreate(e)
	if err != nil {
		return err
	}
	b := &BizAlloc{
		BizTag:  e.BizTag,
		IdArray: make([]*IdArray, 0),
		GetDB:   false,
	}
	b.IdArray = append(b.IdArray, &IdArray{
		Start: data.MaxId,
		End:   data.MaxId + data.Step,
	})
	s.alloc.BizTagMap[e.BizTag] = b
	return nil
}
