package helper

import (
	"errors"
	"github.com/dispatchlabs/disgo/commons/types"
	"github.com/dispatchlabs/disgo/commons/utils"
	"github.com/dispatchlabs/disgo/sdk"
)

var delegates []types.Node
var nbrDelegates int

func GetDelegates(seedIp string) ([]types.Node, error) {
	if delegates == nil || len(delegates) == 0 {
		var err error
		delegates, err = sdk.GetDelegates(seedIp)
		nbrDelegates = len(delegates)
		if nbrDelegates == 0 {
			utils.Fatal(errors.New("No Delegates were returned by the seed"))
		}
		if err != nil {
			utils.Error(err)
		}
	}
	return delegates, nil
}

func GetRandomDelegate(seedIp string) types.Node {
	d, _ := GetDelegates(seedIp)
	rand := utils.Random(0, nbrDelegates)
	return d[rand]
}

