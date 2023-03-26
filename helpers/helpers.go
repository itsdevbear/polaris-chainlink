// Package helpers provides some common test methods and operations to keep actual test code clean
package helpers

import (
	"testing"

	"github.com/smartcontractkit/helmenv/environment"
	"github.com/smartcontractkit/integrations-framework/actions"
	"github.com/smartcontractkit/integrations-framework/client"
	"github.com/smartcontractkit/integrations-framework/testreporters"
	"github.com/smartcontractkit/integrations-framework/utils"
	"github.com/test-go/testify/require"
)

// PrepareTestTeardown prepares the test to be torn down on error or failure. Make sure to call this method as soon
// as the test is ready to be run.
func PrepareTestTeardown(
	t *testing.T,
	testEnvironment *environment.Environment,
	networks *client.Networks,
	chainlinkNodes []client.Chainlink,
	optionalTestReporter testreporters.TestReporter,
) {
	t.Helper()
	t.Cleanup(func() {
		err := actions.TeardownSuite(testEnvironment, networks, utils.ProjectRoot, chainlinkNodes, optionalTestReporter)
		require.NoError(t, err, "Error tearing down test environment")
	})
}
