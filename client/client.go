package client

import (
	"context"
	onlyIdSrv "github.com/jenrain/OnlyID/api"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	etcd     *clientv3.Client
	node     string
	key      string
	change   bool
	ttl      int64
	revision int64
	conn     *grpc.ClientConn
}

func InitGrpc(etcdAddr []string, ttl int64) (*Client, error) {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   etcdAddr,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	cli := &Client{
		etcd:   c,
		key:    "/onlyId/master",
		change: false,
		ttl:    ttl,
	}
	cli.masterCheckLoop()
	return cli, nil
}

func (c *Client) watch() {
	watcher := clientv3.NewWatcher(c.etcd)
	watchChan := watcher.Watch(context.Background(), c.key, clientv3.WithRev(c.revision+1))
	for watchRes := range watchChan {
		for _, event := range watchRes.Events {
			switch event.Type {
			case mvccpb.DELETE:
				go c.getMasterNode()
			}
		}
	}
}

func (c *Client) masterCheckLoop() {
	if err := c.getMasterNode(); err != nil {
		panic(err)
	}
	go c.watch()
	ticker := time.NewTicker(time.Duration(c.ttl) * time.Second)
	go func() {
		for {
			select {
			case _ = <-ticker.C:
				_ = c.getMasterNode()
			}
		}
	}()
}

func (c *Client) getMasterNode() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	res, err := c.etcd.Get(ctx, c.key)
	if err != nil {
		return err
	}
	for _, v := range res.Kvs {
		if string(v.Key) == c.key {
			newNode := string(v.Value)
			if c.node != newNode {
				c.change = true
			}
			c.node = string(v.Value)
		}
	}
	if c.revision != res.Header.Revision {
		c.revision = res.Header.Revision
	}
	return nil
}

func (c *Client) GetOnlyIdGrpcClient() (onlyIdSrv.OnlyIdClient, error) {
	var err error
	// 查看服务端ip是否发生变动
	if c.change {
		c.conn, err = grpc.Dial(c.node, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}
		c.change = false
	}
	return onlyIdSrv.NewOnlyIdClient(c.conn), nil
}
