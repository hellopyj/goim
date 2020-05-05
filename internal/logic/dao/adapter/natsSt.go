package adapter

import (
	"errors"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	"github.com/nats-io/stan.go"
	"time"
)
// natsDao dao for nats
func init()  {
	RegAdapter("natsst",func()PushMsg{
		return NewNatsSt(conf.Conf)
	})
}
type natsStDao struct {
	c    *conf.Config
	push stan.Conn
}
// New new a dao and return.
func NewNatsSt(c *conf.Config) (dao *natsStDao) {
	dao= &natsStDao{c:c}
	_,err:=dao.connect()
	if err!=nil {
		panic(err)
	}
	return
}

// PublishMessage  push message to nats
func (d *natsStDao) PublishMessage( pbmsg *pb.PushMsg) (err error) {
	_,value,err:=FormatMsg(pbmsg)
	if err!=nil {
		return
	}
	if d.push == nil{
		return errors.New("nats error")
	}
	//msg := &stan.Msg{d.c.Nats.Topic,value}
	return d.push.Publish(d.c.Nats.Topic,value)
}
func (d *natsStDao) connect() (conn stan.Conn,err error) {
	conn, err = stan.Connect(d.c.Nats.Cluster,d.c.Nats.Group,stan.NatsURL(d.c.Nats.Brokers),stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
		count:=0
		for {
			if count>100 {
				return
			}
			count++
			_,e:=d.connect()
			if e==nil {
				return
			}
			time.Sleep(time.Second*1)
		}
	}),)
	if err==nil {
		d.push=conn
	}
	return
}

// Close close the resource.
func (d *natsStDao) Close() (err error) {
	if d.push != nil {
		d.push.NatsConn().Close()
		err=d.push.Close()
	}
	return
}
