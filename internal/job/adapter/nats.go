package adapter

import (
	"context"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/Terry-Mao/goim/internal/job/conf"
	"github.com/gogo/protobuf/proto"
	log "github.com/golang/glog"
	"github.com/nats-io/nats.go"
)
func init()  {
	RegAdapter("nats",func()JobConsumer{
		return NewNats(conf.Conf)
	})
}
type natsConsumer struct {
	conf *conf.Config
	consumer *nats.Conn
}
// Consume messages, watch signals
func (c *natsConsumer) Consume(j consume) {
	ctx := context.Background()

	// process push message
	pushMsg := new(pb.PushMsg)

	if _, err := c.consumer.Subscribe(c.conf.Nats.Topic, func(msg *nats.Msg) {

		log.Info("------------> ", string(msg.Data))

		if err := proto.Unmarshal(msg.Data, pushMsg); err != nil {
			log.Errorf("proto.Unmarshal(%v) error(%v)", msg, err)
			return
		}
		if err := j.Push(context.Background(), pushMsg); err != nil {
			log.Errorf("push(%v) error(%v)", pushMsg, err)
		}
		log.Infof("consume: %d  %s \t%+v", msg.Data, pushMsg)

	}); err != nil {
		return
	}

	<-ctx.Done()
	return
}
func NewNats(c *conf.Config) *natsConsumer {

	nc, err := nats.Connect(c.Nats.Brokers)
	if err != nil {
		return nil
	}
	return &natsConsumer{
		consumer: nc,
		conf:conf.Conf,
	}
}

func (c *natsConsumer) Close() error {
	c.consumer.Close()
	return nil
}

