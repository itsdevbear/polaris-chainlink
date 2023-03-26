package tests

// environments_test shows some basic tests that deploy and teardown different testing environments
import (
	"os"
	"testing"

	"github.com/smartcontractkit/chainlink-integration-tests-template/helpers"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/smartcontractkit/helmenv/environment"
	"github.com/smartcontractkit/helmenv/tools"
	"github.com/smartcontractkit/integrations-framework/actions"
	"github.com/smartcontractkit/integrations-framework/client"
	"github.com/smartcontractkit/integrations-framework/config"
	"github.com/stretchr/testify/require"
)

func init() {
	// Sets logging output to use zerologger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// Loads the config files
	actions.LoadConfigs("../")
}

// TestEnvironment deploys a basic environment with 1 simulated Geth node, 1 mockserver, and 1 Chainlink node
func TestEnvironment(t *testing.T) {
	// Deploys a test environment with a single Chainlink node
	testEnvironment, err := environment.DeployOrLoadEnvironment(
		environment.NewChainlinkConfig(
			config.ChainlinkVals(),
			"basic-environment-test",
			config.GethNetworks()...,
		),
		tools.ChartsRoot,
	)
	require.NoError(t, err, "Error deploying test environment")

	// Connects to the deployed environment
	err = testEnvironment.ConnectAll()
	require.NoError(t, err, "Error connecting to test environment")

	// Fetches the connection to any deployed networks
	networkRegistry := client.NewDefaultNetworkRegistry()
	networks, err := networkRegistry.GetNetworks(testEnvironment)
	require.NoError(t, err, "Error retrieving networks from the test environment")
	defaultNetwork := networks.Default

	// Fetches the connection to any deployed Chainlink nodes
	chainlinkNodes, err := client.ConnectChainlinkNodes(testEnvironment)
	require.NoError(t, err, "Creating Chainlink node client shouldn't fail")

	// Sets up the test to elegantly teardown when it fails or passes
	helpers.PrepareTestTeardown(t, testEnvironment, networks, chainlinkNodes, nil)

	// Sets up a client to connect to the mockserver
	mockserver, err := client.ConnectMockServer(testEnvironment)
	require.NoError(t, err, "Creating mockserver clients shouldn't fail")

	require.NotNil(t, networks, "Expected networks to be populated")
	require.NotNil(t, defaultNetwork, "Expected default network to be populated")
	require.NotNil(t, chainlinkNodes, "Expected there to be Chainlink nodes")
	require.NotNil(t, mockserver, "Expected there to be a mockserver client")
	require.Equal(t, 1, len(chainlinkNodes), "Expected there to be 1 Chainlink node")
}

// TestMultiNodeEnvironment deploys a basic environment with 1 simulated Geth node, 1 mockserver, and 3 Chainlink nodes
func TestMultiNodeEnvironment(t *testing.T) {
	// Deploys a test environment with a single Chainlink node
	testEnvironment, err := environment.DeployOrLoadEnvironment(
		environment.NewChainlinkConfig(
			environment.ChainlinkReplicas(2, config.ChainlinkVals()), // Deploys 2 Chainlink nodes
			"multi-node-test",
			config.GethNetworks()...,
		),
		tools.ChartsRoot,
	)
	require.NoError(t, err, "Error deploying test environment")

	// Connects to the deployed environment
	err = testEnvironment.ConnectAll()
	require.NoError(t, err, "Error connecting to test environment")

	// Fetches the connection to any deployed networks
	networkRegistry := client.NewDefaultNetworkRegistry()
	networks, err := networkRegistry.GetNetworks(testEnvironment)
	require.NoError(t, err, "Error retrieving networks from the test environment")
	defaultNetwork := networks.Default

	// Fetches the connection to any deployed Chainlink nodes
	chainlinkNodes, err := client.ConnectChainlinkNodes(testEnvironment)
	require.NoError(t, err, "Creating Chainlink node client shouldn't fail")

	// Sets up the test to elegantly teardown when it fails or passes
	helpers.PrepareTestTeardown(t, testEnvironment, networks, chainlinkNodes, nil)

	// Sets up a client to connect to the mockserver
	mockserver, err := client.ConnectMockServer(testEnvironment)
	require.NoError(t, err, "Creating mockserver clients shouldn't fail")

	require.NotNil(t, networks, "Expected networks to be populated")
	require.NotNil(t, defaultNetwork, "Expected default network to be populated")
	require.NotNil(t, chainlinkNodes, "Expected there to be Chainlink nodes")
	require.NotNil(t, mockserver, "Expected there to be a mockserver client")
	require.Equal(t, 2, len(chainlinkNodes), "Expected there to be 2 Chainlink nodes")
}
