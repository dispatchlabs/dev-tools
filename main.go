package main

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/sdk"
	"github.com/dispatchlabs/dev-tools/common-util/config"
	configHelper "github.com/dispatchlabs/dev-tools/common-util/helper"
	"github.com/dispatchlabs/dev-tools/transactions/helper"
	"os"
	"time"
	"flag"
	"github.com/dispatchlabs/dev-tools/transactions"
	"strconv"
	"github.com/dispatchlabs/disgo/commons/crypto"
	"encoding/hex"
	"math/big"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
 	"errors"
)

//var delay = time.Millisecond * 2
var txCount = 1
var testMap map[string]time.Time
//var queueTimeout = time.Second * 5

// would like to get this set up as a true CLI.
// like: https://github.com/spf13/cobra
// Right now, it's the quick and dirty to speed up work

var (
	version string
	date    string
)


func main() {

	//if version == "" {
	//	fmtd := `date \"+%Y-%m-%d-%H:%M:%S\"`
	//	msg := fmt.Sprintf("You have not specified a version.  You should build with those parameters eg: \n%s\n", "go build -ldflags \"-X main.version=0.0.1 -X main.date=%s", fmtd)
	//	fmt.Printf(msg)
	//}
	//fmt.Printf("version=%s, date=%s\n\n", version, date)

	var delegate = flag.String("delegate", "", "please pass in a delegate ip")
	var seed = flag.String("seed", "", "please pass in a seed ip")
	flag.Parse()
	functionToExecute := flag.Arg(0)
	toAddress := "61c48d9a7838021b55cbaf70956e7e7470ad0c8e"


	fmt.Println(*delegate)
	fmt.Println(*seed)
	fmt.Printf("Executing: %s\n", functionToExecute)
	switch functionToExecute {
	case "upgradeTx":
		tx := helper.NewUpdateTx( "1.0.1")
		dlgt := helper.GetRandomDelegate(transactions.SeedHost)
		// Post transaction.
		httpResponse, err := http.Post(fmt.Sprintf("http://%s:%d/v1/transactions", dlgt.HttpEndpoint.Host, dlgt.HttpEndpoint.Port), "application/json", bytes.NewBuffer([]byte(tx.String())))
		if err != nil {
			utils.Error(err)
		}
		defer httpResponse.Body.Close()

		// Read body.
		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf("%s\n", body)
		// Unmarshal response.
		var response *types.Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf("%s\n", tx.ToPrettyJson())
	case "upgradeSeedTx":
		tx := helper.NewUpdateTx( "1.0.1")
		httpResponse, err := http.Post(fmt.Sprintf("http://%s/v1/transactions", transactions.SeedHost), "application/json", bytes.NewBuffer([]byte(tx.String())))
		if err != nil {
			utils.Error(err)
		}
		defer httpResponse.Body.Close()

		// Read body.
		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			utils.Error(err)
		}

		// Unmarshal response.
		var response *types.Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf(tx.ToPrettyJson())
	case "manualTx":
		tx := helper.NewTranferTx( toAddress)
		fmt.Printf("\n%s\n", tx.ToPrettyJson())
	case "upgrade":
		output := configHelper.Update(configHelper.GetDefaultDirectory(), "2.5.0", "Disgo")
		utils.Info(output)
	case "newLocalCluster":
		var nbrDelegates int
		var err error
		genesisAccount := &types.Account{}
		genesisAccount.Balance = big.NewInt(int64(1800000000000000000))
		genesisAccount.Address = transactions.GenesisAddress
		if len(flag.Args()) == 4 {
			fmt.Printf("4 args")
			nbrDelegates, err = strconv.Atoi(flag.Arg(1))
			if err != nil {
				utils.Error(err)
			}
			genesisAccount.Address = flag.Arg(2)
			val, err := strconv.Atoi(flag.Arg(3))
			if err != nil {
				utils.Error(err)
			}
			genesisAccount.Balance = big.NewInt(int64(val))

			config.CleanAndBuildNewConfig(1, nbrDelegates, genesisAccount, true)
		}
		if(len(flag.Args()) == 2) {
			fmt.Printf("2 args")
			nbrDelegates, err = strconv.Atoi(flag.Arg(1))
			if err != nil {
				utils.Error(err)
			}
			config.CleanAndBuildNewConfig(1, nbrDelegates, genesisAccount, true)
		} else {
			fmt.Printf("1 args")
			config.CleanAndBuildNewConfig(1, 4, genesisAccount, true)
		}
	case "makePrivateKey", "test":
		createGenesisKeyAndAddress()
	case "update":
		config.RefreshDisgoExecutable("")
		configHelper.Exec("pkill disgo")
	case "clearDB":
		config.ClearDB(configHelper.GetDefaultDirectory())
	case "reset":
		config.RefreshDisgoExecutable("")
		config.ClearDB(configHelper.GetDefaultDirectory())
		config.ClearLogs(configHelper.GetDefaultDirectory())
		configHelper.Exec("pkill disgo")

	case "execute":
		if(len(flag.Args()) == 2) {
			txCount, _ = strconv.Atoi(flag.Arg(1))
		}
		if (*seed == ""){
			helper.RunTransferTokens(txCount, transactions.GenesisPrivateKey, transactions.GenesisAddress, toAddress, transactions.SeedHost, 1)
		} else {
			helper.RunTransferTokens(txCount, "a544cca72d88a49ec3afadc4a358125a138ad83b6fce72c1067b5773d8ae688f", "7bf5580620d91b55258a09dc9c1166f5542cc115", toAddress, *seed, 1)
		}

	case "transfer":
		if(len(flag.Args()) == 5) {
			privateKey := flag.Arg(1)
			from := flag.Arg(2)
			to   := flag.Arg(3)
			amount, _ := strconv.Atoi(flag.Arg(4))

			helper.RunTransferTokens(txCount, privateKey, from, to, transactions.SeedHost, int64(amount))
		} else {
			to   := flag.Arg(1)
			amount, _ := strconv.Atoi(flag.Arg(2))
			helper.RunTransferTokens(txCount, transactions.GenesisPrivateKey, transactions.GenesisAddress, to, transactions.SeedHost, int64(amount))
		}

	case "getTxByHash":
		if(len(flag.Args()) != 2) {
			fmt.Printf("Need to specify the transaction hash")
			return
		}
		txHash := flag.Arg(1)

		delegates, err := helper.GetDelegates(transactions.SeedHost)
		if err != nil {
			utils.Fatal(err)
		}
		time.Sleep(time.Second * 1)
		for _, delegate := range delegates {
			tx, err := sdk.GetTransaction(delegate, txHash)
			if err != nil {
				utils.Error("Error getting Transaction ", err)
			}
			if tx == nil {
				fmt.Printf("Transaction from Delegate: %s is not found yet\n", delegate.String())
			} else {
				fmt.Printf("Transaction from Delegate: %s is \n%s\n", delegate.String(), tx.ToPrettyJson())
			}
		}
	case "balance":
		delegates, err := helper.GetDelegates(transactions.SeedHost)
		if err != nil {
			utils.Error(err)
		}
		for _, delegate := range delegates {
			account, err := sdk.GetAccount(delegate, toAddress)
			if err != nil {
				utils.Error(err)
				continue
			}
			fmt.Printf("Account from Delegate: %s is \n%s\n", delegate.String(), account.ToPrettyJson())
		}

	case "deployContractFromFile":
		var receipt *types.Receipt
		var err error
		if(len(flag.Args()) == 2) {
			files := make([]string, 2)
			files[0] = fmt.Sprintf("test_contracts/%s.bin", flag.Args()[1])
			files[1] = fmt.Sprintf("test_contracts/%s.abi", flag.Args()[1])
			receipt, err = helper.DeployContractFromFile(files, *seed)
		} else if(len(flag.Args()) == 3) {
			receipt, err = helper.DeployContractFromFile(flag.Args()[1:], *seed)
		} else {
			files := make([]string, 2)
			files[0] = "test_contracts/Uint256Test.bin"
			files[1] = "test_contracts/Uint256Test.abi"
			receipt, err = helper.DeployContractFromFile(files, *seed)
		}
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf("Receipt: \n%s\n", receipt.ToPrettyJson())

	case "executeWrite":
		fmt.Println(flag.Arg(1))
		fmt.Println(flag.Arg(2))
		fmt.Println(flag.Args()[3:])


		if(len(flag.Args()) == 3) {
			helper.ExecuteContract(flag.Arg(1), flag.Arg(2), nil, *seed)
		} else if(len(flag.Args()) > 3) {
			helper.ExecuteContract(flag.Arg(1), flag.Arg(2), flag.Args()[3:], *seed)
		} else {
			helper.ExecuteContract("0b28be714a683eb119125ecb176724dcd701a597", "set", nil, *seed)
		}

	case "executeRead":
		var receipt = new(types.Receipt)
		var err = errors.New("")
		if(len(flag.Args()) == 3) {
			receipt, err = helper.CallContract(flag.Arg(1), flag.Arg(2), nil, *seed)
		} else if(len(flag.Args()) > 3) {
			receipt, err = helper.CallContract(flag.Arg(1), flag.Arg(2), flag.Args()[3:], *seed)
		} else {
			receipt, err = helper.CallContract("0b28be714a683eb119125ecb176724dcd701a597", "set", nil, *seed)
		}
		if err != nil {
			utils.Error(err)
		} else {
			utils.Info("Result: ", receipt)
		}

	case "executeVarArgContract":
		if len(os.Args) < 4 {
			fmt.Println("executeVarArgContract must have at least 3 arguments")
			break
		}
		helper.ExecuteContract(os.Args[2], os.Args[3], os.Args[4:], *seed)

	case "deployAndExecute":
		deployReceipt, err := helper.DeployContractFromFile(flag.Args()[1:], *seed)
		if err != nil {
			utils.Error(err)
			return
		}
		fmt.Printf("Deploy Receipt: \n%s\n", deployReceipt.ToPrettyJson())

		executeReceipt, err := helper.ExecuteContract(deployReceipt.ContractAddress, flag.Arg(2), nil, *seed)
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf("Execute Receipt: \n%s\n", executeReceipt.ToPrettyJson())

	case "createSignedTx":
		// Create transfer tokens transaction.
		txCount, _ = strconv.Atoi(flag.Arg(4))
		transaction, err := types.NewTransferTokensTransaction(flag.Arg(1), flag.Arg(2), flag.Arg(3), int64(txCount), 0, utils.ToMilliSeconds(time.Now()))
		if err != nil {
			utils.Error(err)
		}
		fmt.Printf(transaction.ToPrettyJson())
	default:
		fmt.Errorf("Invalid argument %s\n", functionToExecute)
	}
	//testMap = map[string]time.Time{}

}

