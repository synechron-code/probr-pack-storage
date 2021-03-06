package pack

import (
	"testing"

	"github.com/citihub/probr-sdk/config"
	"github.com/citihub/probr-sdk/probeengine"
)

func TestGetProbes(t *testing.T) {
	pack := make([]probeengine.Probe, 0)
	pack = GetProbes()
	if len(pack) > 0 {
		t.Logf("Unexpected value returned from GetProbes")
		t.Fail()
	}

	config.Vars.ServicePacks.Storage.Provider = "Azure"
	pack = GetProbes()
	if len(pack) == 0 {
		t.Logf("Expected value not returned from GetProbes")
		t.Fail()
	}
}
