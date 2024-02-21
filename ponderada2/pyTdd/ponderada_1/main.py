import paho.mqtt.client as mqtt
import time
from faker import Faker
import random

# Configuração do cliente MQTT
client = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)

# Conectando ao broker Mosquitto
client.connect("localhost", 1891, 60)

# Função para gerar dados do sensor MiCS-6814 simulados
def generate_sensor_data():
    fake = Faker()
    sensor_outputs = {
        "CO_ppm": fake.pyfloat(min_value=1, max_value=1000, right_digits=2),
        "NO2_ppm": fake.pyfloat(min_value=0.05, max_value=10, right_digits=2),
        "NH3_ppm": fake.pyfloat(min_value=1, max_value=300, right_digits=2),
    }
    return sensor_outputs

# Loop para enviar mensagens continuamente
try:
    while True:
        sensor_data = generate_sensor_data()
        client.publish("sensor/mics6814", str(sensor_data))
        print(f"Mensagem enviada: {sensor_data}")
        time.sleep(2)  # Aguarda 2 segundos antes de enviar a próxima mensagem
except KeyboardInterrupt:
    print("Publicação encerrada")

# Desconecta do broker MQTT ao encerrar
client.disconnect()
