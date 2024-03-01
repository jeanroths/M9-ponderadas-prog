# Integração simulador com HiveMQ

Este é um conjunto de testes para garantir o funcionamento correto do simulador de sensor MiCS-6814 com MQTT. Nessa atividade, houve integração entre o simulador desenvolvido nas duas primeiras atividades ponderadas e um cluster configurado no HiveMQ. Para tal, o simulador é capaz de se comunicar utilizando autenticação em camada de transporte (TLS).

## Requisitos

1. Golang (Go)

2. Navegue até o diretório do projeto:

    ```bash
    cd ponderada4
    ```

## Uso

### Instalando Dependências - Go Mod

Acesse o diretorio que contem as dependências necessárias para cada função: 

Para o publisher e subscriber:
```
/src/ponderada2
```

Acione as dependências para cada uma das pastas, com: 
```
go mod tidy
```
Após isso, execute o publisher com:

```
go run publisher.go
```

E o subscriber no diretorio `sub`com:

```
go run subscriber.go
```

### Testes

No diretório que você executou o `publisher.go` execute o seguinte comando:

```
go test -v -cover
```

Resultado esperado:

```
=== RUN   TestConection
    simulation_test.go:21: Conection with broker MQTT successful
--- PASS: TestConection (0.00s)
=== RUN   TestValidationData
    simulation_test.go:38: Data validation successful.
--- PASS: TestValidationData (0.00s)
=== RUN   TestPublishMessages
    simulation_test.go:94: Message received successfull.
--- PASS: TestPublishMessages (0.00s)
PASS
coverage: 10.5% of statements
ok  	paho-go	0.004s
```
## Teste de Conexão com o Broker

### Propósito

Recebimento - garante que os dados enviados pelo simulador são recebidos pelo broker.

### Entrada:
- Nenhuma.

### Saída Esperada:
- Conection with broker MQTT successful.

## Explicação Publisher

### Definição de Handlers

```
var messagePubHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
}

var connectHandler MQTT.OnConnectHandler = func(client MQTT.Client) {
    fmt.Println("Connected")
}

var connectLostHandler MQTT.ConnectionLostHandler = func(client MQTT.Client, err error) {
    fmt.Printf("Connection lost: %v", err)
}
```
Esses trechos definem três handlers para eventos MQTT:

- messagePubHandler: É chamado quando uma mensagem é publicada no tópico em que o cliente está inscrito.
- connectHandler: É chamado quando o cliente MQTT se conecta com sucesso ao broker MQTT.
- connectLostHandler: É chamado quando a conexão MQTT com o broker é perdida.

### Função MsgSender

```
func MsgSender() map[string]int {
    data := map[string]int{
        "NH3_ppm": rand.Intn(300),
        "CO_ppm":  rand.Intn(1000),
        "NO2_ppm": rand.Intn(10),
    }
    return data
}
```
Essa função MsgSender é responsável por gerar dados simulados do sensor. Ela retorna um mapa contendo os valores dos gases NH3, CO e NO2 em partes por milhão (ppm), sendo os valores gerados aleatoriamente dentro de um intervalo específico.

### Função Client

```
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
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnect = connectHandler
    opts.OnConnectionLost = connectLostHandler

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

        token := client.Publish("/pond4", 1, false, msg) //QoS 1
        token.Wait()

        fmt.Println("Published:", msg)
        time.Sleep(2 * time.Second)
    }
    client.Disconnect(250)
}
```
Essa função Client é responsável por configurar e iniciar o cliente MQTT. Ela carrega as configurações do ambiente de um arquivo .env, define as opções do cliente MQTT, incluindo o broker, porta, identificação, credenciais, handlers de eventos e outros. Em seguida, ela cria um novo cliente MQTT, conecta-se ao broker e, em um loop infinito, gera dados simulados do sensor, converte-os em formato JSON e os publica no tópico /pond4 com um intervalo de 2 segundos entre as publicações. Finalmente, desconecta-se do broker quando termina.


## Vídeo
