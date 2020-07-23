package ava_services

import (
	"time"

	"github.com/kurtosis-tech/ava-e2e-tests/gecko_client"
	"github.com/kurtosis-tech/kurtosis/commons/services"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
)

// Implements ServiceAvailabilityCheckerCore
type GeckoServiceAvailabilityCheckerCore struct{}

// Set 20 second initial delay to allow the initial health check
// func (g GeckoServiceAvailabilityCheckerCore) InitialDelay() time.Duration { return 20 * time.Second }

func (g GeckoServiceAvailabilityCheckerCore) IsServiceUp(toCheck services.Service, dependencies []services.Service) bool {
	castedService := toCheck.(GeckoService)
	jsonRpcSocket := castedService.GetJsonRpcSocket()
	client := gecko_client.NewGeckoClient(jsonRpcSocket.GetIpAddr(), jsonRpcSocket.GetPort())
	healthInfo, err := client.HealthApi().GetLiveness()
	if err != nil {
		// TODO increase log level when InitialDelay is set to a conservative value
		logrus.Info(stacktrace.Propagate(err, "Error occurred in getting liveness info"))
		return false
	}

	return healthInfo.Healthy
}

func (g GeckoServiceAvailabilityCheckerCore) GetTimeout() time.Duration {
	return 90 * time.Second
}
