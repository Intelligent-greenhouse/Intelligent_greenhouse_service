package infra

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"intelligent-greenhouse-service/conf"
	"strconv"
)

type Mqtt struct {
	Mq *mqtt.Client
}

func NewMqttClient(c *conf.Trigger) (*Mqtt, error) {
	opts := mqtt.NewClientOptions().AddBroker(c.Mqtt.Host + ":" + strconv.Itoa(int(c.Mqtt.Port)))
	mqttClient := mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &Mqtt{Mq: &mqttClient}, nil
}
