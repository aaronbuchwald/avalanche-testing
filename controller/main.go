package main

import (
	"flag"
	"fmt"
	"github.com/kurtosis-tech/ava-e2e-tests/commons/ava_testsuite"
	"github.com/kurtosis-tech/ava-e2e-tests/commons/logging"
	"github.com/kurtosis-tech/kurtosis/controller"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	testVolumeArg := flag.String(
		"test-volume",
		"",
		"The name of the volume that will have been created by the initializer and mounted on this controller, which can also be mounted on other nodes in the network",
	)

	testVolumeMountpointArg := flag.String(
		"test-volume-mountpoint",
		"",
		"The filepath in the test controller's filesystem where the test volume will have been mounted by the initializer",
	)

	testNameArg := flag.String(
		"test",
		"",
		"Comma-separated list of specific tests to run (leave empty or omit to run all tests)",
	)

	testImageNameArg := flag.String(
		"test-image-name",
		"",
		"Name of Docker image of the service being tested",
	)

	dockerNetworkArg := flag.String(
		"docker-network",
		"",
		"Name of Docker network that the container is running in, and in which all services should be started",
	)

	subnetMaskArg := flag.String(
		"subnet-mask",
		"",
		"Subnet mask of the Docker network that the test controller is running in",
	)

	testControllerIpArg := flag.String(
		"test-controller-ip",
		"",
		"IP address of the Docker container running this test controller",
	)

	gatewayIpArg := flag.String(
		"gateway-ip",
		"",
		"IP address of the gateway address on the Docker network that the test controller is running in",
	)

	logLevelArg := flag.String(
		"log-level",
		"info",
		fmt.Sprintf("Log level to use for the controller (%v)", logging.GetAcceptableStrings()),
	)
	flag.Parse()


	logLevelPtr := logging.LevelFromString(*logLevelArg)
	if logLevelPtr == nil {
		// It's a little goofy that we're logging an error before we've set the loglevel, but we do so at the highest
		//  level so that whatever the default the user should see it
		logrus.Fatalf("Invalid initializer log level %v", *logLevelArg)
		os.Exit(1)
	}
	logrus.SetLevel(*logLevelPtr)

	logrus.Debugf(
		"Controller CLI arguments: dockerNetwork: %v, subnetMask %v, gatewayIp %v, testControllerIp %v, testImageName %v",
		*dockerNetworkArg,
		*subnetMaskArg,
		*gatewayIpArg,
		*testControllerIpArg,
		*testImageNameArg)

	controller := controller.NewTestController(
		*testVolumeArg,
		*testVolumeMountpointArg,
		*dockerNetworkArg,
		*subnetMaskArg,
		*gatewayIpArg,
		*testControllerIpArg,
		ava_testsuite.AvaTestSuite{},
		*testImageNameArg)

	logrus.Infof("Running test '%v'...", *testNameArg)
	setupErr, testErr := controller.RunTest(*testNameArg)
	if setupErr != nil {
		logrus.Errorf("Test %v encountered an error during setup (test did not run):", *testNameArg)
		logrus.Error(setupErr)
		os.Exit(1)
	}
	if testErr != nil {
		logrus.Errorf("Test %v failed:", *testNameArg)
		logrus.Error(testErr)
		os.Exit(1)
	}
	logrus.Infof("Test %v succeeded", *testNameArg)
}
