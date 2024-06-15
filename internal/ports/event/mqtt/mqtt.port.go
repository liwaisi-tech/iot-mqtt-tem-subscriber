package ports

type MQTTPorts interface {
	RunConsumer(topic string)
}
