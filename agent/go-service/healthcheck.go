package main

import (
	"github.com/1204244136/MDA/agent/go-service/pkg/maafocus"
	maa "github.com/MaaXYZ/maa-framework-go/v4"
)

// agentHealthCheck verifies the Agent callback without interacting with the game.
type agentHealthCheck struct{}

func (*agentHealthCheck) Run(ctx *maa.Context, _ *maa.CustomActionArg) bool {
	maafocus.Print(ctx, "Go Agent connected / Go Agent 已连接")
	return true
}
