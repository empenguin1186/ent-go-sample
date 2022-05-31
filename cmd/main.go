package main

import (
	"context"
	"ent-go-sample/ent"
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

	CreateUser(context.Background(), client)
}

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
