package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

func CreateNewConnection() (*supabase.Client, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %v", err)
	}

	client, error := supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_KEY"),
		&supabase.ClientOptions{},
	)
	if error != nil {
		log.Fatalf("Erro ao criar cliente Supabase: %v", error)
		return nil, error
	}

	return client, nil

}
