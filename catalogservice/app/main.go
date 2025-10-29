package main

import (
	"catalogservice/configs"
	"catalogservice/modules/logs"
	"catalogservice/modules/servers"
	"log"
	"os"

	databases "catalogservice/pkg/databases/postgresql"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func main() {
	// Load dotenv config
	if err := godotenv.Load("../config.env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	// Fiber configs
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	// Database Configs
	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	// Kafka Config
	cfg.Sarama.Host = []string{os.Getenv("KAFKA_HOST")}
	cfg.Sarama.Group = os.Getenv("KAFKA_GROUP")

	//Redis Config
	cfg.Redis.Addr = os.Getenv("REDIS_ADDR")

	// New Database
	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//defer db.Close()

	// Kafka Producer Config
	producer, err := sarama.NewSyncProducer(cfg.Sarama.Host, nil)
	if err != nil {
		logs.Error(err)
	}
	defer producer.Close()

	// Kafka Consumer Config
	consumer, err := sarama.NewConsumerGroup(cfg.Sarama.Host, cfg.Sarama.Group, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	s := servers.NewServer(cfg, db, producer, consumer)
	s.Start()
}
