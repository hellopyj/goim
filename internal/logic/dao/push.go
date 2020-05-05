package dao

import (
	"context"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/Terry-Mao/goim/dpg/bean/msg/chat"
	"github.com/Terry-Mao/goim/dpg/config/operation"
	log "github.com/golang/glog"
)
// PushMsg push a message to databus.
func (d *Dao) PushMsg(c context.Context, op int32, server string, keys []string, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_PUSH,
		Operation: op,
		Server:    server,
		Keys:      keys,
	}
	pushMsg.Msg,err=d.DealMsg(op,msg)
	if err!=nil {
		return
	}
	if err = d.push.PublishMessage(pushMsg); err != nil {
		log.Errorf("PushMsg.send(push pushMsg:%v) error(%v)", pushMsg, err)
	}
	return
}

// BroadcastRoomMsg push a message to databus.
func (d *Dao) BroadcastRoomMsg(c context.Context, op int32, room string, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_ROOM,
		Operation: op,
		Room:      room,
	}
	pushMsg.Msg,err=d.DealMsg(op,msg)
	if err!=nil {
		return
	}
	if err = d.push.PublishMessage(pushMsg); err != nil {
		log.Errorf("PushMsg.send(broadcast_room pushMsg:%v) error(%v)", pushMsg, err)
	}
	return
}

// BroadcastMsg push a message to databus.
func (d *Dao) BroadcastMsg(c context.Context, op, speed int32, msg []byte) (err error) {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_BROADCAST,
		Operation: op,
		Speed:     speed,
	}
	pushMsg.Msg,err=d.DealMsg(op,msg)
	if err!=nil {
		return
	}
	if err = d.push.PublishMessage(pushMsg); err != nil {
		log.Errorf("PushMsg.send(broadcast pushMsg:%v) error(%v)", pushMsg, err)
	}
	return
}
func (d *Dao) DealMsg(op int32, msg []byte) (ret []byte,err error) {
	switch op {
	case operation.OpChat:
		tmsg,e:= chat.InitFromByte(msg)
		if e!=nil {
			return nil,e
		}
		ret=tmsg.Bytes()
		if !tmsg.IsSkip() {
			if tmsg.IsUserChat() {
				d.SaveUserChat(tmsg.GetRKeys(),tmsg.ID,string(ret))
			}else if tmsg.IsGroupChat() {
				d.SaveGroupChat(tmsg.GetRKeys(),tmsg.ID,string(ret))
			}
		}

	}
	return msg,nil
}
