package config

import (
	"github.com/dispatchlabs/dev-tools/common-util/helper"
	"testing"
)

func TestUpdateDisgoExecutable(t *testing.T) {
	RefreshDisgoExecutable(helper.GetDefaultDirectory())
}

func TestDeleteDB(t *testing.T) {
	ClearDB(helper.GetDefaultDirectory())
}

func TestBuildRuntimeCluster(t *testing.T) {
	CleanAndBuildNewConfig(1, 4, nil, true)
}
