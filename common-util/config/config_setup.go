package config

import (
	"fmt"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/dev-tools/common-util/configTypes"
	"github.com/dispatchlabs/dev-tools/common-util/helper"
	"github.com/dispatchlabs/dev-tools/common-util/util"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/dispatchlabs/disgo/commons/types"
)

func CleanAndBuildNewConfig(nbrSeeds, nbrDelegates int, genesisAccount *types.Account, restricted bool) {
	host := "127.0.0.1"
	delegateStartingPort := 3502
	seedStartingPort := 1973
	//clear out what is there
	defaultDir := helper.GetDefaultDirectory()
	util.DeleteSubDirs(defaultDir)

	//create new dirs with config
	SetupDefaultLocalCluster(host, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort, genesisAccount, restricted)
	RefreshDisgoExecutable(defaultDir)
}

func RefreshDisgoExecutable(baseDir string) {
	//build latest code
	if baseDir == "" {
		baseDir = helper.GetDefaultDirectory()
	}
	fmt.Println("Base Dir:")
	fmt.Println(baseDir)
	err := BuildDisgoExecutable()
	if err != nil {
		panic(err)
	}

	//update with newest disgo build for all directories
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			updateDisgoExecutable(file.Name())
		}
	}
}

func ClearLogs(baseDir string) {
	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.IsDir() {
			dir := filepath.Join(baseDir, file.Name())
			innerFiles, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range innerFiles {
				if file.Name() == "disgo.log" {
					innerMost := filepath.Join(dir, file.Name())
					fmt.Println(innerMost)
					util.DeleteFile(innerMost)
				}
			}
		}
	}
}

func ClearDB(baseDir string) {

	files, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(filepath.Join(baseDir, file.Name()))
		if file.IsDir() {
			dir := filepath.Join(baseDir, file.Name())
			innerFiles, err := ioutil.ReadDir(dir)
			if err != nil {
				log.Fatal(err)
			}
			for _, file := range innerFiles {
				if file.IsDir() && file.Name() == "db" {
					innerMost := filepath.Join(dir, file.Name())
					fmt.Println(innerMost)
					util.DeleteDir(innerMost)
				}
			}
		}
	}
}

func BuildDisgoExecutable() error {
	fmt.Println("BuildDisgoExecutable")
	util.DeleteFile(fmt.Sprintf("%s/disgo", helper.GetDisgoDirectory()))

	fmt.Println("getDisgoDirectory:")
	fmt.Println(helper.GetDisgoDirectory())

	helper.CheckCommand("go")
	cmd := "go build -ldflags \"-X main.version=2.4.0 -X main.date=`date \"+%Y-%m-%d-%H:%M:%S\"` -X github.com/dispatchlabs/disgo/commons/types.Password=test\""
	//cmd := "go build"
	//cmd := "ls -al"

	fmt.Println(cmd)


	err := helper.ExecFromDir(cmd, helper.GetDisgoDirectory())
	if err != nil {
		utils.Error(err)
		return err
	}
	return nil
}

func updateDisgoExecutable(nodeName string) {
	nodeDir := helper.GetDefaultDirectory() + string(os.PathSeparator) + nodeName + string(os.PathSeparator)
	//util.DeleteFile(fmt.Sprintf("%sdisgo", nodeDir))

	cmd := fmt.Sprintf("cp %s/disgo %s", helper.GetDisgoDirectory(), nodeDir)
	fmt.Println(cmd)

	utils.Debug("Command: " + cmd)
	err := helper.Exec(cmd)
	if err != nil {
		utils.Error(err)
	}
}

func SetupDefaultLocalCluster(host string, nbrSeeds, nbrDelegates, seedStartingPort, delegateStartingPort int, genesisAccount *types.Account, restricted bool) map[string]*configTypes.NodeInfo {
	clusterStructure := configTypes.NewClusterStructure(helper.GetDisgoDirectory(), helper.GetDefaultDirectory(), nbrSeeds, nbrDelegates)

	seedNodes := make([]*configTypes.NodeInfo, nbrSeeds)
	for i := 0; i < nbrSeeds; i++ {
		seedName := fmt.Sprintf("seed-%d", i)
		grpcPort := seedStartingPort + (i * 4)
		httpPort := grpcPort + 2

		seedInfo := &configTypes.NodeInfo{seedName, host, int64(httpPort), int64(grpcPort), 0,nil, nil, nil, genesisAccount}
		seedNodes[i] = seedInfo
	}

	delegateNodes := make([]*configTypes.NodeInfo, nbrDelegates)
	for i := 0; i < nbrDelegates; i++ {
		delegateName := fmt.Sprintf("delegate-%d", i)

		grpcPort := delegateStartingPort + (i * 3)
		httpPort := grpcPort + 1
		localPort := grpcPort + 2

		delegateNodes[i] = &configTypes.NodeInfo{delegateName, host, int64(httpPort), int64(grpcPort), int64(localPort),nil, nil, nil, genesisAccount}
	}
	configMap := helper.CreateNewLocalConfigs(clusterStructure, seedNodes, delegateNodes, restricted)
	for _, v := range configMap {
		clusterStructure.SaveAccountAndConfigFiles(v)
		v.GenesisAccount = genesisAccount
		clusterStructure.SaveGenesisAccount(v)
	}
	return configMap
}

func CreateRemoteClusterConfig(seedNodes, delegateNodes []*configTypes.NodeInfo) map[string]*configTypes.NodeInfo {
	//seedNode := &configTypes.NodeInfo{"stage-seed-0", "35.203.143.69", 1975,1973, nil, nil}
	//
	//delegateNodes = append(delegateNodes, &configTypes.NodeInfo{"stage-delegate-0", "35.233.231.3", 1975, 1973, nil, nil})
	//delegateNodes = append(delegateNodes, &configTypes.NodeInfo{"stage-delegate-1", "35.233.241.115", 1975, 1973, nil, nil})
	//delegateNodes = append(delegateNodes, &configTypes.NodeInfo{"stage-delegate-2", "35.230.0.126", 1975, 1973, nil, nil})

	configMap := helper.GetNewRemoteConfigs(seedNodes, delegateNodes)
	for _, v := range configMap {
		fmt.Printf("%s\n", v.ToPrettyJson())
	}
	return configMap
}
