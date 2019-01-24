package helper

import (
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

func CreateSeedConfig(ipAddress string, httpPort, grpcPort int64, localPort int64, seedAccount *types.Account) *types.Config {
	seedConfig := createConfig(ipAddress, httpPort, grpcPort, localPort)

	seedConfig.Seeds = []*types.Node{
		{
			Address:      seedAccount.Address,
			GrpcEndpoint: seedConfig.GrpcEndpoint,
			HttpEndpoint: seedConfig.HttpEndpoint,
			LocalHttpApiPort:   localPort,
		},
	}
	return seedConfig
}

func CreateDelegateConfig(ipAddress string, httpPort, grpcPort int64, localPort int64, seedNodes []*types.Node) *types.Config {
	delegateConfig := createConfig(ipAddress, httpPort, grpcPort, localPort)
	delegateConfig.Seeds = seedNodes
	delegateConfig.IsBookkeeper = true
	return delegateConfig
}

func createConfig(ipAddress string, httpPort int64, grpcPort int64, localPort int64) *types.Config {

	//fileLocation := utils.GetConfigDir() + string(os.PathSeparator) + "myDisgoKey.json"
	//if GetConfig().KeyLocation != "" {
	//	fileLocation = GetConfig().KeyLocation
	//}

	configInstance := &types.Config{
		HttpEndpoint:       &types.Endpoint{Host: ipAddress, Port: httpPort},
		GrpcEndpoint:       &types.Endpoint{Host: ipAddress, Port: grpcPort},
		LocalHttpApiPort:   int(localPort),
		LocalHttpApiUsername: "Disgo",
		LocalHttpApiPassword: utils.RandomString(10),
		DelegateAddresses:  []string{},
		GrpcTimeout:        5,
	}
	return configInstance
}

func GetConfig(dirName, nodeName string) (*types.Config, error) {
	fileName := dirName + string(os.PathSeparator) + "config.json"
	if !utils.Exists(fileName) {
		return nil, errors.New("Config file does not exist")
	}
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}

	config, err := types.ToConfigFromJson(bytes)
	if err != nil {
		utils.Fatal("unable to read config.json", err)
	}
	return config, nil
}
