package helper

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
)

func GetTransaction(toAddress string) *types.Transaction {
	utils.Info("GetTransaction")
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	var tx, err = types.NewTransferTokensTransaction(
		privateKey,
		from,
		toAddress,
		1,
		1,
		utils.ToMilliSeconds(time.Now().Add(time.Minute*4)),
	)
	if err != nil {
		utils.Error(err)
	} else {
		fmt.Printf("Created Tx: %s\n", tx.Hash)
	}

	return tx
}

func GetNewExecuteTxWithVarableParams(toAddress string, method string, args []string, write bool) *types.Transaction {
	utils.Info("GetNewExecuteTxWithVarableParams")
	// Taken from Genesis
	var privateKey = "0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a"
	var from = "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2c"

	//abi, err := ioutil.ReadFile(abi_file)
	//if err != nil {
	//	panic(err)
	//}

	var tx, _ = types.NewExecuteContractTransaction(
		privateKey,
		from,
		toAddress,
		method,
		GetVariableParamsForContract(args),
		write,
		utils.ToMilliSeconds(time.Now()),
	)
	return tx
}

// wants args in the format int:23 string:foo uint:22
func GetVariableParamsForContract(args []string) string {
	var argList = make([]interface{}, 0)

	for _, arg := range args {
		data := strings.Split(arg, ":")
		kind, val := data[0], data[1]

		switch kind {
		case "uint":
			u64, _ := strconv.ParseUint(val, 10, 32)
			argList = append(argList, u64)
		case "int":
			i64, _ := strconv.ParseInt(val, 10, 32)
			argList = append(argList, i64)
		case "string":
			argList = append(argList, val)
		case "bytes":
			byteArray := []byte(val)
			argList = append(argList, byteArray)
		case "bool":
			if val == "true" {
				argList = append(argList, true)
			} else {
				argList = append(argList, false)
			}
		}
	}

	fmt.Println("ArgList:", argList)
	bytes, err := json.Marshal(argList)
	if err != nil {
		utils.Error(bytes)
	}
	return string(bytes)
}

func GetParamsForMethod(method string) string {
	var params = make([]interface{}, 0)

	switch method {
	case "setAProxy":
		params = append(params, "8dbc69f38d71cf68757742e19d5642a3c200cbff")
		params = append(params, 25)
		break
	case "setA":
		params = append(params, 20)
		break
	case "setVar5":
		params = append(params, "Abcdefg")
		break
	case "setVar6Var4":
		params = append(params, "test value for var4")
		break
	case "intParam":
		params = append(params, 20)
		break
	case "plusOne":
		params = append(params, 1)
		break
	case "uintParam":
		params = append(params, uint(30))
		break
	case "boolParamType":
		params = append(params, true)
		break
	case "multiParams":
		params = append(params, "test1")
		params = append(params, "test2")
		break
	case "arrayParam":
		var array = make([]interface{}, 1)
		array[0] = uint(20)
		params = append(params, array)
		break
		//all fall through below
	case "set":
		bytes := []byte("someValueHere")
		params = append(params, "keyValue")
		params = append(params, bytes)
	case "test_me":
		params = append(params, uint(30))
		break
	case "getVar5":
	case "returnBool":
	case "testLog":
	case "returnInt":
	case "returnUint":
	case "throwException":

		break

	}
	fmt.Println("Params: %v\n", params)
	bytes, err := json.Marshal(params)
	if err != nil {
		utils.Error(bytes)
	}
	return string(bytes)
}

