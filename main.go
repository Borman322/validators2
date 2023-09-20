package main

import (
	"context"
	"validators2/src/common"
	"validators2/src/config"
	"validators2/src/container"
	"validators2/src/service"
)

func main() {
	di := container.CreateContainer()

	ctx := context.Background()

	container.MustInvoke(di, func(
		config *config.Config,
		ValMonitoring *service.ValidatorService,

	) {

		ValMonitoring.Start(ctx)

		common.WaitForSignal()
	})
}