func createGenesisKeyAndAddress() {
	publicKey, privateKey := crypto.GenerateKeyPair()
	address := crypto.ToAddress(publicKey)

	fmt.Printf("\nGenesis key & address: \n" +
		"***********************************************************************************************************************************************************\n" +
		"Public Key:\t\t%v\n" +
		"Private Key:\t\t%v\n" +
		"Address:\t\t%s\n" +
		"***********************************************************************************************************************************************************\n",
		hex.EncodeToString(publicKey), hex.EncodeToString(privateKey), hex.EncodeToString(address))

}

func getTxResult(seedUrl, hash string) {
	delegates, err := sdk.GetDelegates(seedUrl)
	if err != nil {
		utils.Fatal(err)
	}
	time.Sleep(time.Second * 1)
	for _, delegate := range delegates {
		tx, err := sdk.GetTransaction(delegate, hash)
		if err != nil {
			utils.Error("Error getting Transaction ", err)
		}
		if tx == nil {
			fmt.Printf("Transaction from Delegate: %s is not found yet\n", delegate.String())
		} else {
			fmt.Printf("Transaction from Delegate: %s is \n%s\n", delegate.String(), tx.ToPrettyJson())
		}
	}
}

//func sendGrpcTransactions(toAddress string) *types.Transaction {
//	var tx *types.Transaction
//
//	for i := 0; i < txCount; i++ {
//		tx = helper.GetTransaction(toAddress)
//		gossipResponse, err := SendGrpcTransaction(tx, getRandomDelegate().GrpcEndpoint, toAddress)
//		if err != nil {
//			utils.Error(err)
//		} else {
//			//fmt.Printf("grpc response: %v\n", gossipResponse)
//			fmt.Printf("Transaction Hash: %v\n", gossipResponse.Transaction.Hash)
//		}
//		time.Sleep(delay)
//	}
//	return tx
//}


