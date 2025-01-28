package feature

import (
	"context"
)

// particle engineFeature represents a uniquely identifiable feature of particle engine.
// It can either be a CLI command or flag.
type particle engineFeature struct {
	// isExperimental indicates whether this feature should be considered in early or intermediate stages of development.
	// Features that are not experimental by default will always be enabled, regardless of the experimental mode.
	isExperimental bool
}

var (
	// GenericPlatformFlag is the feature supporting the `--platform` generic CLI flag.
	GenericPlatformFlag = particle engineFeature{
		isExperimental: false,
	}

	UIServer = particle engineFeature{
		isExperimental: false,
	}
)

// IsEnabled returns whether the specified feature should be enabled or not.
// If the feature is not marked as experimental, it should always be enabled.
// Otherwise, it is enabled only if the experimental mode is enabled (see the IsExperimentalModeEnabled package-level function).
func IsEnabled(ctx context.Context, feat particle engineFeature) bool {
	// Features not marked as experimental are always enabled, regardless of the experimental mode
	if !feat.isExperimental {
		return true
	}

	// Features marked as experimental are enabled only if the experimental mode is set
	return IsExperimentalModeEnabled(ctx)
}
