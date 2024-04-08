## Integração entre HiveMQ, Kafka e Banco de Dados SQLite

Esta aplicação integra um simulador de dados MQTT com um banco de dados SQLite via Kafka. O simulador gera dados aleatórios de NH3, CO e NO2 e os publica em um tópico MQTT no HiveMQ. Um consumidor Kafka assina esse tópico, lê os dados e os insere no banco de dados SQLite.

### Funcionamento dos Códigos

#### Publisher (publisher.go)

O código `publisher.go` atua como um publisher MQTT que gera dados aleatórios de NH3, CO e NO2 e os publica em um tópico MQTT no HiveMQ. Ele utiliza a biblioteca Eclipse Paho MQTT para se conectar ao broker MQTT e publicar mensagens contendo os dados gerados. O processo de geração e publicação de dados ocorre em um loop infinito, com um intervalo de 2 segundos entre cada publicação.

Para executar o publisher, siga estas etapas:

1. Configure as variáveis de ambiente `BROKER_ADDR`, `HIVE_USER` e `HIVE_PSWD` no arquivo `.env` com as informações necessárias para se conectar ao broker MQTT.
2. Execute o comando `go run main.go` para iniciar o publisher.

#### Consumer (consumer.go)

O código `consumer.go` atua como um consumidor Kafka que se inscreve em um tópico Kafka, lê mensagens contendo os dados gerados pelo publisher e os insere em um banco de dados SQLite. Ele utiliza a biblioteca confluent-kafka-go para se conectar ao cluster Kafka e consumir mensagens do tópico especificado. Ao receber uma mensagem, o consumidor Kafka a converte em um formato adequado e a insere no banco de dados SQLite.

Para executar o consumidor, siga estas etapas:

1. Certifique-se de ter um broker Kafka em execução e configurado corretamente.
2. Execute o comando `go run consumer.go` para iniciar o consumidor.

### Execução

1. Certifique-se de que o broker MQTT (HiveMQ) e o broker Kafka estejam configurados e em execução.
2. Configure as variáveis de ambiente no arquivo `.env` com as informações necessárias para se conectar ao broker MQTT.
3. Execute o publisher MQTT com o comando `go run main.go`.
4. Execute o consumidor Kafka com o comando `go run consumer.go`.

Após a execução bem-sucedida do publisher e do consumidor, os dados gerados pelo publisher serão publicados no tópico MQTT e consumidos pelo consumidor Kafka, que os inserirá no banco de dados SQLite.

Para visualizar os dados no banco de dados SQLite, você pode usar ferramentas como o SQLite CLI ou um cliente de banco de dados SQLite GUI.

### Considerações

Certifique-se de ter todas as dependências instaladas e configuradas corretamente. Além disso, verifique se as configurações de segurança, como credenciais de acesso, estão corretas para garantir uma conexão bem-sucedida aos brokers MQTT e Kafka.

# Demonstração

https://github.com/jeanroths/M9-ponderadas-prog/assets/99195775/5b24d3c9-00de-4062-a98c-8ce621ae8bb8

