package tool

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits uint8 = 10 // 每台机器的id位数 10位最大可以用2 ^ 10 = 1024个字节
	numberBits uint8 = 22 // 表示每个集群下的每个节点，1秒内可生成的id序号的二进制位数 即每秒可生成 2^22-1=4194304个唯一id(0-4194303)

	workerMax   int64 = ^(-1 << workerBits)     // 10位节点id的最大值
	numberMax   int64 = ^(-1 << numberBits)     // 1秒内可以生成的id序号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 节点id向左的偏移量
	// 31位字节作为时间戳数值的话 大约68年就会用完
	// 假如你2010年1月1日开始开发系统 如果不减去2010年1月1日的时间戳 那么白白浪费40年的时间戳啊！
	// 这个一旦定义且开始生成ID后千万不要改了 不然可能会生成相同的ID
	epoch int64 = 1594364131 //这个是我在写epoch这个变量时的时间戳(秒)
)

type Worker struct {
	mu        sync.RWMutex // 添加互斥锁 确保并发安全
	timestamp int64        // 记录时间戳
	workerId  int64        // 该节点的ID
	number    int64        // 当前毫秒已经生成的id序列号（从0开始累加）1秒内最多生成4194304个id
}

func NewWorker(workerId int64) (*Worker, error) {
	// 是否溢出
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("workerId is overflow")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *Worker) GetId() int64 {
	// 解决并发安全
	w.mu.Lock()
	defer w.mu.Unlock()
	// 获取生成时的时间戳
	now := w.Now()
	// 在当前秒内
	if now == w.timestamp {
		w.number++
		// 判断当前工作节点是否在1秒内已经生成numberMax个id
		if w.number > numberMax {
			// 如果当前工作节点在1秒内生成的id已经超过上限 需要等待1秒再继续生成
			w.number = 0
			for now <= w.timestamp {
				now = w.Now()
			}
			w.timestamp = now
		}
	} else if now > w.timestamp {
		// 超过当前的秒，重置工作节点生成id的序号
		w.number = 0
		// 将机器上一次生成id的时间更新为当前时间
		w.timestamp = now
	} else {
		// 时钟回拨
		for now < w.timestamp {
			now = w.Now()
		}
		// 等待机器时钟回拨到内存中记录的时间戳
		w.number++
		// 判断当前工作节点是否在1秒内已经生成numberMax个id
		if w.number > numberMax {
			w.number = 0
			for now <= w.timestamp {
				now = w.Now()
			}
			w.timestamp = now
		}
	}
	id := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.number))
	return id
}

func (w *Worker) Now() int64 {
	return time.Now().Unix()
}
