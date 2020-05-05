package handler

import (
	"context"
	"errors"
	"github.com/Terry-Mao/goim/dpg/proto/ghandler"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

var (
	// grpc options
	grpcKeepAliveTime    = time.Duration(10) * time.Second
	grpcKeepAliveTimeout = time.Duration(3) * time.Second
	grpcBackoffMaxDelay  = time.Duration(3) * time.Second
	grpcMaxSendMsgSize   = 1 << 24
	grpcMaxCallMsgSize   = 1 << 24
)

const (
	// grpc options
	grpcInitialWindowSize     = 1 << 24
	grpcInitialConnWindowSize = 1 << 24
)

type HandlerClient struct {
	client ghandler.HandlerServerClient
}
var _clients map[string]*HandlerClient

//func getCtx() (ctx context.Context) {
//	md := metadata.Pairs("platform", "goim")
//	ctx = metadata.NewOutgoingContext(context.Background(), md)
//	return
//}
func New(c *conf.Config) {
	hmap:=HandlerMap{c:c}
	hmap.watchHandler()
}
func newHandler(url string)(hand *HandlerClient,err error)  {
	if len(url)>0 {
		//creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second))
		defer cancel()
		conn, e := grpc.DialContext(ctx,url, []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithInitialWindowSize(grpcInitialWindowSize),
			grpc.WithInitialConnWindowSize(grpcInitialConnWindowSize),
			grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(grpcMaxCallMsgSize)),
			grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(grpcMaxSendMsgSize)),
			grpc.WithBackoffMaxDelay(grpcBackoffMaxDelay),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                grpcKeepAliveTime,
				Timeout:             grpcKeepAliveTimeout,
				PermitWithoutStream: true,
			})}...)
		if e == nil {
			hand=&HandlerClient{
				client:ghandler.NewHandlerServerClient(conn),
			}
			return
		}else {
			log.Println(e)
		}
		err=errors.New("Handler初始化失败")
	}
	return

}
func Handler(key string,betreq *ghandler.HandlerReq)(replay *ghandler.HandlerReply,err error)  {
	replay=&ghandler.HandlerReply{}
	handlerClient:=getHandler(key)
	if handlerClient==nil {
		err=errors.New("hander_server为空")
		return
	}
	gret,err:=(*handlerClient).client.Handler(context.Background(),betreq)
	if err==nil {
		replay=gret
	}
	return
}
func getHandler(key string)*HandlerClient  {
	client,has:=_clients[key]
	if has {
		return client
	}
	return nil
}


