package bean

type server struct {
	Urls []string
}
type Xserver struct {
	Handler *server
	Push *server
}
type Nats struct {
	Cluster  string
	Topic    string
	Group    string
	TopicID  string
	Brokers  string
	AckInbox string
	Durable string
}