func GetCode() string {
	return "608060405234801561001057600080fd5b506040805190810160405280600d81526020017f61616161616161616161616161000000000000000000000000000000000000008152506000908051906020019061005c9291906100f8565b5060006002600001819055506000600260010160006101000a81548160ff0219169083151502179055506001600260010160016101000a81548160ff021916908360ff1602179055506040805190810160405280600b81526020017f62626262626262626262620000000000000000000000000000000000000000008152506002800190805190602001906100f29291906100f8565b5061019d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013957805160ff1916838001178555610167565b82800160010185558215610167579182015b8281111561016657825182559160200191906001019061014b565b5b5090506101749190610178565b5090565b61019a91905b8082111561019657600081600090555060010161017e565b5090565b90565b610e7f80620001ad6000396000f300608060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c6f56a1461010c5780631af35da3146101b25780631e358c3e146101e1578063216a52e51461025b578063222e04121461030a57806330bc6db21461032157806333e538e91461034c57806334e45f53146103dc5780633a458b1f146104455780634a846e02146104f45780634aea8b14146105f057806378d8866e1461069657806379af647314610726578063943640c31461073d578063a5fe087214610794578063af445500146107ab578063cb69e300146107d6578063d13f25ad1461083f578063fd213d0c1461086a575b600080fd5b34801561011857600080fd5b5061013760048036038101908080359060200190929190505050610912565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017757808201518184015260208101905061015c565b50505050905090810190601f1680156101a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101be57600080fd5b506101c7610951565b604051808215151515815260200191505060405180910390f35b3480156101ed57600080fd5b5061024560048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610959565b6040518082815260200191505060405180910390f35b34801561026757600080fd5b50610308600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061097b565b005b34801561031657600080fd5b5061031f61097f565b005b34801561032d57600080fd5b50610336610993565b6040518082815260200191505060405180910390f35b34801561035857600080fd5b5061036161099c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103a1578082015181840152602081019050610386565b50505050905090810190601f1680156103ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103e857600080fd5b50610443600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610a3e565b005b34801561045157600080fd5b5061045a610a5a565b60405180858152602001841515151581526020018360ff1660ff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104b657808201518184015260208101905061049b565b50505050905090810190601f1680156104e35780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561050057600080fd5b50610509610b2a565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561054d578082015181840152602081019050610532565b50505050905090810190601f16801561057a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156105b3578082015181840152602081019050610598565b50505050905090810190601f1680156105e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156105fc57600080fd5b5061061b60048036038101908080359060200190929190505050610ba1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561065b578082015181840152602081019050610640565b50505050905090810190601f1680156106885780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106a257600080fd5b506106ab610be0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106eb5780820151818401526020810190506106d0565b50505050905090810190601f1680156107185780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561073257600080fd5b5061073b610c7e565b005b34801561074957600080fd5b50610752610c95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107a057600080fd5b506107a9610c9d565b005b3480156107b757600080fd5b506107c0610d40565b6040518082815260200191505060405180910390f35b3480156107e257600080fd5b5061083d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d46565b005b34801561084b57600080fd5b50610854610d60565b6040518082815260200191505060405180910390f35b34801561087657600080fd5b50610897600480360381019080803515159060200190929190505050610d69565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108d75780820151818401526020810190506108bc565b50505050905090810190601f1680156109045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b600080905090565b600081600081518110151561096a57fe5b906020019060200201519050919050565b5050565b600080600581151561098d57fe5b04905050565b60006014905090565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a345780601f10610a0957610100808354040283529160200191610a34565b820191906000526020600020905b815481529060010190602001808311610a1757829003601f168201915b5050505050905090565b80600280019080519060200190610a56929190610dae565b5050565b60028060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b205780601f10610af557610100808354040283529160200191610b20565b820191906000526020600020905b815481529060010190602001808311610b0357829003601f168201915b5050505050905084565b6060806040805190810160405280600581526020017f74657374310000000000000000000000000000000000000000000000000000008152506040805190810160405280600581526020017f7465737432000000000000000000000000000000000000000000000000000000815250915091509091565b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b505050505081565b600260000160008154809291906001019190505550565b600030905090565b7fb20aa8922321b2e5be1e9784294eda54d640a58038ceede50492f7d7ffc8ad62604051808060200180602001838103835260058152602001807f7465737431000000000000000000000000000000000000000000000000000000815250602001838103825260058152602001807f74657374320000000000000000000000000000000000000000000000000000008152506020019250505060405180910390a1565b60015481565b8060009080519060200190610d5c929190610dae565b5050565b60006014905090565b606060008290506040805190810160405280600481526020017f7465737400000000000000000000000000000000000000000000000000000000815250915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610def57805160ff1916838001178555610e1d565b82800160010185558215610e1d579182015b82811115610e1c578251825591602001919060010190610e01565b5b509050610e2a9190610e2e565b5090565b610e5091905b80821115610e4c576000816000905550600101610e34565b5090565b905600a165627a7a72305820dcd6870229051d00b5ba7652f6df20b6bbe95d60d4342eb36e22411fab359aa70029"
	//return  "608060405234801561001057600080fd5b506040805190810160405280600d81526020017f61616161616161616161616161000000000000000000000000000000000000008152506000908051906020019061005c9291906100f8565b5060006002600001819055506000600260010160006101000a81548160ff0219169083151502179055506001600260010160016101000a81548160ff021916908360ff1602179055506040805190810160405280600b81526020017f62626262626262626262620000000000000000000000000000000000000000008152506002800190805190602001906100f29291906100f8565b5061019d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061013957805160ff1916838001178555610167565b82800160010185558215610167579182015b8281111561016657825182559160200191906001019061014b565b5b5090506101749190610178565b5090565b61019a91905b8082111561019657600081600090555060010161017e565b5090565b90565b610e7f80620001ad6000396000f300608060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806304c6f56a1461010c5780631af35da3146101b25780631e358c3e146101e1578063216a52e51461025b578063222e04121461030a57806330bc6db21461032157806333e538e91461034c57806334e45f53146103dc5780633a458b1f146104455780634a846e02146104f45780634aea8b14146105f057806378d8866e1461069657806379af647314610726578063943640c31461073d578063a5fe087214610794578063af445500146107ab578063cb69e300146107d6578063d13f25ad1461083f578063fd213d0c1461086a575b600080fd5b34801561011857600080fd5b5061013760048036038101908080359060200190929190505050610912565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561017757808201518184015260208101905061015c565b50505050905090810190601f1680156101a45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101be57600080fd5b506101c7610951565b604051808215151515815260200191505060405180910390f35b3480156101ed57600080fd5b5061024560048036038101908080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284378201915050505050509192919290505050610959565b6040518082815260200191505060405180910390f35b34801561026757600080fd5b50610308600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061097b565b005b34801561031657600080fd5b5061031f61097f565b005b34801561032d57600080fd5b50610336610993565b6040518082815260200191505060405180910390f35b34801561035857600080fd5b5061036161099c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103a1578082015181840152602081019050610386565b50505050905090810190601f1680156103ce5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103e857600080fd5b50610443600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610a3e565b005b34801561045157600080fd5b5061045a610a5a565b60405180858152602001841515151581526020018360ff1660ff16815260200180602001828103825283818151815260200191508051906020019080838360005b838110156104b657808201518184015260208101905061049b565b50505050905090810190601f1680156104e35780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b34801561050057600080fd5b50610509610b2a565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561054d578082015181840152602081019050610532565b50505050905090810190601f16801561057a5780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156105b3578082015181840152602081019050610598565b50505050905090810190601f1680156105e05780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156105fc57600080fd5b5061061b60048036038101908080359060200190929190505050610ba1565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561065b578082015181840152602081019050610640565b50505050905090810190601f1680156106885780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156106a257600080fd5b506106ab610be0565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156106eb5780820151818401526020810190506106d0565b50505050905090810190601f1680156107185780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561073257600080fd5b5061073b610c7e565b005b34801561074957600080fd5b50610752610c95565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156107a057600080fd5b506107a9610c9d565b005b3480156107b757600080fd5b506107c0610d40565b6040518082815260200191505060405180910390f35b3480156107e257600080fd5b5061083d600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610d46565b005b34801561084b57600080fd5b50610854610d60565b6040518082815260200191505060405180910390f35b34801561087657600080fd5b50610897600480360381019080803515159060200190929190505050610d69565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108d75780820151818401526020810190506108bc565b50505050905090810190601f1680156109045780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b600080905090565b600081600081518110151561096a57fe5b906020019060200201519050919050565b5050565b600080600581151561098d57fe5b04905050565b60006014905090565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a345780601f10610a0957610100808354040283529160200191610a34565b820191906000526020600020905b815481529060010190602001808311610a1757829003601f168201915b5050505050905090565b80600280019080519060200190610a56929190610dae565b5050565b60028060000154908060010160009054906101000a900460ff16908060010160019054906101000a900460ff1690806002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b205780601f10610af557610100808354040283529160200191610b20565b820191906000526020600020905b815481529060010190602001808311610b0357829003601f168201915b5050505050905084565b6060806040805190810160405280600581526020017f74657374310000000000000000000000000000000000000000000000000000008152506040805190810160405280600581526020017f7465737432000000000000000000000000000000000000000000000000000000815250915091509091565b60606040805190810160405280600481526020017f74657374000000000000000000000000000000000000000000000000000000008152509050919050565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c765780601f10610c4b57610100808354040283529160200191610c76565b820191906000526020600020905b815481529060010190602001808311610c5957829003601f168201915b505050505081565b600260000160008154809291906001019190505550565b600030905090565b7fb20aa8922321b2e5be1e9784294eda54d640a58038ceede50492f7d7ffc8ad62604051808060200180602001838103835260058152602001807f7465737431000000000000000000000000000000000000000000000000000000815250602001838103825260058152602001807f74657374320000000000000000000000000000000000000000000000000000008152506020019250505060405180910390a1565b60015481565b8060009080519060200190610d5c929190610dae565b5050565b60006014905090565b606060008290506040805190810160405280600481526020017f7465737400000000000000000000000000000000000000000000000000000000815250915050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610def57805160ff1916838001178555610e1d565b82800160010185558215610e1d579182015b82811115610e1c578251825591602001919060010190610e01565b5b509050610e2a9190610e2e565b5090565b610e5091905b80821115610e4c576000816000905550600101610e34565b5090565b905600a165627a7a7230582041658a57dddbe6b1c9f0bfce0db152db8095a3035b6d6d0ff37666dafdd367f30029"
}

