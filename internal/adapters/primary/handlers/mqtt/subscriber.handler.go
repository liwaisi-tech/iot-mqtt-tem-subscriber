package handlers

import (
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	repositories "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/repositories/climate_data"
	"github.com/rs/zerolog/log"
)

type SubscriberHandler struct {
	climateDataRepository repositories.ClimateDataRepositoryPort
}
type mqttMessage struct {
	topic   string
	message string
}

func New(
	climateDataRepository repositories.ClimateDataRepositoryPort,
) *SubscriberHandler {
	return &SubscriberHandler{
		climateDataRepository: climateDataRepository,
	}
}

func (sh *SubscriberHandler) RunConsumer(topic string) {
	qos := byte(0)
	choke := make(chan *mqttMessage)
	clientOptions := getClientOptions()
	clientOptions.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		incomingMessage := &mqttMessage{
			topic:   msg.Topic(),
			message: string(msg.Payload()),
		}
		choke <- incomingMessage
	})

	client := mqtt.NewClient(clientOptions)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal().Err(token.Error()).Msg("Error connecting to MQTT Broker")
		panic(token.Error())
	}

	if token := client.Subscribe(topic, qos, nil); token.Wait() && token.Error() != nil {
		log.Fatal().Err(token.Error()).Msg("Error subscribing to MQTT Broker")
		panic(token.Error())
	}
	log.Info().Str("topic", topic).Msg("Subscribed to MQTT Broker")
	for {
		incoming := <-choke
		log.Info().Msgf("Received message: %v", incoming.message)
	}
}

func getClientOptions() (clientOptions *mqtt.ClientOptions) {
	clientOptions = mqtt.NewClientOptions()
	clientOptions.AddBroker(os.Getenv("MQTT_BROKER"))
	clientOptions.SetClientID(os.Getenv("MQTT_CLIENT_ID"))
	clientOptions.SetUsername(os.Getenv("MQTT_USERNAME"))
	clientOptions.SetPassword(os.Getenv("MQTT_PASSWORD"))
	clientOptions.SetCleanSession(true)
	return
}
