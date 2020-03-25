package yta_test

import (
	"encoding/json"
	"fmt"

	yta "github.com/ystar-foundation/yta-go"
)

func ExampleAPI_GetInfo() {
	api := yta.New(getAPIURL())

	info, err := api.GetInfo()
	if err != nil {
		panic(fmt.Errorf("get info: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
