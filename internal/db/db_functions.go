package db

import (
	"encoding/json"
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
		return nil, err
	}

	data, _, errI := client.From("users").Insert(UserDB{
		Username: "admin",
		Password: passwordHash,
		Role:     "admin",
	}, true, "id", "representation", "exact").Execute()

	return data, errI
}

func Login(client *supabase.Client, username string, password string) (bool, error) {
	data, _, err := client.From("users").Select("password", "exact", false).
		Eq("username", username).Execute()
	if err != nil {
		return false, err
	}

	type passwordStruct struct {
		Password string `json:"password"`
	}

	var passwordData []passwordStruct

	err = json.Unmarshal(data, &passwordData)

	if len(passwordData) > 0 {
		compare := hash.CheckPassword(password, passwordData[0].Password)

		return compare, err
	}

	return false, err
}
