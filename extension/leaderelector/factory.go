package leaderelector

import (
	"context"
	"errors"
	"fmt"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"time"
)

// CreateDefaultConfig returns the default configuration for the extension.
func CreateDefaultConfig() component.Config {
	return &Config{
		LeaseDuration: 15 * time.Second,
		RenewDuration: 10 * time.Second,
		RetryPeriod:   2 * time.Second,
		// Set default values for your configuration
	}
}

// CreateExtension creates the extension instance based on the configuration.
func CreateExtension(
	ctx context.Context,
	set extension.Settings,
	cfg component.Config,
) (extension.Extension, error) {
	baseCfg, ok := cfg.(*Config)
	if !ok {
		return nil, errors.New("Invalid config, cannot create extension leaderelector")
	}
	fmt.Printf("Creating leaderelector extension with config: %+v\n", baseCfg)
	return &leaderElectionExtension{
		config: baseCfg,
		logger: set.Logger,
	}, nil
}

// NewFactory creates a new factory for your extension.
func NewFactory() extension.Factory {
	return extension.NewFactory(
		component.MustNewType("leaderelector"),
		CreateDefaultConfig,
		CreateExtension,
		component.StabilityLevelDevelopment,
	)
}
