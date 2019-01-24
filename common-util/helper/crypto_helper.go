package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/crypto"
	"encoding/json"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
)

func GetEncryptedPrivateKey(account *types.Account) (*crypto.EncryptedKeyJSONV3, error) {
	jsn := &crypto.EncryptedKeyJSONV3{}
	password := "Disgo"
	keystore, err := types.CreateFromKey(account.PrivateKey, password)
	if err != nil {
		return nil, err
	}
	fmt.Printf("JSON: %s\n\n", string(keystore))
	err = json.Unmarshal(keystore, jsn)
	if err != nil {
		utils.Error(err)
	}
	return jsn, nil
}
