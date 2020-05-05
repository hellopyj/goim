package adapter

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
)
// natsDao dao for nats
func init()  {
	RegAdapter("nats",func()PushMsg{
		return NewNats(conf.Conf)
	})
}
type natsDao struct {
	c    *conf.Config
	push *nats.Conn
}
// New new a dao and return.
func NewNats(c *conf.Config) (dao *natsDao) {
	conn, err := nats.Connect(c.Nats.Brokers)
	if err != nil {
		return
	}
	dao= &natsDao{
		c:    c,
		push: conn,
	}
	return
}

// PublishMessage  push message to nats
func (d *natsDao) PublishMessage( pbmsg *pb.PushMsg) (err error) {
	_,value,err:=FormatMsg(pbmsg)
	if err!=nil {
		return
	}
	if d.push == nil||!d.push.IsConnected() {
		return errors.New("nats error")
	}
	msg := &nats.Msg{Subject: d.c.Nats.Topic, Reply: d.c.Nats.AckInbox, Data: value}
	if(d.push.IsConnected()){
		return d.push.PublishMsg(msg)
	}else {
		return errors.New("")
	}

}

// Close close the resource.
func (d *natsDao) Close() error {
	if d.push != nil {
		d.push.Close()
	}
	return nil
}

