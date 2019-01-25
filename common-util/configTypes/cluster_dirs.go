package configTypes

import (
	"fmt"
	"github.com/dispatchlabs/dev-tools/common-util/util"
	"os"
)

type ClusterStructure struct {
	DisgoDir        	string
	ClusterRoot     	string
	NodeDirs        	map[string]string
	DelegateDirs    	map[string]string
	AccountFileName 	string
	ConfigFileName  	string
	KeyFileName			string
	GenesisAccountName	string
}

func NewClusterStructure(disgoDir, clusterRoot string, nbrSeeds, nbrDelegates int) *ClusterStructure {

	return &ClusterStructure{
		DisgoDir:        	disgoDir,
		ClusterRoot:     	clusterRoot,
		NodeDirs:        	getNodeDirs(clusterRoot, nbrSeeds, nbrDelegates),
		AccountFileName: 	"account.json",
		ConfigFileName:  	"config.json",
		KeyFileName:		"myDisgoKey.json",
		GenesisAccountName:	"genesis_account.json",
	}
}

func getNodeDirs(clusterRoot string, nbrSeeds, nbrDelegates int) map[string]string {
	nodeDirs := map[string]string{}
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		nodeDirs[seedName] = clusterRoot + string(os.PathSeparator) + seedName
	}
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)
		nodeDirs[delegateName] = clusterRoot + string(os.PathSeparator) + delegateName
	}
	return nodeDirs
}

func (this ClusterStructure) SaveAccountAndConfigFiles(nodeInfo *NodeInfo) {
	configDir := this.NodeDirs[nodeInfo.Name] + string(os.PathSeparator) + "config/"
	util.WriteFile(configDir, configDir+string(os.PathSeparator)+this.AccountFileName, nodeInfo.Account.ToPrettyJson())
	util.WriteFile(configDir, configDir+string(os.PathSeparator)+this.ConfigFileName, nodeInfo.Config.ToPrettyJson())
	util.WriteFile(configDir, configDir+string(os.PathSeparator)+this.KeyFileName, nodeInfo.EncryptedKey.ToPrettyJson())
}

func (this ClusterStructure) SaveGenesisAccount(nodeInfo *NodeInfo) {
	if nodeInfo.GenesisAccount == nil {
		return
	}
	configDir := this.NodeDirs[nodeInfo.Name] + string(os.PathSeparator) + "config/"
	genesisString := fmt.Sprintf("{\"address\": \"%s\",\"balance\": \"%d\"}", nodeInfo.GenesisAccount.Address, nodeInfo.GenesisAccount.Balance)
	util.WriteFile(configDir, configDir+string(os.PathSeparator)+this.GenesisAccountName, genesisString)
}

func (this ClusterStructure) getAccountFileLocation(nodeName string) string {
	baseDir := this.NodeDirs[nodeName]
	return baseDir + string(os.PathSeparator) + this.AccountFileName
}

func (this ClusterStructure) getConfigFileLocation(nodeName string) string {
	baseDir := this.NodeDirs[nodeName]
	return baseDir + string(os.PathSeparator) + this.ConfigFileName
}
