package adapter

import (
	"context"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
	"github.com/Terry-Mao/goim/internal/job/conf"
	"github.com/bsm/sarama-cluster"
	"github.com/gogo/protobuf/proto"
	log "github.com/golang/glog"
)
func init()  {
	RegAdapter("kafka",func()JobConsumer{
		return NewKafka(conf.Conf)
	})
}
type kafkaConsumer struct {
	consumer *cluster.Consumer
	conf *conf.Config
}
func NewKafka(c *conf.Config) *kafkaConsumer {
	return &kafkaConsumer{
		consumer: newKafkaSub(c.Kafka),
		conf: conf.Conf,
	}
}

func (c *kafkaConsumer) Close() error {
	return c.consumer.Close()
}

// Consume messages, watch signals
func (c *kafkaConsumer) Consume(j consume) {
	for {
		select {
		case err := <-c.consumer.Errors():
			log.Errorf("consumer error(%v)", err)
		case n := <-c.consumer.Notifications():
			log.Infof("consumer rebalanced(%v)", n)
		case msg, ok := <-c.consumer.Messages():
			if !ok {
				return
			}
			c.consumer.MarkOffset(msg, "")
			// process push message
			pushMsg := new(pb.PushMsg)
			if err := proto.Unmarshal(msg.Value, pushMsg); err != nil {
				log.Errorf("proto.Unmarshal(%v) error(%v)", msg, err)
				continue
			}
			if err := j.Push(context.Background(), pushMsg); err != nil {
				log.Errorf("c.push(%v) error(%v)", pushMsg, err)
			}
			log.Infof("consume: %s/%d/%d\t%s\t%+v", msg.Topic, msg.Partition, msg.Offset, msg.Key, pushMsg)
		}
	}
}
func newKafkaSub(c *conf.Kafka) *cluster.Consumer {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	consumer, err := cluster.NewConsumer(c.Brokers, c.Group, []string{c.Topic}, config)
	if err != nil {
		panic(err)
	}
	return consumer
}