func GetAbi() string {
	return `[
	{
		"constant": true,
		"inputs": [
			{
				"name": "param",
				"type": "int256"
			}
		],
		"name": "intParam",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "returnBool",
		"outputs": [
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "param",
				"type": "uint256[]"
			}
		],
		"name": "arrayParam",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "value",
				"type": "string"
			},
			{
				"name": "value2",
				"type": "string"
			}
		],
		"name": "multiParams",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "throwException",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "returnUint",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getVar5",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "value",
				"type": "string"
			}
		],
		"name": "setVar6Var4",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "var6",
		"outputs": [
			{
				"name": "var1",
				"type": "uint256"
			},
			{
				"name": "var2",
				"type": "bool"
			},
			{
				"name": "var3",
				"type": "uint8"
			},
			{
				"name": "var4",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "getMultiReturn",
		"outputs": [
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "param",
				"type": "uint256"
			}
		],
		"name": "uintParam",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "var5",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "incVar6Var1",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "returnAddress",
		"outputs": [
			{
				"name": "",
				"type": "address"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [],
		"name": "logEvent",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "intVar",
		"outputs": [
			{
				"name": "",
				"type": "int256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "value",
				"type": "string"
			}
		],
		"name": "setVar5",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "returnInt",
		"outputs": [
			{
				"name": "",
				"type": "int256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "value",
				"type": "bool"
			}
		],
		"name": "boolParamType",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"anonymous": false,
		"inputs": [
			{
				"indexed": false,
				"name": "test1",
				"type": "string"
			},
			{
				"indexed": false,
				"name": "test2",
				"type": "string"
			}
		],
		"name": "testLog",
		"type": "event"
	}
]`
}

