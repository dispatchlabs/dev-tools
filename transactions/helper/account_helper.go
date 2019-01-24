package helper

import (
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/sdk"
	"github.com/dispatchlabs/tools/transactions"
	"time"
	"github.com/dispatchlabs/disgo/commons/types"
	"fmt"
)

func GetAccount(address string) *types.Account {
	for {
		utils.Info("Get Account")
		acct, err := sdk.GetAccount(GetRandomDelegate(transactions.SeedHost), address)
		if err != nil {
			utils.Error(err)
		}
		if acct.HertzAvailable == 0 || acct.Balance.Uint64() == acct.HertzAvailable {
			fmt.Printf("   Acct: %v :: %v\n", acct.Balance.Uint64(), acct.HertzAvailable)
			time.Sleep(time.Millisecond * 200)
		} else {
			return acct
		}
	}
}