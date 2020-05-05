package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/Terry-Mao/goim/dpg/bean"
	"github.com/Terry-Mao/goim/dpg/proto/gcmd"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"time"
)

const (
	LE_PLATFORM="leapp"
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
//鉴权========
type AuthItem struct {
	Token string
}
//鉴权========

var _cmdclient gcmd.CmdDealClient

func getCtx() (ctx context.Context) {
	md := metadata.Pairs("platform", "goim")
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	return
}
func New(c *conf.Config,) {
	urls:=c.Xserver.Handler.Urls
	if len(urls)>0 {
		//creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
		conn, err := grpc.Dial(urls[0], []grpc.DialOption{
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
		if err != nil {
			fmt.Println(err)
		}
		//defer conn.Close()
		_cmdclient= gcmd.NewCmdDealClient(conn)
	}

}
func Cmd(cmd []byte,uid int,token string,roomid int)(ret *gcmd.RetData,err error)  {
	if _cmdclient==nil {
		err=errors.New("Cmd_server为空")
		return
	}
	ret=&gcmd.RetData{}
	betreq:=&gcmd.CmdReq{
		Token:token,
		Uid:int64(uid),
		Roomid:int64(roomid),
		Command:cmd,
	}
	gret,err:=_cmdclient.Cmd(getCtx(),betreq)
	if err==nil {
		ret=gret
	}
	return
}
func Auth(uid int,token string,room bean.Room )(ret *gcmd.RetData,err error) {
	if _cmdclient==nil {
		err=errors.New("Cmd_server为空")
		return
	}
	ret=&gcmd.RetData{}
	authreq:=&gcmd.AuthReq{
		Token:token,
		Roomid:int64(room.ID),
		Uid:int64(uid),
		Platform:LE_PLATFORM,
	}
	gret, err := _cmdclient.Auth(getCtx(), authreq)
	if err==nil {
		ret=gret
	}
	return
}
