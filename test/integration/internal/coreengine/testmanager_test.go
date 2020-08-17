package coreengine_test

import (
	"testing"

	"citihub.com/probr/internal/coreengine"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Test manager integration tests, so actually calling out to Kube ...
// TODO: complete

func TestExecTest(t *testing.T) {
	// create a test and add it to the TestStore

	//test descriptor ... (general)
	grp := coreengine.CloudDriver
	cat := coreengine.General
	name := "account_manager"
	td := coreengine.TestDescriptor{Group: grp, Category: cat, Name: name}

	uuid := uuid.New().String()
	sat := coreengine.Pending

	test := coreengine.Test{
		UUID:           &uuid,
		TestDescriptor: &td,
		Status:         &sat,
	}

	assert.NotNil(t, test)

	// get the test mgr
	tm := coreengine.NewTestManager()

	assert.NotNil(t, tm)

	tsuuid := tm.AddTest(&test)

	s, err := tm.ExecTest(tsuuid)
	if err != nil {
		t.Fatalf("Error executing test: %v", err)
	}

	assert.True(t, s == 0)

}

func TestExecAllTests(t *testing.T) {

	tm := coreengine.NewTestManager()

	//add some tests and add them to the TM
	addTest(tm, "account_manager", coreengine.CloudDriver, coreengine.General)
	addTest(tm, "pod_security_policy", coreengine.Kubernetes, coreengine.PodSecurityPolicies)
	addTest(tm, "internet_access", coreengine.Kubernetes, coreengine.InternetAccess)

	tm.ExecAllTests()
}

func addTest(tm *coreengine.TestStore, testname string, grp coreengine.Group, cat coreengine.Category) {
	
	td := coreengine.TestDescriptor{Group: grp, Category: cat, Name: testname}

	uuid1 := uuid.New().String()
	sat := coreengine.Pending

	test := coreengine.Test{
		UUID:           &uuid1,
		TestDescriptor: &td,
		Status:         &sat,
	}

	//add - don't worry about the rtn uuid
	tm.AddTest(&test)

}