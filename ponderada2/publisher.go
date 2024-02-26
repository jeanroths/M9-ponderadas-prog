package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"math/rand"
	"time"
)

func MsgSender() map[string]int {
	data := map[string]int{
		"NH3_ppm": rand.Intn(300),
		"CO_ppm":  rand.Intn(1000),
		"NO2_ppm": rand.Intn(10),
	}
	return data
}

func Client() {

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("publisher")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	for {
		data := MsgSender()
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error converting data to JSON", err)
			return
		}

		msg := string(jsonData) + time.Now().Format(time.RFC3339)

		token := client.Publish("/pond2", 1, false, msg) //QoS 1
		token.Wait()

		fmt.Println("Published:", msg)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	Client()
}

// for {
// 	data := map[string]interface{}{
// 		"NH3_ppm":  rand.Intn(300),
// 		"CO_ppm":   rand.Intn(1000),
// 		"NO2_ppm":  rand.Intn(10),
// 	}

// 	jsonData, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Println("Error converting data to JSON", err)
// 		return
// 	}

// 	msg := string(jsonData) + time.Now().Format(time.RFC3339)

// 	token := client.Publish("/pond2", 0, false, msg)
// 	token.Wait()

// 	fmt.Println("Published:", msg)
// 	time.Sleep(2 * time.Second)
// }

//}
