package broker

import MQTT "github.com/eclipse/paho.mqtt.golang"

type Command interface {
	Execute(currentTopic string, message []byte)
	GetTopic() string
	SetMQTTClient(client MQTT.Client)
}

type TopicsToConsumer struct {
	topicsCommand []Command
}

func NewTopicsToConsumer() *TopicsToConsumer {
	return &TopicsToConsumer{}
}

func (ttc *TopicsToConsumer) Add(topic Command) *TopicsToConsumer {
	ttc.topicsCommand = append(ttc.topicsCommand, topic)
	return ttc
}

func (ttc *TopicsToConsumer) GetAll() []Command {
	return ttc.topicsCommand
}
