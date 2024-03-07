# Integração do simulador com Metabase

Esta aplicação tem como objetivo fornecer uma visão geral do processo de publicação de informações em um tópico MQTT por um publicador (publisher), leitura dessas informações por um assinante (subscriber) e inserção desses dados em um banco de dados SQLite. Além disso, abordará a integração do banco de dados com um dashboard no Metabase.

## Requisitos

- GoLang instalado no sistema
- Docker instalado no sistema
- MQTT Broker configurado (HiveMQ ou outro)
- Metabase Docker Container

## Configuração

1. **Configuração do Ambiente MQTT**

- Certifique-se de que o broker MQTT esteja configurado corretamente e acessível.
- Defina as variáveis de ambiente BROKER_ADDR, HIVE_USER, e HIVE_PSWD em um arquivo .env contendo as informações necessárias para se conectar ao broker MQTT.

2. **Configuração do Metabase**

- Execute o Metabase em um contêiner Docker usando o seguinte comando:

```
docker run -d -p 3000:3000 -v $(pwd)/db.db:/db.db --name metabase metabase/metabase
```
## Execução 

1. **Compilação e Execução do Código Go**

Para iniciar o publisher, subscriber e criar o banco de dados com a função para inserir dados, basta apenas rodar o seguinte comando:
```
go run *.go
```

2. **Visualização de dados no Metabase**

- Acesse o Metabase no navegador utilizando o endereço `localhost:3000`.
- Configure o banco de dados SQLite como uma fonte de dados no Metabase.
- Crie consultas e dashboards para visualizar os dados inseridos pelo subscriber no banco de dados SQLite.

## Explicação

- O `publisher` gera dados aleatórios representando níveis de NH3, CO e NO2, e os publica em um tópico MQTT.
- O `subscriber` se inscreve nesse tópico e, ao receber os dados, os converte em um formato adequado e os insere no banco de dados SQLite.
- O Metabase está configurado para se conectar ao banco de dados SQLite e fornecer uma interface para visualização dos dados.

## Vídeo
