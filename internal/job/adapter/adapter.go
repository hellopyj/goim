package adapter

import (
	"context"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
)
type consume interface {
	Push(ctx context.Context, pushMsg *pb.PushMsg) (err error)
}
type JobConsumer interface {
	Consume(j consume)
	Close() (err error)
}
var(
	_daomap=make(map[string]func()JobConsumer)
)
//注册消息插件
func RegAdapter(key string,pub func()JobConsumer){
	_daomap[key]=pub
}
//获得消息插件
func GetAdapter(key string)(pub JobConsumer){
	pubfn,has:=_daomap[key]
	if has {
		return pubfn()
	}else {
		panic("no pub")
	}
}
