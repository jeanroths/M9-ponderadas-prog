package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"math/rand"
	"os"
	"time"
	"strconv"
)

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
	fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func MsgSender() map[string]int {
	data := map[string]int{
		"NH3_ppm": rand.Intn(300),
		"CO_ppm":  rand.Intn(1000),
		"NO2_ppm": rand.Intn(10),
	}
	return data
}

func Client() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for {
		data := MsgSender()

		msg := time.Now().Format(time.RFC3339) + " - " + strconv.Itoa(data["NH3_ppm"]) + " - " + strconv.Itoa(data["CO_ppm"]) + " - " + strconv.Itoa(data["NO2_ppm"])

		token := client.Publish("topic_pond7", 1, false, msg) //QoS 1
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
	client.Disconnect(250)
}


func main() {
	Client()
}

