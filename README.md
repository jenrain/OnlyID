# OnlyID

## 简介

onlyId是一个分布式id生成器，用户可以根据业务需要基于数据库号段、雪花算法或者redis的incr命令来生成id。
服务端多机部署提供高可用服务，多节点竞争etcd分布式锁成为主，对外提供GRPC和HTTP服务。

## 特性

- 三种分布式id生成算法：数据库号段算法、雪花算法、redis incr命令
- 数据库号段算法采用预分配机制，分配id只访问内存
- 基于etcd lease实现自动抢主，主节点挂掉从节点自动申请为主节点
- 对外提供GRPC和HTTP服务
- 高性能，单机qps可达十万级

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

### 环境

| Parameter  | Value                             |
|------------|-----------------------------------|
| Go version | 1.20                              |
| Machine    | 腾讯云CVM 计算型C6                      |
| System     | Centos 7.6                        |
| CPU        | Intel Ice Lake(3.2GHz/3.5Ghz) 16核 |
| Memory     | 32 GB                             |

### ghz压测grpc接口

**GetId**
```
Summary:
  Count:        2677242
  Total:        20.00 s
  Slowest:      12.84 ms
  Fastest:      0.06 ms
  Average:      1.18 ms
  Requests/sec: 133854.88
```

**GetSnowFlakeId**
```
Summary:
  Count:        3021127
  Total:        20.00 s
  Slowest:      10.80 ms
  Fastest:      0.05 ms
  Average:      0.99 ms
  Requests/sec: 151052.31
```

**GetRedisId**
```
Summary:
  Count:        2276938
  Total:        20.00 s
  Slowest:      1.01 s
  Fastest:      0.09 ms
  Average:      1.49 ms
  Requests/sec: 113841.71
```

### wrk压测http接口

**getId**
```
  8 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   103.28ms  136.97ms 560.60ms   79.19%
    Req/Sec     3.56k     3.45k   23.74k    83.42%
  Latency Distribution
     50%    3.67ms
     75%  203.58ms
     90%  336.79ms
     99%  427.25ms
  586778 requests in 30.03s, 76.66MB read
Requests/sec:  19542.88
Transfer/sec:      2.55MB
```

**getSnowFlakeId**
```
  8 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   105.31ms  140.31ms 614.58ms   79.69%
    Req/Sec     3.25k     2.93k   22.98k    85.60%
  Latency Distribution
     50%    3.65ms
     75%  204.15ms
     90%  341.56ms
     99%  460.46ms
  588146 requests in 30.04s, 83.01MB read
Requests/sec:  19580.94
Transfer/sec:      2.76MB
```

**getRedisId**
```
  8 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   104.76ms  143.57ms   1.46s    80.45%
    Req/Sec     3.31k     2.86k   18.94k    77.13%
  Latency Distribution
     50%    3.61ms
     75%  200.63ms
     90%  335.55ms
     99%  486.03ms
  569566 requests in 30.03s, 73.77MB read
Requests/sec:  18969.31
Transfer/sec:      2.46MB
```

## 参考

[Leaf—美团点评分布式ID生成系统](https://tech.meituan.com/2017/04/21/mt-leaf.html)
