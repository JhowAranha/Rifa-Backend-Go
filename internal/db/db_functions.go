package db

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

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
