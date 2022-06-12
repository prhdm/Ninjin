package mqtt

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	/*"strconv"ذ
	"strings"*/
	"sync"
	log "github.com/sirupsen/logrus"
)

// MQTT credentials(you may have username and password too)
const mqttServer = "79.175.177.48:1883"
const mqttClientID = "IAmUsagi"
const userName = "usagi"
const password = "usagi1234"

// MQTT topics(channels) that we work with.
const tempTopic = "fleet/v1/#"



var wg = sync.WaitGroup{}

func Start() {
	wg.Add(1)

	greeter()

	c := createClient()

	handler := NewLogHandler()
	log.Info("new log handler created")
	if token := c.Subscribe(tempTopic, 0, func(client MQTT.Client, message MQTT.Message) {
		handler.Handle(string(message.Payload()), message.Topic())
	}); token.Wait() && token.Error() != nil {
		log.Info(token.Error())
		os.Exit(1)
	}

	wg.Wait()
}

// createClient returns a new MQTT client object.
func createClient() MQTT.Client {
	opts := MQTT.NewClientOptions().AddBroker("tcp://" + mqttServer).SetClientID(mqttClientID)
	opts.AutoReconnect = true
	opts.SetUsername(userName)
	opts.SetPassword(password)

	c := MQTT.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return c
}

// greeter prints a short introduction text to the terminal.
func greeter() {
	log.Info("=============================================")
	log.Info("* * * HELLO FROM MQTT MONITORING SERVER * * *")
	log.Info("=============================================")
}

