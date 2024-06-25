package handlers

import (
	"context"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type HandlerMessage func(ctx context.Context, eventMessage []byte) (err error)

type SubscriberHandler struct {
	ctx            context.Context
	handlerMessage HandlerMessage
	choke          chan mqtt.Message
}

func New(
	ctx context.Context,
	handlerMessage HandlerMessage,
) *SubscriberHandler {
	return &SubscriberHandler{
		ctx:            ctx,
		handlerMessage: handlerMessage,
	}
}

func (sh *SubscriberHandler) RunConsumer(topic string) {
	qos := byte(0)
	sh.choke = make(chan mqtt.Message)
	clientOptions := getClientOptions()
	clientOptions.SetDefaultPublishHandler(
		func(client mqtt.Client, msg mqtt.Message) {
			sh.choke <- msg
		},
	)
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
		incoming := <-sh.choke
		log.Info().Str("topic", "").Msgf("Received message: %v", string(incoming.Payload()))
		err := sh.handlerMessage(sh.ctx, incoming.Payload())
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
	clientOptions.SetClientID(fmt.Sprintf("%s-%s", os.Getenv("MQTT_CLIENT_ID"), uuid.New().String()))
	clientOptions.SetUsername(os.Getenv("MQTT_USERNAME"))
	clientOptions.SetPassword(os.Getenv("MQTT_PASSWORD"))
	clientOptions.SetCleanSession(true)
	return
}
