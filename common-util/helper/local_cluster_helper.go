package helper

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/dev-tools/common-util/configTypes"
	"os/user"
	"github.com/dispatchlabs/dev-tools/common-util/util"
)

/*
 * Functions specifically for creating local cluster to run
 */

// Restricted config specifies delegate list in the seed node so any nodes joining are a only type Node
// Non-Restricted config does not specify delegate list in the seed node so any nodes joining are a delegate
func CreateNewLocalConfigs(clusterStructure *configTypes.ClusterStructure, seedNodes []*configTypes.NodeInfo, delegateNodes []*configTypes.NodeInfo, restricted bool) map[string]*configTypes.NodeInfo {
	configMap := map[string]*configTypes.NodeInfo{}
	seedsConfig := make([]*types.Node, len(seedNodes))
	for i := 0; i < len(seedNodes); i++ {
		seedAccount := CreateSeedAccount()
		seedConfig := CreateSeedConfig(seedNodes[i].Host, seedNodes[i].HttpPort, seedNodes[i].GrpcPort, seedNodes[i].LocalPort, seedAccount)

		seedEncryptedKey, err := GetEncryptedPrivateKey(seedAccount)
		if err != nil {
			utils.Error(err)
		}

		seedNodes[i].Account = seedAccount
		seedNodes[i].Config = seedConfig
		seedNodes[i].EncryptedKey = seedEncryptedKey
		seedsConfig[i] = seedConfig.Seeds[0]
		configMap[seedNodes[i].Name] = seedNodes[i]
	}

	delegateAddressList := make([]string, len(delegateNodes))

	for i := 0; i < len(delegateNodes); i++ {
		delegateNodes[i].Config = CreateDelegateConfig(delegateNodes[i].Host, delegateNodes[i].HttpPort, delegateNodes[i].GrpcPort, delegateNodes[i].LocalPort, seedsConfig)
		delegateNodes[i].Account = CreateDelegateAccount(delegateNodes[i].Name)
		delegateEncryptedKey, err := GetEncryptedPrivateKey(delegateNodes[i].Account)
		if err != nil {
			utils.Error(err)
		}
		delegateNodes[i].EncryptedKey = delegateEncryptedKey
		delegateAddressList[i] = delegateNodes[i].Account.Address

		configMap[delegateNodes[i].Name] = delegateNodes[i]
	}
	if restricted {
		for _, seedNode := range seedNodes {
			seedNode.Config.DelegateAddresses = delegateAddressList
		}
	}
	return configMap
}

func SetupDefaultConfig(host string, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort int) {
	clusterStructure := configTypes.NewClusterStructure(GetDisgoDirectory(), GetDefaultDirectory(), nbrSeeds, nbrDelegates)

	seedNodes := make([]*configTypes.NodeInfo, nbrSeeds)
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		grpcPort := seedStartingPort + (i * 4)
		httpPort := grpcPort + 2

		seedInfo := &configTypes.NodeInfo{seedName, host, int64(httpPort), int64(grpcPort), int64(0), nil, nil, nil, nil}
		seedNodes[i] = seedInfo
	}

	delegateNodes := make([]*configTypes.NodeInfo, nbrDelegates)
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)

		grpcPort := delegateStartingPort + (i * 3)
		httpPort := grpcPort + 1
		localPort := grpcPort + 2

		delegateNodes[i] = &configTypes.NodeInfo{delegateName, host, int64(httpPort), int64(grpcPort), int64(localPort), nil, nil, nil, nil}

	}
	configMap := CreateNewLocalConfigs(clusterStructure, seedNodes, delegateNodes, true)
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}
}

var defaultDirectory string

func GetDefaultDirectory() string {
	if defaultDirectory == "" {
		usr, err := user.Current()
		if err != nil {
			utils.Fatal(err)
		}
		defaultDirectory = usr.HomeDir + "/disgo_cluster"
		util.MakeDir(defaultDirectory)
	}
	return defaultDirectory
}

var disgoDir string

func GetDisgoDirectory() string {
	if disgoDir == "" {
		usr, err := user.Current()
		if err != nil {
			utils.Fatal(err)
		}
		disgoDir = usr.HomeDir + "/go/src/github.com/dispatchlabs/disgo"
	}
	return disgoDir
}
