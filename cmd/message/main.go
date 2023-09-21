package main

import (
	"douyin/cmd/api/rpc"
	"douyin/cmd/message/dal"
	"douyin/kitex_gen/message/messageservice"
	"douyin/pkg/constants"
	"douyin/pkg/logrus"
	"douyin/pkg/mw"
	tracer "douyin/pkg/trace"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"time"
)

const serviceAddr = "127.0.0.1:9093"

func main() {
	logrus.InitLogger("message_log.txt")
	dal.Init()
	rpc.InitRPC()
	tracerSuit, closer := tracer.Init(constants.MessageServiceName)
	defer closer.Close()

	r, err := etcd.NewEtcdRegistry([]string{
		constants.EtcdAddress,
	})

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)

	svr := messageservice.NewServer(new(MessageServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.MessageServiceName}), // server name
		server.WithMiddleware(mw.CommonMiddleware),                                                        // middleware
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithSuite(tracerSuit),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithRegistry(r),                                             // registry
		server.WithReadWriteTimeout(10*time.Minute),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
