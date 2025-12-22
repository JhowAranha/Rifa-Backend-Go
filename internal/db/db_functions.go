package db

import (
	"fmt"

	"github.com/JhowAranha/Rifa-Backend-Go/internal/hash"
	"github.com/supabase-community/supabase-go"
)

type UserDB struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func TestDBConnection() {
	_, err := CreateNewConnection()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Teste de conexão bem-sucedido!")
	}
}

func GetNums(client *supabase.Client) ([]uint8, error) {
	data, _, err := client.From("numero").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetNumByID(client *supabase.Client, id string) ([]uint8, error) {
	data, _, err := client.From("numero").Select("*", "exact", false).Eq("id", id).Execute()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ToggleNumByID(client *supabase.Client, id int) string {
	params := map[string]interface{}{
		"num_id": id,
	}

	res := client.Rpc("toggle_num_checked", "", params)
	return res
}

func AddNewUser(client *supabase.Client) ([]uint8, error) {
	passwordHash, err := hash.CreateNewHash("admin123")
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	data, _, errI := client.From("users").Insert(UserDB{
		Username: "admin",
		Password: passwordHash,
		Role:     "admin",
	}, true, "id", "representation", "exact").Execute()

	if errI != nil {
		fmt.Println("Erro ao inserir usuário no banco de dados: ", err)
		return nil, nil
	}

	return data, nil
}
