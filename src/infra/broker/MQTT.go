package broker

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type RegisterHumidityCtrlnput struct {
	Humidity int64  `json:"humidity,omitempty"`
	SensorID string `json:"sensor_id,omitempty"`
}

type RegisterTemperatureCtrlnput struct {
	Temperature float64 `json:"temperature,omitempty"`
	Humidity float64 `json:"humidity,omitempty"`
	SensorID    string  `json:"sensor_id,omitempty"`
}

type MQTTBroker struct {
	topicsToConsumer *TopicsToConsumer
	client           MQTT.Client
}

func NewMQTTBroker(clientId string) *MQTTBroker {
	broker := &MQTTBroker{
		topicsToConsumer: &TopicsToConsumer{},
	}

	broker.initClient(clientId)
	return broker
}

func (mqb *MQTTBroker) onConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to broker")

	for _, topic := range mqb.topicsToConsumer.GetAll() {
		topic.SetMQTTClient(client)

		if token := client.Subscribe(topic.GetTopic(), 0, nil); token.Wait() && token.Error() != nil {
			panic(token.Error())
		}

		fmt.Println("Subscribed to topic:", topic.GetTopic())
	}
}

func (mqb *MQTTBroker) onMessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())

	for _, topic := range mqb.topicsToConsumer.GetAll() {
		go topic.Execute(msg.Topic(), msg.Payload())
	}
}
func (mqb *MQTTBroker) SetSubscribeTopics(topicsToConsumer *TopicsToConsumer) {
	mqb.topicsToConsumer = topicsToConsumer
}

func (mqb *MQTTBroker) Connect() {
	if token := mqb.client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Connecting to broker")
}

func (mqb *MQTTBroker) GetClient() MQTT.Client {
	return mqb.client
}

func (mqb *MQTTBroker) initClient(clientId string) {
	basePath, _ := os.Getwd()
	keysPath := basePath + "/keys/"

	brokerURL := os.Getenv("BROKER_URL")
	certFile := keysPath + os.Getenv("CERT_FILE")
	keyFile := keysPath + os.Getenv("PRIVATE_FILE")
	caFile := keysPath + os.Getenv("CA_FILE")

	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientId)
	opts.SetTLSConfig(mqb.newTLSConfig(caFile, certFile, keyFile))
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(mqb.onConnectHandler)
	opts.SetDefaultPublishHandler(mqb.onMessageHandler)
	client := MQTT.NewClient(opts)
	mqb.client = client

}

func (mqb *MQTTBroker) Disconnect() {
	mqb.client.Disconnect(250)
}

func (mqb *MQTTBroker) StartConsumers() {
	mqb.Connect()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	panic("Finish StartConsumers")
}

func (mqb *MQTTBroker) newTLSConfig(caFile, certFile, keyFile string) *tls.Config {
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal("Error reading CA certificate file:", err)
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal("Error loading certificate file:", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}
}
