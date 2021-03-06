package services

import (
	"bytes"
	"fmt"
	"net"
	"testing"

	"github.com/ava-labs/avalanche-testing/avalanche/services/certs"
	"github.com/kurtosis-tech/kurtosis/commons/services"
	"github.com/stretchr/testify/assert"
)

var testPublicIP = net.ParseIP("172.17.0.2")

func TestNoDepsStartCommand(t *testing.T) {
	initializerCore := NewGeckoServiceInitializerCore(
		1,
		1,
		0,
		false,
		make(map[string]string),
		[]string{},
		certs.NewStaticGeckoCertProvider(bytes.Buffer{}, bytes.Buffer{}),
		INFO,
	)

	expected := []string{
		avalancheBinary,
		"--public-ip=" + testPublicIP.String(),
		"--network-id=local",
		"--http-port=9650",
		"--http-host=",
		"--staking-port=9651",
		"--log-level=info",
		"--snow-sample-size=1",
		"--snow-quorum-size=1",
		"--staking-enabled=false",
		"--tx-fee=0",
	}
	actual, err := initializerCore.GetStartCommand(make(map[string]string), testPublicIP, make([]services.Service, 0))
	assert.NoError(t, err, "An error occurred getting the start command")
	assert.Equal(t, expected, actual)
}

func TestWithDepsStartCommand(t *testing.T) {
	testNodeID := "node1"
	testDependencyIP := "1.2.3.4"

	bootstrapperNodeIDs := []string{
		testNodeID,
	}
	initializerCore := NewGeckoServiceInitializerCore(
		1,
		1,
		0,
		false,
		make(map[string]string),
		bootstrapperNodeIDs,
		certs.NewStaticGeckoCertProvider(bytes.Buffer{}, bytes.Buffer{}),
		INFO,
	)

	expected := []string{
		avalancheBinary,
		"--public-ip=" + testPublicIP.String(),
		"--network-id=local",
		"--http-port=9650",
		"--http-host=",
		"--staking-port=9651",
		"--log-level=info",
		"--snow-sample-size=1",
		"--snow-quorum-size=1",
		"--staking-enabled=false",
		"--tx-fee=0",
		fmt.Sprintf("--bootstrap-ips=%v:9651", testDependencyIP),
	}

	testDependency := GeckoService{
		ipAddr:      "1.2.3.4",
		jsonRPCPort: "9650/tcp",
		stakingPort: "9651/tcp",
	}
	testDependencySlice := []services.Service{
		testDependency,
	}
	actual, err := initializerCore.GetStartCommand(make(map[string]string), testPublicIP, testDependencySlice)
	assert.NoError(t, err, "An error occurred getting the start command")
	assert.Equal(t, expected, actual)
}
