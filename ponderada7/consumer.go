package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var ConsumerPointer *kafka.Consumer
var db *sql.DB

type Data struct {
	NH3_ppm  int
	CO_ppm   int
	NO2_ppm  int
}

func ReadConfig() kafka.ConfigMap {
    // reads the client configuration from client.properties
    // and returns it as a key-value map
    m := make(map[string]kafka.ConfigValue)

    file, err := os.Open("client.properties")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to open file: %s", err)
        os.Exit(1)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if !strings.HasPrefix(line, "#") && len(line) != 0 {
            kv := strings.Split(line, "=")
            parameter := strings.TrimSpace(kv[0])
            value := strings.TrimSpace(kv[1])
            m[parameter] = value
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Failed to read file: %s", err)
        os.Exit(1)
    }

    return m
}

func Consumer(db *sql.DB) *kafka.Consumer {
	if ConsumerPointer == nil {
		GenerateConsumer(db)
		return ConsumerPointer
	}
	return ConsumerPointer
}

func GenerateConsumer(db *sql.DB) {
	// Configurações do consumidor
	conf := ReadConfig()
	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		panic(err)
	}

	ConsumerPointer = consumer
}

func Subscribe(consumer *kafka.Consumer, db *sql.DB, topic string) {
    // Se o consumidor ou o banco de dados forem nulos, não há nada a fazer
    if consumer == nil || db == nil {
        fmt.Println("Consumer or database is nil")
        return
    }
    fmt.Printf("aqui: %v\n", consumer)
    // Assinar tópico
    err := consumer.SubscribeTopics([]string{topic}, nil)
    if err != nil {
        panic(err)
    }

    // Consumir mensagens
    for {
        msg, err := consumer.ReadMessage(-1)
        if err == nil {
            fmt.Printf("Received message: %s\n", string(msg.Value))
            result := strings.Split(string(msg.Value), " - ")
            nh3, _ := strconv.Atoi(result[1])
            co, _ := strconv.Atoi(result[2])
            no2, _ := strconv.Atoi(result[3])
            data := &Data{NH3_ppm: nh3, CO_ppm: co, NO2_ppm: no2}
            insertData(db, *data)
        } else {
            fmt.Printf("Consumer error: %v (%v)\n", err, msg)
            break
        }
    }
}

func insertData(db *sql.DB, data Data) {
	// Prepara a consulta SQL para inserir os dados no banco de dados
	query := "INSERT INTO sensor_data (NH3_ppm, CO_ppm, NO2_ppm) VALUES (?, ?, ?)"
	_, err := db.Exec(query, data.NH3_ppm, data.CO_ppm, data.NO2_ppm)
	if err != nil {
		fmt.Println("Error inserting data:", err)
	}
}

func main() {
    // Abre o banco de dados SQLite
    database, err := sql.Open("sqlite3", "./db.db")
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer database.Close()

    // Cria a tabela se ela não existir
    _, err = database.Exec(`CREATE TABLE IF NOT EXISTS sensor_data (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        NH3_ppm INTEGER,
        CO_ppm INTEGER,
        NO2_ppm INTEGER
    )`)
    if err != nil {
        fmt.Println("Error creating table:", err)
        return
    }

    // Inicia o consumidor Kafka
    consumer := Consumer(database)
    Subscribe(consumer, database, "topic_pond7")
}

