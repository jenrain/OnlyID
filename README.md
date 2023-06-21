# OnlyID

## 简介

onlyId是一个分布式id生成器，用户可以根据业务需要基于数据库号段、雪花算法或者redis的incr命令来生成id。
服务端多机部署提供高可用服务，多节点竞争etcd分布式锁成为主，GRPC对外提供服务。

## 特性

- 三种分布式id生成算法：数据库号段算法、雪花算法、redis incr命令
- 数据库号段算法采用预分配机制，分配id只访问内存
- 基于etcd lease实现自动抢主，主节点挂掉从节点自动申请为主节点
- GRPC对外提供服务
- 高性能，qps可达百万级

## 安装部署

### 部署服务端

```bash
# 1.克隆代码
git clone https://github.com/jenrain/OnlyID.git

# 2.根据需要启动mysql或者redis服务

# 3.启动etcd服务

# 3.本地编译，在项目根目录执行make命令，交叉编译项目

# 4.将onlyId二进制文件和only_id.toml文件上传到linux服务器同一目录下

# 5.根据业务需要，修改配置文件的id生成模式(mode字段)

# 6.运行项目
./onlyId -conf only_id.toml &
```

### 客户端调用

1. 以sdk的形式安装
```bash
go get -u github.com/jenrain/OnlyID
```

2. 同步依赖
```bash
go mod tidy
```
3. 客户端调用

```go
package main

import (
	"context"
	"fmt"
	onlyIdSrv "github.com/jenrain/OnlyID/api"
	"github.com/jenrain/OnlyID/client"
)

func main() {
	// 1.访问etcd获取节点ip地址
	addr := []string{"127.0.0.1:2379"}
	cli, _ := client.InitGrpc(addr, 15)
	c, _ := cli.GetOnlyIdGrpcClient()

	// 2.数据库号段id
	res, err := c.GetId(context.TODO(), &onlyIdSrv.ReqId{
		BizTag: "001",
	})
	if err != nil {
		fmt.Println("get id fail || reason: ", err.Error())
	}
	fmt.Println(res.Id, res.Message)

	// 3.雪花算法id
	//flakeRes, err := c.GetSnowFlakeId(context.TODO(), &empty.Empty{})
	//if err != nil {
	//	fmt.Println("get flakeRes fail || reason: ", err.Error())
	//}
	//fmt.Println(flakeRes.Id, res.Message)

	// 4.数据库号段id
	//res, err := c.GetRedisId(context.TODO(), &onlyIdSrv.ReqId{
	//	BizTag: "001",
	//})
	//if err != nil {
	//	fmt.Println("get id fail || reason: ", err.Error())
	//}
	//fmt.Println(res.Id, res.Message)
}
```

## 压测

TODO

## 参考

[Leaf—美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)

[gid](https://github.com/hwholiday/gid)