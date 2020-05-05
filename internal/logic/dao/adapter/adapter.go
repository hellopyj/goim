package adapter

import (
	"errors"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/gogo/protobuf/proto"
	"strconv"
)
var(
	_daomap=make(map[string]func()PushMsg)
)
type PushMsg interface {
	PublishMessage(pbmsg *pb.PushMsg) (err error )
	Close() (err error)
}
//注册消息插件
func RegAdapter(key string,pub func()PushMsg){
	_daomap[key]=pub
}
//获得消息插件
func GetAdapter(key string)(pub PushMsg){
	pubfn,has:=_daomap[key]
	if has {
		return pubfn()
	}else {
		panic("no pub")
	}
}
func FormatMsg(pbmsg *pb.PushMsg)(key string,b []byte,err error)  {
	b, err = proto.Marshal(pbmsg)
	if err != nil {
		return
	}
	switch pbmsg.Type {
	case pb.PushMsg_PUSH:
		key=pbmsg.Keys[0]
	case pb.PushMsg_ROOM:
		key=pbmsg.Room
	case pb.PushMsg_BROADCAST:
		key=strconv.FormatInt(int64(pbmsg.Operation), 10)
	default:
		err=errors.New("formatmsg no key")
	}
	return
}
