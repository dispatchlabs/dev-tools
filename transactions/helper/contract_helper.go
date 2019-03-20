package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/commons/helper"
	"github.com/dispatchlabs/disgo/sdk"
	"github.com/dispatchlabs/dev-tools/transactions"
	"fmt"
	"time"
	"github.com/dispatchlabs/dev-tools/common-util/util"
	"github.com/pkg/errors"
	"encoding/hex"
)


func DeployContract(contractCode, abi string) (*types.Receipt, error) {
	deployHash, err := sdk.DeploySmartContract(
		GetRandomDelegate(transactions.SeedHost),
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		contractCode,
		abi)
	if err != nil {
		return nil, err
	}
	time.Sleep(3 * time.Second)
	deployRcpt := GetReceipt(deployHash)
	return deployRcpt, nil
}

func DeployContractFromFile(args []string) (*types.Receipt, error) {
	if len(args) != 2 {
		return nil, errors.New("deployContractFromFile needs a binary file (arg 1) and abi file (arg 2)")
	}
	code, err := util.ReadFileAsString(args[0])
	if err != nil {
		return nil, err
	}

	abi, err := util.ReadFileAsString(args[1])
	if err != nil {
		return nil, err
	}

	printAbiMethodsAndParams(abi)

	deployHash, err := sdk.DeploySmartContract(
		GetRandomDelegate(transactions.SeedHost),
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		string(code),
		string(abi),
	)

	if err != nil {
		utils.Error(err)
	}
	time.Sleep(3 * time.Second)
	deployRcpt := GetReceipt(deployHash)
	return deployRcpt, nil
}

func ExecuteContract(contractAddress string, method string, args []string) (*types.Receipt, error) {
	var params string
	if args == nil {
		params = GetParamsForMethod(method)
	} else {
		params = GetVariableParamsForContract(args)
	}
	utils.Info(fmt.Sprintf("executeContract --> \t%s\t%s\t%s", contractAddress, method, params))
	hash, err := sdk.ExecuteWriteTransaction(
		GetRandomDelegate(transactions.SeedHost),
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		contractAddress,
		method,
		params,
	)
	if err != nil {
		utils.Error(err)
	}
	time.Sleep(3 * time.Second)
	execRcpt := GetReceipt(hash)

	return execRcpt, nil
}

func CallContract(contractAddress string, method string, args []string) (*types.Receipt, error) {
	var params string
	if args == nil {
		params = GetParamsForMethod(method)
	} else {
		params = GetVariableParamsForContract(args)
	}
	utils.Info(fmt.Sprintf("executeReadSmartContract --> \t%s\t%s\t%s", contractAddress, method, params))
	receipt, err := sdk.ExecuteReadTransaction(
		GetRandomDelegate(transactions.SeedHost),
		transactions.GenesisPrivateKey,
		transactions.GenesisAddress,
		contractAddress,
		method,
		params,
	)
	if err != nil {
		utils.Error(err)
		return nil, err
	} else {
		return receipt, nil
	}

}


func validateAbi(deployedTxHash string) {
	contractTx, err := sdk.GetTransaction(GetRandomDelegate(transactions.SeedHost), deployedTxHash)
	if err != nil {
		utils.Error(err, utils.GetCallStackWithFileAndLineNumber())
		return
	}
	params, err := helper.GetConvertedParams(contractTx)
	if err != nil {
		utils.Error(err, utils.GetCallStackWithFileAndLineNumber())
		return
	}
	fmt.Printf("params being sent are %s\n", params)
}

func printAbiMethodsAndParams(abi string) error {
	encoded := hex.EncodeToString([]byte(abi))
	theABI, err := helper.GetABI(encoded)
	if err != nil {
		utils.Error(err)
		return err
	}

	for k, v := range theABI.Methods {
		fmt.Printf("Method: %s\n", k)
		for i := 0; i < len(v.Inputs); i++ {
			arg := v.Inputs[i]
			fmt.Printf("\tInput: %s\n", arg.Type)

		}
	}
	return nil
}
