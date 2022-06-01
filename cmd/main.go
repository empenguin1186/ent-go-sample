package main

import (
	"context"
	"ent-go-sample/domain/model"
	"ent-go-sample/ent"
	"ent-go-sample/infra"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	if err != nil {
		log.Fatal("Error creating dataSource")
	}

	client, err := ent.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// オートマイグレーションツールを実行する
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	repository := infra.NewUserRepositoryImpl(client, context.Background())
	repository.Create(&model.User{Age: 10, Name: "hoge"})
}
