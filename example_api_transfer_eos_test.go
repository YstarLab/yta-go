package yta_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	yta "github.com/ystar-foundation/yta-go"
	"github.com/ystar-foundation/yta-go/token"
)

func ExampleAPI_PushTransaction_transfer_EOS() {
	api := yta.New(getAPIURL())

	keyBag := &yta.KeyBag{}
	err := keyBag.ImportPrivateKey(readPrivateKey())
	if err != nil {
		panic(fmt.Errorf("import private key: %s", err))
	}
	api.SetSigner(keyBag)

	from := yta.AccountName("eosuser1")
	to := yta.AccountName("eosuser2")
	quantity, err := yta.NewEOSAssetFromString("1.0000 YTA")
	memo := ""

	if err != nil {
		panic(fmt.Errorf("invalid quantity: %s", err))
	}

	txOpts := &yta.TxOptions{}
	if err := txOpts.FillFromChain(api); err != nil {
		panic(fmt.Errorf("filling tx opts: %s", err))
	}

	tx := yta.NewTransaction([]*yta.Action{token.NewTransfer(from, to, quantity, memo)}, txOpts)
	signedTx, packedTx, err := api.SignTransaction(tx, txOpts.ChainID, yta.CompressionNone)
	if err != nil {
		panic(fmt.Errorf("sign transaction: %s", err))
	}

	content, err := json.MarshalIndent(signedTx, "", "  ")
	if err != nil {
		panic(fmt.Errorf("json marshalling transaction: %s", err))
	}

	fmt.Println(string(content))
	fmt.Println()

	response, err := api.PushTransaction(packedTx)
	if err != nil {
		panic(fmt.Errorf("push transaction: %s", err))
	}

	fmt.Printf("Transaction [%s] submitted to the network succesfully.\n", hex.EncodeToString(response.Processed.ID))
}

func readPrivateKey() string {
	// Right now, the key is read from an environment variable, it's an example after all.
	// In a real-world scenario, would you probably integrate with a real wallet or something similar
	envName := "EOS_GO_PRIVATE_KEY"
	privateKey := os.Getenv(envName)
	if privateKey == "" {
		panic(fmt.Errorf("private key environment variable %q must be set", envName))
	}

	return privateKey
}
