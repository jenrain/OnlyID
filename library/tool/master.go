package tool

import (
	"context"
	"github.com/jenrain/OnlyID/library/log"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"time"
)

type master struct {
	cli          *clientv3.Client
	ip           string
	key          string
	ttl          int64
	isMasterNode bool
	revision     int64
	id           clientv3.LeaseID
	isClose      bool
}

var MasterNode *master

func InitMasterNode(etcdAddr []string, ip string, ttl int64) error {
	MasterNode = &master{
		key: "/onlyId/master",
		ttl: ttl,
		ip:  ip,
	}
	var err error
	MasterNode.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", ip), zap.Error(err))
		return err
	}
	go MasterNode.leaseTTL()
	return nil
}

func (m *master) leaseTTL() {
	if m == nil {
		panic("InitMasterNode is nil")
	}
	if err := m.applyMasterNode(); err != nil {
		panic(err)
	}
	// 监听key兜底
	m.watch()
	// 定时续约
	ticker := time.NewTicker(time.Duration(m.ttl) * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				_ = m.applyMasterNode()
			}
		}
	}()

}

// 监听到key被删除了，尝试去抢锁成为主
func (m *master) watch() {
	go func() {
		watcher := clientv3.NewWatcher(m.cli)
		watchChan := watcher.Watch(context.Background(), m.key, clientv3.WithRev(m.revision+1))
		for watchResp := range watchChan {
			for _, event := range watchResp.Events {
				switch event.Type {
				case mvccpb.DELETE:
					if !m.isClose {
						go m.applyMasterNode()
					}
				}
			}
		}
	}()
}

// 抢锁成为master
// 或者续约master的ttl
func (m *master) applyMasterNode() error {
	if m == nil {
		panic("InitMasterNode is nil")
	}
	lease := clientv3.NewLease(m.cli)
	// slave抢锁
	if !m.isMasterNode {
		txn := clientv3.NewKV(m.cli).Txn(context.TODO())
		grantRes, err := lease.Grant(context.TODO(), m.ttl+1)
		if err != nil {
			log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", m.ip), zap.Error(err))
			m.isMasterNode = false
			return err
		}
		m.id = grantRes.ID
		txn.If(clientv3.Compare(clientv3.CreateRevision(m.key), "=", 0)).
			Then(clientv3.OpPut(m.key, m.ip, clientv3.WithLease(grantRes.ID))).
			Else()
		txnRes, err := txn.Commit()
		if err != nil {
			log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", m.ip), zap.Error(err))
			m.isMasterNode = false
			return err
		}
		if txnRes.Succeeded {
			m.isMasterNode = true
		} else {
			m.isMasterNode = false
		}
		if m.revision != txnRes.Header.Revision {
			m.revision = txnRes.Header.Revision
		}
	}
	// master续约ttl
	_, err := lease.KeepAliveOnce(context.TODO(), m.id)
	if err != nil {
		m.isMasterNode = false
		log.GetLogger().Error("[ApplyMasterNode] New", zap.Any("ip", m.ip), zap.Error(err))
		return err
	}
	return nil
}

func (m *master) CloseApplyMasterNode() {
	if m != nil {
		m.isClose = true
		if _, err := m.cli.Delete(context.Background(), m.key); err != nil {
			log.GetLogger().Error("[CloseApplyMasterNode] Delete", zap.Any("ip", m.ip), zap.Error(err))
		}
	}
}
