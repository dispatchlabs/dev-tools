package helper

import (
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/commons/types"
	"time"
	"fmt"
	"github.com/dispatchlabs/tools/transactions"
)

func NewTranferTx(toAddress string) *types.Transaction {
	utils.Info("NewTranferTx")
	transaction, err := types.NewTransferTokensTransaction(
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		toAddress,
		999,
		0,
		utils.ToMilliSeconds(time.Now().Add(time.Minute)),
	)

	if err != nil {
		return nil
	}
	return transaction
}

func NewManualDeployTx(code, abi string) *types.Transaction {
	utils.Info("NewManualDeployTx")
	tx, _ := types.NewDeployContractTransaction(
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		code,
		abi,
		utils.ToMilliSeconds(time.Now()),
	)
	//
	fmt.Printf("DEPLOY: %s\n", tx.ToPrettyJson())
	return tx
}

func NewManualExecuteTx(toAddress string, method string, readOnly bool) *types.Transaction {
	utils.Info("NewManualExecuteTx")
	var tx, _ = types.NewExecuteContractTransaction(
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		toAddress,
		method,
		GetParamsForMethod(method),
		utils.ToMilliSeconds(time.Now()),
		readOnly,
	)
	return tx
}

func NewUpdateTx(version string) *types.Transaction {
	utils.Info("NewUpdateTx")
	var tx, _ = types.NewUpdateTransaction(
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		version,
		utils.ToMilliSeconds(time.Now()),
	)
	return tx
}
