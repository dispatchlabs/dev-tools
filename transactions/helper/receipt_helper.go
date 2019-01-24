package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/sdk"
	"fmt"
	"time"
	"github.com/dispatchlabs/tools/transactions"
)

func GetReceipt(hash string) *types.Receipt {
	for {
		utils.Info("Get Reciept")
		tx, err := sdk.GetTransaction(GetRandomDelegate(transactions.SeedHost), hash)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf(tx.String())
		receipt := tx.Receipt
		fmt.Printf("Hash: %s\n%s\n", hash, receipt.ToPrettyJson())
		if receipt.Status == "Pending" {
			time.Sleep(time.Second * 5)
		} else {
			return &receipt
		}
	}
}

