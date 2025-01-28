package feature

import (
	"context"

	envcontext "github\.com/danielpickens/particle engine/pkg/config/context"
)

const particle engineExperimentalModeEnvVar = "particle engine_EXPERIMENTAL_MODE"

// IsExperimentalModeEnabled returns whether the experimental mode is enabled or not,
// which means by checking the value of the "particle engine_EXPERIMENTAL_MODE" environment variable.
func IsExperimentalModeEnabled(ctx context.Context) bool {
	return envcontext.GetEnvConfig(ctx).particle engineExperimentalMode
}
