package main

import (
	"douyin/cmd/api/rpc"
	"douyin/cmd/user/dal"
	"douyin/kitex_gen/user/userservice"
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
)

const serviceAddr = "127.0.0.1:9096"

func main() {
	logrus.InitLogger("user_log.txt")
	dal.Init()
	rpc.InitRPC()
	tracerSuit, closer := tracer.Init(constants.UserServiceName)
	defer closer.Close()

	r, err := etcd.NewEtcdRegistry([]string{
		constants.EtcdAddress,
	})

	addr, err := net.ResolveTCPAddr("tcp", serviceAddr)

	svr := userservice.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithMiddleware(mw.CommonMiddleware),                                                     // middleware
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithServiceAddr(addr), // address
		server.WithSuite(tracerSuit),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithRegistry(r),                                             // registry
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