func GetDaveCode() string {
	return "6080604052348015600f57600080fd5b50609c8061001e6000396000f300608060405260043610603e5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663f5a6259f81146043575b600080fd5b348015604e57600080fd5b506058600435606a565b60408051918252519081900360200190f35b600101905600a165627a7a7230582052a887255bee69b86c68b80729c39cb6d8c2651404d12f5b12ce002ebf8f1b0b0029"
}

func GetDaveAbi() string {
	return `[{"constant":true,"inputs":[{"name":"y","type":"uint256"}],"name":"plusOne","outputs":[{"name":"x","type":"uint256"}],"payable":false,"stateMutability":"pure","type":"function"}]`
}

func GetComplexContractCode() string {
	return "608060405234801561001057600080fd5b5060df8061001f6000396000f3006080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063d46300fd14604e578063ee919d50146076575b600080fd5b348015605957600080fd5b50606060a0565b6040518082815260200191505060405180910390f35b348015608157600080fd5b50609e6004803603810190808035906020019092919050505060a9565b005b60008054905090565b80600081905550505600a165627a7a723058205906547745a52855a1b22685e079cbdec04bad5d24c4c243d60837b39fb845890029"
}

func GetComplexContractAbi() string {
	return `[
	{
		"constant": true,
		"inputs": [],
		"name": "getA",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "a",
				"type": "uint256"
			}
		],
		"name": "setA",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
}

func GetContractToContractCode() string {
	return "608060405234801561001057600080fd5b5061025f806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063f3189d4614610051578063f474bd9e146100a8575b600080fd5b34801561005d57600080fd5b50610092600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506100f5565b6040518082815260200191505060405180910390f35b3480156100b457600080fd5b506100f3600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291905050506101a2565b005b6000808290508073ffffffffffffffffffffffffffffffffffffffff1663d46300fd6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561015f57600080fd5b505af1158015610173573d6000803e3d6000fd5b505050506040513d602081101561018957600080fd5b8101908080519060200190929190505050915050919050565b60008290508073ffffffffffffffffffffffffffffffffffffffff1663ee919d50836040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180828152602001915050600060405180830381600087803b15801561021657600080fd5b505af115801561022a573d6000803e3d6000fd5b505050505050505600a165627a7a72305820357833ed1cddada528e80a109eb86fdea4cadc72eba11ed595a73294811c87920029"
}

func GetContractToContractAbi() string {
	return `[
	{
		"constant": true,
		"inputs": [
			{
				"name": "originalContract",
				"type": "address"
			}
		],
		"name": "getAPrxoy",
		"outputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "originalContract",
				"type": "address"
			},
			{
				"name": "a",
				"type": "uint256"
			}
		],
		"name": "setAProxy",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`
}

func GetDenisContractCode() string {
	return "608060405234801561001057600080fd5b506101cf806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634a846e028114610045575b600080fd5b34801561005157600080fd5b5061005a610138565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561009b578181015183820152602001610083565b50505050905090810190601f1680156100c85780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156100fb5781810151838201526020016100e3565b50505050905090810190601f1680156101285780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60408051808201825260058082527f746573743100000000000000000000000000000000000000000000000000000060208084019190915283518085019094529083527f746573743200000000000000000000000000000000000000000000000000000090830152915600a165627a7a723058203e88e2519dc4b1f0aedbf85899eecfdafd91bd94e2854b0559923a0f0dddeee50029"
}

func GetDenisContractAbi() string {
	return `
[
	{
		"constant": true,
		"inputs": [],
		"name": "getMultiReturn",
		"outputs": [
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`
}

func GetAveryContractCode() string {
	return "608060405234801561001057600080fd5b50610464806100206000396000f30060806040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632b29c0fa14610051578063693ec85e14610118575b600080fd5b34801561005d57600080fd5b506100fe600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506101fa565b604051808215151515815260200191505060405180910390f35b34801561012457600080fd5b5061017f600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610286565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101bf5780820151818401526020810190506101a4565b50505050905090810190601f1680156101ec5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000816000846040518082805190602001908083835b6020831015156102355780518252602082019150602081019050602083039250610210565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020908051906020019061027b929190610393565b506001905092915050565b60606000826040518082805190602001908083835b6020831015156102c0578051825260208201915060208101905060208303925061029b565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103875780601f1061035c57610100808354040283529160200191610387565b820191906000526020600020905b81548152906001019060200180831161036a57829003601f168201915b50505050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106103d457805160ff1916838001178555610402565b82800160010185558215610402579182015b828111156104015782518255916020019190600101906103e6565b5b50905061040f9190610413565b5090565b61043591905b80821115610431576000816000905550600101610419565b5090565b905600a165627a7a72305820ea80fc901400a55aa856e5ede33521d0d6e5971ad71e89abb43fe21f2f2209ee0029"
}

func GetAveryContractAbi() string {
	return `[{"constant":false,"inputs":[{"name":"key","type":"string"},{"name":"value","type":"bytes"}],"name":"set","outputs":[{"name":"","type":"bool"}],"payable":false,"stateMutability":"nonpayable","type":"function"},{"constant":true,"inputs":[{"name":"key","type":"string"}],"name":"get","outputs":[{"name":"","type":"bytes"}],"payable":false,"stateMutability":"view","type":"function"}]`
}
