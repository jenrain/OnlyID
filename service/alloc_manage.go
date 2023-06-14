package service

import (
	"context"
	"errors"
	"github.com/jenrain/OnlyID/entity"
	"sync"
	"time"
)

// Alloc 缓存在内存中的所有id
type Alloc struct {
	Mu        sync.RWMutex
	BizTagMap map[string]*BizAlloc
}

// BizAlloc 对应一种Biz的所有id
type BizAlloc struct {
	Mu      sync.Mutex
	BizTag  string
	IdArray []*IdArray
	GetDB   bool // 并发控制，标记当前是否在读取DB
}

// IdArray id列表
type IdArray struct {
	Cur   int64
	Start int64
	End   int64
}

// NewAllocId 将DB中的id都加载进内存
func (s *Service) NewAllocId() (a *Alloc, err error) {
	var res []entity.Segments
	// 加载DB中的号段
	if res, err = s.r.SegmentsGetAll(); err != nil {
		return
	}
	a = &Alloc{
		BizTagMap: make(map[string]*BizAlloc),
	}
	// 根据DB中的号段，在内存中构造所有Biz的id列表
	for _, v := range res {
		a.BizTagMap[v.BizTag] = &BizAlloc{
			BizTag:  v.BizTag,
			IdArray: make([]*IdArray, 0),
			GetDB:   false,
		}
		a.BizTagMap[v.BizTag].IdArray = append(a.BizTagMap[v.BizTag].IdArray, &IdArray{Start: v.MaxId, End: v.MaxId + v.Step})
	}
	s.alloc = a
	return
}

func (b *BizAlloc) GetId(s *Service) (id int64, err error) {
	var (
		// 标记是否可以获取到id
		canGetId    bool
		ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	)
	// 获取id需要加锁
	b.Mu.Lock()
	// 当前号段还有id
	if b.LeftIdCount() > 0 {
		id = b.PopId()
		canGetId = true
	}
	// 如果内存中缓存的id号段只剩下一个，就去DB中加载
	if len(b.IdArray) <= 1 && !b.GetDB {
		b.GetDB = true
		b.Mu.Unlock()
		// 异步去数据库加载新号段
		go b.GetIdArray(cancel, s)
	} else {
		b.Mu.Unlock()
		defer cancel()
	}
	// 成功从内存获取到了id，直接返回
	if canGetId {
		return
	}
	// 内存中的id耗尽，使用select阻塞，直到从DB获取到新的号段
	select {
	case <-ctx.Done():
	}
	b.Mu.Lock()
	if b.LeftIdCount() > 0 {
		id = b.PopId()
	} else {
		err = errors.New("no get id")
	}
	b.Mu.Unlock()
	return
}

// GetIdArray 去数据库中加载新的号段
func (b *BizAlloc) GetIdArray(cancel context.CancelFunc, s *Service) {
	var (
		tryNum int
		ids    *entity.Segments
		err    error
	)
	defer cancel()
	for {
		// 失败重试3次
		if tryNum >= 3 {
			b.GetDB = false
			break
		}
		b.Mu.Lock()
		// 再次检查号段是否只剩下一个
		if len(b.IdArray) <= 1 {
			b.Mu.Unlock()
			ids, err = s.r.SegmentsIdNext(b.BizTag)
			if err != nil {
				tryNum++
			} else {
				tryNum = 0
				b.Mu.Lock()
				b.IdArray = append(b.IdArray, &IdArray{
					Cur:   0,
					Start: ids.MaxId,
					End:   ids.MaxId + ids.Step,
				})
				if len(b.IdArray) > 1 {
					b.GetDB = false
					b.Mu.Unlock()
					break
				} else {
					b.Mu.Unlock()
				}
			}
		} else {
			b.Mu.Unlock()
		}
	}
}

// LeftIdCount 获取当前号段的剩余id个数
func (b *BizAlloc) LeftIdCount() (count int64) {
	for _, v := range b.IdArray {
		arr := v
		// 结束位置 - 开始位置 - 已经分配的次数
		count += arr.End - arr.Start - arr.Cur
	}
	return
}

func (b *BizAlloc) PopId() (id int64) {
	id = b.IdArray[0].Start + b.IdArray[0].Cur // 开始位置加上分配次数
	b.IdArray[0].Cur++                         // 分配次数 + 1
	if id+1 >= b.IdArray[0].End {              // 当前号段已经没有id了
		b.IdArray = append(b.IdArray[:0], b.IdArray[1:]...) // 把分配完的数组移除
	}
	return
}
