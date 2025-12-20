package main

import (
	"fmt"
	"net/http"

	"github.com/JhowAranha/Rifa-Backend-Go/internal/db"
	"github.com/supabase-community/supabase-go"
)

func main() {
	client, err := db.CreateNewConnection()
	if err != nil {
		fmt.Print("Erro ao criar nova conexão com o banco de dados: ", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /toggle", handleToggle(client))
	mux.HandleFunc("GET /nums", handleGetNums(client))
	mux.HandleFunc("GET /nums/{id}", handleGetNumByID(client))

	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", mux)
}

func handleToggle(client *supabase.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, db.ToggleNumByID(client, "1"))
	}
}

func handleGetNums(client *supabase.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		nums, err := db.GetNums(client)
		if err != nil {
			http.Error(w, "Erro ao buscar números: ", http.StatusBadRequest)
		}

		w.Write(nums)
	}
}

func handleGetNumByID(client *supabase.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		num, err := db.GetNumByID(client, r.PathValue("id"))
		if err != nil {
			http.Error(w, "Erro ao buscar número por id: ", http.StatusBadRequest)
		}

		w.Write(num)
	}
}
