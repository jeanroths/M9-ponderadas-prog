# Testes de Simulação de comunicação MQTT e dados do sensor MiCS-6814

Este é um conjunto de testes para garantir o funcionamento correto do simulador de sensor MiCS-6814 com MQTT.

## Requisitos

Golang (Go)
Mosquitto MQTT Broker

## Instalação

1. Instale o Mosquitto MQTT Broker. Você pode seguir as instruções de instalação no site oficial: [mosquitto.org](https://mosquitto.org/download/).


2. Navegue até o diretório do projeto:

    ```bash
    cd ponderada2
    ```

## Uso

1. Inicie o broker MQTT Mosquitto (se ainda não estiver em execução). Se estiver usando Linux ou macOS, você pode iniciar o broker com:

    ```bash
    mosquitto -c mosquitto.conf
    ```

    Certifique-se de que o arquivo de configuração `mosquitto.conf` esteja apontando para o listener na porta 1891 (ou ajuste o código e o arquivo de configuração conforme necessário).

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

## Teste de Validação dos dados

### Propósito

Validação dos dados - garante que os dados enviados pelo simulador chegam sem alterações.

### Entrada:
Nenhuma

### Saída:

Mapa com os campos:

 - CO: Concentração de CO.
 - NH3: Concentração de NH3.
 - NO2: Concetração de NO2.

## Teste de Taxa de Disparo de Mensagens

### Propósito

Confirmação da taxa de disparo - garante que o simulador atende às especificações de taxa de disparo de mensagens dentro de uma margem de erro razoável.

### Entrada:

JSON com dados dos sensores gerados aleatoriamente

### Saída:

Mensagem JSON recebida com sucesso, contendo os mesmos campos e valores da mensagem publicada

## Vídeo

https://github.com/jeanroths/M9-ponderadas-prog/assets/99195775/49b4902f-59ed-4019-b693-df485ac5f5ae




