package handlers

import (
	"context"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	usecaseports "github.com/liwaisi-tech/iot-mqtt-tem-subscriber/internal/ports/usecases"
	"github.com/rs/zerolog/log"
)

type SubscriberHandler struct {
	ctx                    context.Context
	saveClimateDataUseCase usecaseports.SaveClimateDataUseCasePort
}
type mqttMessage struct {
	topic   string
	message []byte
}

func New(
	ctx context.Context,
	saveClimateDataUseCase usecaseports.SaveClimateDataUseCasePort,
) *SubscriberHandler {
	return &SubscriberHandler{
		saveClimateDataUseCase: saveClimateDataUseCase,
	}
}

func (sh *SubscriberHandler) RunConsumer(topic string) {
	qos := byte(0)
	choke := make(chan *mqttMessage)
	clientOptions := getClientOptions()
	clientOptions.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		incomingMessage := &mqttMessage{
			topic:   msg.Topic(),
			message: msg.Payload(),
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
		log.Info().Msgf("Received message: %v", string(incoming.message))
		err := sh.saveClimateDataUseCase.Execute(sh.ctx, incoming.message)
		if err != nil {
			log.Error().Err(err).Msg("Failed processing MQTT message")
			continue
		}
		log.Info().Msg("MQTT message processed successfully")
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
