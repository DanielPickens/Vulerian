package helper

import (
	"os"

	_ "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/feature"
)

// EnableExperimentalMode enables the experimental mode, so that experimental features of particle engine can be used.
func EnableExperimentalMode() {
	err := os.Setenv(feature.particle engineExperimentalModeEnvVar, "true")
	Expect(err).ShouldNot(HaveOccurred())
}

// ResetExperimentalMode disables the experimental mode.
//
// Note that calling any experimental feature of particle engine right is expected to error out if experimental mode is not enabled.
func ResetExperimentalMode() {
	if _, ok := os.LookupEnv(feature.particle engineExperimentalModeEnvVar); ok {
		err := os.Unsetenv(feature.particle engineExperimentalModeEnvVar)
		Expect(err).ShouldNot(HaveOccurred())
	}
}
