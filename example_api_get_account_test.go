package yta_test

import (
	"encoding/json"
	"fmt"

	yta "github.com/ystar-foundation/yta-go"
)

func ExampleAPI_GetAccount() {
	api := yta.New(getAPIURL())

	account := yta.AccountName("yta.rex")
	info, err := api.GetAccount(account)
	if err != nil {
		if err == yta.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
