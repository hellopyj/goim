package adapter

import (
	"github.com/Shopify/sarama"
	"github.com/Terry-Mao/goim/internal/logic/conf"
	kafka "gopkg.in/Shopify/sarama.v1"
	pb "github.com/Terry-Mao/goim/api/logic/grpc"
)
func init()  {
	RegAdapter("kafka",func()PushMsg{
		return NewKafka(conf.Conf)
	})
}
// Dao dao.
type kafkaDao struct {
	c    *conf.Config
	push kafka.SyncProducer
}
// New new a dao and return.
func NewKafka(c *conf.Config) (dao *kafkaDao) {
	dao= &kafkaDao{c:c,push:newKafkaPub(c.Kafka)}
	return
}

// PublishMessage  push message to kafka
func (d *kafkaDao) PublishMessage(pbmsg *pb.PushMsg) (err error) {
	key,value,err:=FormatMsg(pbmsg)
	if err!=nil {
		return
	}
	m := &kafka.ProducerMessage{
		Key:   sarama.StringEncoder(key),
		Topic: d.c.Kafka.Topic,
		Value: sarama.ByteEncoder(value),
	}
	_, _, err= d.push.SendMessage(m)
	return
}

// Close close the resource.
func (d *kafkaDao) Close() error {
	return d.push.Close()
}

func newKafkaPub(c *conf.Kafka) kafka.SyncProducer {
	var err error
	kc := kafka.NewConfig()
	kc.Producer.RequiredAcks = kafka.WaitForAll // Wait for all in-sync replicas to ack the message
	kc.Producer.Retry.Max = 10                  // Retry up to 10 times to produce the message
	kc.Producer.Return.Successes = true
	pub, err := kafka.NewSyncProducer(c.Brokers, kc)
	if err != nil {
		panic(err)
	}
	return pub
}
