package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Terry-Mao/goim/api/comet/grpc"
	"github.com/Terry-Mao/goim/dpg/bean"
	"github.com/Terry-Mao/goim/dpg/bean/msg"
	"github.com/Terry-Mao/goim/dpg/bean/msg/chat"
	"github.com/Terry-Mao/goim/dpg/bean/msg/cmd"
	"github.com/Terry-Mao/goim/dpg/bean/msg/cmd/toast"
	"github.com/Terry-Mao/goim/dpg/config/operation"
	"github.com/Terry-Mao/goim/dpg/proto/ghandler"
	"github.com/Terry-Mao/goim/dpg/rpc/server/handler"
	"github.com/Terry-Mao/goim/dpg/tools/cast"
	"github.com/Terry-Mao/goim/internal/logic/model"
	log "github.com/golang/glog"
	"github.com/google/uuid"
	"strconv"
	"time"
)

// Connect connected a conn.
func (l *Logic) Connect(c context.Context, server, cookie string, token []byte) (mid int64, key, roomID string, accepts []int32, hb int64, err error) {
	var params struct {
		Mid      int64   `json:"mid"`
		Key      string  `json:"key"`
		RoomID   string  `json:"room_id"`
		Platform string  `json:"platform"`
		Accepts  []int32 `json:"accepts"`
		Token  string `json:"token"`
	}
	if err = json.Unmarshal(token, &params); err != nil {
		log.Errorf("json.Unmarshal(%s) error(%v)", token, err)
		return
	}
	//xlz==========================
	//鉴权
	if len(params.Accepts)<=0||params.Accepts[0]<=0 {
		err=errors.New("链接失败:无Accepts")
		return
	}
	rtype,rid,err:=model.DecodeRoomKey(params.RoomID)
	if err!=nil {
		err=errors.New("链接失败:Romm解析失败")
		return
	}
	roomid,_:=strconv.Atoi(rid)
	room:=bean.Room{Type:rtype,ID:roomid}
	//ret:=ghandler.HandlerReply{Data:"sasa"}
	//ret,err:= cmd.Auth(int(params.Mid),params.Token,room)
	//if err!=nil||ret.Code<=0 {
	//	err=errors.New("链接失败:"+ret.Summary)
	//	return
	//}
	cinfo:=bean.ConInfo{Token:params.Token,Platform:params.Platform,Accepts:params.Accepts,Room:room,Uid:int(params.Mid),CreateTime:time.Now().Unix()}
	err=l.dao.SetConInfo(int(params.Mid),cinfo)
	if err!=nil {
		err=errors.New("设置Token失败")
		return
	}
	params.Accepts=append(params.Accepts,operation.OpChat,operation.OpCmd,grpc.OpRaw,operation.OpClose)
	//xlz==========================
	mid = params.Mid
	roomID = params.RoomID
	accepts = params.Accepts
	hb = int64(l.c.Node.Heartbeat) * int64(l.c.Node.HeartbeatMax)
	if key = params.Key; key == "" {
		key = uuid.New().String()
	}
	key=params.Token
	if err = l.dao.AddMapping(c, mid, key, server); err != nil {
		log.Errorf("l.dao.AddMapping(%d,%s,%s) error(%v)", mid, key, server, err)
	}
	log.Infof("conn connected key:%s server:%s mid:%d token:%s", key, server, mid, token)
	return
}

// Disconnect disconnect a conn.
func (l *Logic) Disconnect(c context.Context, mid int64, key, server string) (has bool, err error) {
	if has, err = l.dao.DelMapping(c, mid, key, server); err != nil {
		log.Errorf("l.dao.DelMapping(%d,%s) error(%v)", mid, key, server)
		return
	}
	log.Infof("conn disconnected key:%s server:%s mid:%d", key, server, mid)
	return
}

// Heartbeat heartbeat a conn.
func (l *Logic) Heartbeat(c context.Context, mid int64, key, server string) (err error) {
	fmt.Println("心跳")
	has, err := l.dao.ExpireMapping(c, mid, key)
	if err != nil {
		log.Errorf("l.dao.ExpireMapping(%d,%s,%s) error(%v)", mid, key, server, err)
		return
	}
	if !has {
		if err = l.dao.AddMapping(c, mid, key, server); err != nil {
			log.Errorf("l.dao.AddMapping(%d,%s,%s) error(%v)", mid, key, server, err)
			return
		}
	}
	log.Infof("conn heartbeat key:%s server:%s mid:%d", key, server, mid)
	return
}

// RenewOnline renew a server online.
func (l *Logic) RenewOnline(c context.Context, server string, roomCount map[string]int32) (map[string]int32, error) {
	online := &model.Online{
		Server:    server,
		RoomCount: roomCount,
		Updated:   time.Now().Unix(),
	}
	if err := l.dao.AddServerOnline(context.Background(), server, online); err != nil {
		return nil, err
	}
	return l.roomCount, nil
}

// Receive receive a message.
func (l *Logic) Receive(c context.Context, mid int64, proto *grpc.Proto) (err error) {
	//fmt.Println(mid, proto)
	//xlz========================
	//cinfo,err:=l.dao.GetConInfo(int(mid))
	//if  err!=nil{
	//	return
	//}
	//room:=cinfo.Room
	//_,err= cmd.Cmd(proto.Body,int(mid),cinfo.Token,room.Id)
	//if err!=nil {
	//	msg:=coder.Encode("chat",[]message.Message{{Msg:"维护中请稍后尝试",Name:"网络提示"}})
	//	err = l.PushMids(context.TODO(), int32(cinfo.Operation), []int64{mid}, []byte(msg))
	//}
	//xlz=========================
	info,err:=l.dao.GetConInfo(int(mid))
	if err!=nil {
		return
	}
	var content msg.Content
	switch proto.Op {
	case operation.OpChat:
		var chatmsg chat.Chat
		chatmsg,err= chat.InitFromByte(proto.Body)
		if err==nil {
			content=&chatmsg
		}
	case operation.OpCmd:
		var cmdmsg cmd.Cmd
		cmdmsg,err= cmd.InitFromByte(proto.Body)
		if err==nil {
			content=&cmdmsg
		}
	default:
		err=l.PushMids(c, operation.OpChat, []int64{mid}, []byte(time.Now().String()))
	}
	if content==nil {
		cmdmsg:=cmd.Init().SetContent(toast.Toast{Title:"提示",Msg:fmt.Sprintf("格式异常:%v",err)})
		err=l.PushMids(c, operation.OpCmd, []int64{mid}, cmdmsg.Bytes())
		return
	}
	req:=&ghandler.HandlerReq{
		Timestamp:time.Now().Unix(),
		Operation:int64(proto.Op),
		Uid:cast.ToString(mid),
		Token:info.Token,
		Group:"",
		Content:content.String(),
	}
	replay,_:=handler.Handler("dpg.rp",req)
	if replay.Code==-1 {
		return nil
	}
	log.Infof("receive mid:%d message:%+v", mid, proto)
	return
}
