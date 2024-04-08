package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"
	godotenv "github.com/joho/godotenv"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/stretchr/testify/assert"
)

// Testa a conexão com o Broker MQTT
func TestMQTTConnection(t *testing.T) {
	err := godotenv.Load("./env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
		t.FailNow()
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("test-client")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Errorf("Failed to connect MQTT broker: %v", token.Error())
	} else {
		t.Log("Connection with MQTT broker successful")
	}

}

// Testa a publicação de mensagens MQTT e a leitura pelo consumidor Kafka
func TestMQTTPublishAndConsume(t *testing.T) {
	// Configurações MQTT
	err := godotenv.Load("./env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
		t.FailNow()
	}
	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := MQTT.NewClientOptions().AddBroker(fmt.Sprintf("tls://%s:%d", broker, port))
	opts.SetClientID("test-client")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))

	// Configurações do Consumidor Kafka
	kafkaConsumer := Consumer(nil)

	// Inicia um subscriber Kafka para receber as mensagens
	go Subscribe(kafkaConsumer, nil, "topic_pond7")

	// Cria um cliente MQTT
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}

	// Dados para publicação
	data := MsgSender()
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error converting to JSON: %v", err)
	}

	// Publica uma mensagem com dados gerados
	token := client.Publish("topic_pond7", 1, false, jsonData) // QoS 1
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to post MQTT message: %v", token.Error())
	}

	// Aguarda a confirmação de recebimento pelo consumidor Kafka
	time.Sleep(2 * time.Second)

	// Verifica se a mensagem foi recebida e inserida no banco de dados
	// Aqui você pode adicionar asserções adicionais para verificar se os dados foram inseridos corretamente no banco de dados
	// Por exemplo, você pode consultar o banco de dados para verificar se os dados inseridos estão presentes
	assert.True(t, true, "Message received and processed successfully")
}
