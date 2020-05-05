package push

import (
	"context"
	"fmt"
	"github.com/Terry-Mao/goim/dpg/proto/gpush"
	"github.com/Terry-Mao/goim/internal/logic"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)
const (
	RET_SUCCESS=1
)
func New(c *conf.Config,l *logic.Logic)  {
	urls:=c.Xserver.Push.Urls
	if  len(urls)>0{
		_push_server:=PushService{logic:l}
		listen, err := net.Listen("tcp", urls[0])
		if err != nil {
			fmt.Printf("RPC监听失败:%v", err)
		}
		keepParams := grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     time.Duration(time.Second * 60),
			MaxConnectionAgeGrace: time.Duration(time.Second * 20),
			Time:             time.Duration(time.Second * 60),
			Timeout:          time.Duration(time.Second * 20),
			MaxConnectionAge: time.Duration(time.Hour * 2),
		})
		grs:= grpc.NewServer(keepParams)
		gpush.RegisterPushDealServer(grs, _push_server)
		go func() {
			if err:=grs.Serve(listen);err!=nil{
				fmt.Println("RPC关闭:",err.Error())
			}
		}()
	}
}
type PushService struct{
	logic  *logic.Logic
}
func (p PushService) PushMids(ctx context.Context, in *gpush.PushMidsReq) (ret *gpush.RetData, err error) {
	ret=&gpush.RetData{}
	if err = p.logic.PushMids(context.TODO(), int32(in.Op), in.Mids, []byte(in.Data)); err != nil {
		ret.Summary=err.Error()
		return
	}
	ret.Code=RET_SUCCESS
	ret.Data="发送成功"
	return
}
func (p PushService) PushKeys(ctx context.Context, in *gpush.PushKeysReq) (ret *gpush.RetData, err error) {
	ret=&gpush.RetData{}
	if err = p.logic.PushKeys(context.TODO(), int32(in.Op), in.Keys, []byte(in.Data)); err != nil {
		ret.Summary=err.Error()
		return
	}
	ret.Code=RET_SUCCESS
	ret.Data="发送成功"
	return
}
func (p PushService) PushRoom(ctx context.Context, in *gpush.PushRoomReq) (ret *gpush.RetData, err error) {
	ret=&gpush.RetData{}
	if err = p.logic.PushRoom(context.TODO(), int32(in.Op), in.Type,in.Room, []byte(in.Data)); err != nil {
		ret.Summary=err.Error()
		return
	}
	ret.Code=RET_SUCCESS
	ret.Data="发送成功"
	return
}
func (p PushService) PushAll(ctx context.Context, in *gpush.PushAllReq) (ret *gpush.RetData, err error) {
	ret=&gpush.RetData{}
	if err = p.logic.PushAll(context.TODO(), int32(in.Op),int32(in.Speed), []byte(in.Data)); err != nil {
		ret.Summary=err.Error()
		return
	}
	ret.Code=RET_SUCCESS
	ret.Data="发送成功"
	return
}
