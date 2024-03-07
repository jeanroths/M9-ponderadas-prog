package main

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"database/sql"
)

var db* sql.DB

var messageSubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
	result := strings.Split(string(msg.Payload()), " - ")

	nh3, _ := strconv.Atoi(result[1])
	co, _ := strconv.Atoi(result[2])
	no2, _ := strconv.Atoi(result[3])

	data := Sensor{NH3_ppm: nh3, CO_ppm: co, NO2_ppm: no2}
	insertData(data, db)
	
}


func Sub(dbPointer *sql.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	db = dbPointer 
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("subscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	opts.SetDefaultPublishHandler(messageSubHandler)

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("/pond5", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
	fmt.Println("Subscriber está rodando. Pressione CTRL+C para sair.")
	select {}
	client.Disconnect(250)
}
