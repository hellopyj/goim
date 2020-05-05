package adapter

import (
	"context"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/Terry-Mao/goim/internal/job/conf"
	"github.com/gogo/protobuf/proto"
	log "github.com/golang/glog"
	"github.com/nats-io/stan.go"
	"time"
)
func init()  {
	RegAdapter("natsst",func()JobConsumer{
		return NewNatsSt(conf.Conf)
	})
}
type natsStConsumer struct {
	conf *conf.Config
	consumer stan.Conn
	job consume
}

// Consume messages, watch signals
func (c *natsStConsumer) Consume(j consume) {
	c.job=j
	ctx := context.Background()
	// process push message
	pushMsg := new(pb.PushMsg)

	if _, err := c.consumer.Subscribe(c.conf.Nats.Topic, func(msg *stan.Msg) {
		log.Info("------------> ", string(msg.Data))

		if err := proto.Unmarshal(msg.Data, pushMsg); err != nil {
			log.Errorf("proto.Unmarshal(%v) error(%v)", msg, err)
			return
		}
		if err := j.Push(context.Background(), pushMsg); err != nil {
			log.Errorf("push(%v) error(%v)", pushMsg, err)
		}
		log.Infof("consume: %d  %s \t%+v", msg.Data, pushMsg)

	},stan.DurableName(c.conf.Nats.Durable)); err != nil {
		return
	}

	<-ctx.Done()
	return
}
func NewNatsSt(c *conf.Config) *natsStConsumer {
	consumer:=natsStConsumer{conf:c}
	consumer.connect()
	return &consumer
}
func (d *natsStConsumer) connect() (conn stan.Conn,err error) {
	conn, err = stan.Connect(d.conf.Nats.Cluster,d.conf.Nats.Group,stan.NatsURL(d.conf.Nats.Brokers),stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
		count:=0
		for {
			if count>100 {
				return
			}
			count++
			_,e:=d.connect()
			if e==nil {
				d.Consume(d.job)
				return
			}
			time.Sleep(time.Second*1)
		}
	}))
	if err==nil {
		d.consumer=conn
	}
	return
}
func (d *natsStConsumer) Close() (err error) {
	if d.consumer!= nil {
		d.consumer.NatsConn().Close()
		err=d.consumer.Close()
	}
	return
}

