package main

import (
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

// Testa a conexão com o Broker
func TestConection(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	//mensagem := MsgSender(resultado, "/pond2")

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Errorf("Error in connection with broker MQTT: %v", token.Error())
	} else {
		t.Log("Conection with broker MQTT successful")
	}

}

// Testa validação dos dados
func TestValidationData(t *testing.T) {
	msg := MsgSender()

	// Verifique se todos os campos esperados estão presentes nos dados gerados
	expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
	for _, field := range expectedFields {
		if _, ok := msg[field]; !ok {
			t.Errorf("field expected: %s", field)
			return
		}
	}
	t.Log("Data validation successful.")
}

// Testa a taxa de disparo de mensagens
func TestPublishMessages(t *testing.T) {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1891")
	opts.SetClientID("test-client")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to connect MQTT broker: %v", token.Error())
	}

	// Inicia um subscriber MQTT para receber as publicações
	received := make(chan bool)
	token := client.Subscribe("/pond2", 1, func(client MQTT.Client, msg MQTT.Message) {
		// Verifique se a mensagem recebida é válida
		var data map[string]int
		if err := json.Unmarshal(msg.Payload(), &data); err != nil {
			t.Errorf("Erro ao decodificar a mensagem JSON: %v", err)
			return
		}

		// Verifica se todos os campos esperados estão presentes 
		expectedFields := []string{"NH3_ppm", "CO_ppm", "NO2_ppm"}
		for _, field := range expectedFields {
			if _, ok := data[field]; !ok {
				t.Errorf("Field %s expected but not received", field)
				return
			}
		}

		//recebido
		received <- true
	})
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to subscribe MQTT topic: %v", token.Error())
	}

	// Publica uma mensagem com dados gerados
	msg := MsgSender()
	jsonData, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Error to convert to JSON: %v", err)
	}

	data := string(jsonData)
	token = client.Publish("/pond2", 0, false, data)
	if token.Wait() && token.Error() != nil {
		t.Fatalf("Failed to post MQTT message: %v", token.Error())
	}

	// Aguarda a confirmação de recebimento
	select {
	case <-received:
		// Mensagem recebida com sucesso
		t.Log("Message received successfull.")
	case <-time.After(5 * time.Second):
		// Timeout - nenhum sinal de recebimento
		t.Fatalf("Timeout: Any message was received after 5 sec.")
	}
}
