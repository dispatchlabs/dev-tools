package helper

import (
	"github.com/dispatchlabs/disgo/sdk"
	"github.com/dispatchlabs/disgo/commons/utils"
	"time"
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
)

func RunTransferTokens(txCount int, privateKey, fromAddress, toAddress, seedIp string, amount int64) {
	hashes := sendHttpTransactions(txCount, privateKey, fromAddress, toAddress, seedIp, amount)
	time.Sleep(time.Second * 15)
	getTxResult(seedIp, hashes)
	printBalances(seedIp, toAddress)
}


func sendHttpTransactions(txCount int, privateKey, fromAddress, toAddress, seedIp string, amount int64) []string {
	hashes := make([]string, txCount)
	var err error
	for i := 0; i < txCount; i++ {
		dlgt := GetRandomDelegate(seedIp)
		hashes[i], err = sdk.TransferTokens(dlgt, privateKey, fromAddress, toAddress, amount)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf("%d\tSent Transaction: %s to: %d\n", (i+1), hashes[i], dlgt.HttpEndpoint.Port)
	}
	return hashes
}

func getTxResult(seedUrl string, hashes []string) {
	delegates, err := GetDelegates(seedUrl)
	if err != nil {
		utils.Error(err)
	}
	for _, hash := range hashes {
		verifyTx(delegates, hash)
	}
}

func verifyTx(delegates []types.Node, hash string) {
	//utils.Info("veryify")
	list := make([]*types.Transaction, len(delegates))

	for index, delegate := range delegates {
		tx, err := sdk.GetTransaction(delegate, hash)
		if err != nil {
			msg := fmt.Sprintf("Error getting Transaction for hash: %s for delegate %d", hash, delegate.HttpEndpoint.Port)
			utils.Error(msg, err)
			return
		}
		if tx == nil {
			fmt.Printf("Transaction from Delegate: %s does not exist\n", delegate.String())
		} else if tx.Receipt.Status != "Ok" {
			printIssues(delegate, tx)

		} else {
			//fmt.Printf("Transaction from Delegate: %s is \n%s\n", delegate.String(), tx.ToPrettyJson())
		}
		list[index] = tx
	}
}

func printBalances(seedUrl, toAddress string) {
	delegates, err := GetDelegates(seedUrl)
	if err != nil {
		utils.Error(err)
	}
	fmt.Printf("********************************************************************************\n")
	for _, delegate := range delegates {
		account, err := sdk.GetAccount(delegate, toAddress)
		if err != nil {
			utils.Error(err)
			continue
		}
		fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
	}
	fmt.Printf("********************************************************************************\n")

}

func printIssues(delegate types.Node, transaction *types.Transaction) {
	reasons := make([]string, 0)
	if transaction.Receipt.Status != "Ok" {
		msg := fmt.Sprintf("Transaction %s from Delegate: %s had a status: %s\n", transaction.Hash, delegate.String(), transaction.Receipt.Status)
		fmt.Printf("\n%s\n", msg)
		reasons = append(reasons, msg)
	}
	fmt.Printf("%s\n\n", transaction.ToPrettyJson())
}