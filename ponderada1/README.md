# Simulador de Sensor MiCS-6814 com MQTT

Este é um simulador de sensor MiCS-6814 que gera dados fictícios para simular a leitura de gases como CO, NO2 e NH3, e os envia via MQTT.

## Requisitos

- Python 3.x
- Mosquitto MQTT Broker

## Instalação

1. Certifique-se de ter o Python instalado. Você pode baixá-lo em [python.org](https://www.python.org/).

2. Instale o Mosquitto MQTT Broker. Você pode seguir as instruções de instalação no site oficial: [mosquitto.org](https://mosquitto.org/download/).


3. Navegue até o diretório do projeto:

    ```bash
    cd ponderada1
    ```

5. Instale as dependências do Python:

    ```bash
    pip install -r requirements
    ```

## Uso

1. Inicie o broker MQTT Mosquitto (se ainda não estiver em execução). Se estiver usando Linux ou macOS, você pode iniciar o broker com:

    ```bash
    mosquitto -c mosquitto.conf
    ```

    Certifique-se de que o arquivo de configuração `mosquitto.conf` esteja apontando para o listener na porta 1891 (ou ajuste o código Python e o arquivo de configuração conforme necessário).

2. Execute o script Python `main.py` para iniciar a simulação:

    ```bash
    python3 main.py #caso esteja usando Windows trocar python3 por python
    ```

    Isso iniciará o simulador e ele começará a enviar dados simulados do sensor para o broker MQTT.

3. Para visualizar os dados que estão sendo enviados, use o `mosquitto_sub` em outro terminal:

    ```bash
    mosquitto_sub -h localhost -p 1891 -v -t "sensor/mics6814" 
    ```

    Isso mostrará as mensagens recebidas no tópico "sensor/mics6814".

4. Para interromper a execução do simulador, pressione `Ctrl+C` no terminal onde o simulador está sendo executado.

## Explicação

- `main.py`: Este script Python é responsável por gerar dados simulados do sensor MiCS-6814 e enviá-los para um broker MQTT usando a biblioteca Paho MQTT. Os dados são gerados com valores aleatórios usando a biblioteca Faker.

- `mosquitto.conf`: Este arquivo de configuração do Mosquitto MQTT Broker define um listener na porta 1891 e permite conexões anônimas para simplificar o teste do simulador.


## Vídeo