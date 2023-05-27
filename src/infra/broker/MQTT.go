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

	"encoding/json"

	application "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/application"
	domain "github.com/Vinicius-Santos-da-Silva/greenhouse_api/src/domain"
)

type RegisterHumidityCtrlnput struct {
	Humidity int64  `json:"humidity,omitempty"`
	SensorID string `json:"sensor_id,omitempty"`
}

type RegisterTemperatureCtrlnput struct {
	Temperature float64 `json:"temperature,omitempty"`
	SensorID    string  `json:"sensor_id,omitempty"`
}

type MQTTBroker struct {
	humidityRepository    domain.SoilRepository
	temperatureRepository domain.TemperatureRepository
}

func NewMQTTBroker(humidityRepository domain.SoilRepository, temperatureRepository domain.TemperatureRepository) *MQTTBroker {
	return &MQTTBroker{
		humidityRepository,
		temperatureRepository,
	}
}

func (mqb *MQTTBroker) onConnectHandler(client MQTT.Client) {
	fmt.Println("Connected to broker")
	if token := client.Subscribe(os.Getenv("WATER_PUMP_SUBSCRIBE"), 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe(os.Getenv("TEMPERATURE_SUBSCRIBE"), 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Subscribed to topic:", os.Getenv("WATER_PUMP_SUBSCRIBE"))
	fmt.Println("Subscribed to topic:", os.Getenv("TEMPERATURE_SUBSCRIBE"))
}

func (mqb *MQTTBroker) onMessageHandler(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Received message on topic: %s\n", msg.Topic())
	fmt.Printf("Message: %s\n", msg.Payload())

	if msg.Topic() == os.Getenv("WATER_PUMP_SUBSCRIBE") {

		var inputJSON RegisterHumidityCtrlnput
		err := json.Unmarshal(msg.Payload(), &inputJSON)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Humidity: ", inputJSON.Humidity)
		_, err = mqb.humidityRepository.Create(&domain.HumidityRepositoryDTO{
			SensorID: inputJSON.SensorID,
			Value:    inputJSON.Humidity,
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		go application.NewManageWaterPump(mqb.humidityRepository).GetCommand(client)
		return
	}

	var inputJSON RegisterTemperatureCtrlnput
	err := json.Unmarshal(msg.Payload(), &inputJSON)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Temperature: ", inputJSON.Temperature)
	_, err = mqb.temperatureRepository.Create(&domain.TemperatureRepositoryDTO{
		SensorID: inputJSON.SensorID,
		Value:    inputJSON.Temperature,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go application.NewManageTemperature(mqb.temperatureRepository).GetCommand(client)
}

func (mqb *MQTTBroker) MQTTClient(clientId string) MQTT.Client {
	basePath, _ := os.Getwd()
	fmt.Println(basePath)

	keysPath := basePath + "/keys/"
	fmt.Println("\n\n", keysPath)

	brokerURL := os.Getenv("BROKER_URL")
	certFile := keysPath + os.Getenv("CERT_FILE")
	keyFile := keysPath + os.Getenv("PRIVATE_FILE")
	caFile := keysPath + os.Getenv("CA_FILE")

	opts := MQTT.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID(clientId)
	// opts.SetClientID("sdk-nodejs-v2")
	opts.SetTLSConfig(mqb.NewTLSConfig(caFile, certFile, keyFile))
	opts.SetAutoReconnect(true)
	opts.SetOnConnectHandler(mqb.onConnectHandler)
	opts.SetDefaultPublishHandler(mqb.onMessageHandler)
	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Println("Connecting to broker")

	return client
}

func (mqb *MQTTBroker) MQTTConsumer() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc

	panic("Fimmm")
}

func (mqb *MQTTBroker) NewTLSConfig(caFile, certFile, keyFile string) *tls.Config {
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
