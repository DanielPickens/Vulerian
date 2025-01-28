//go:build !race
// +build !race

package integration

import (
	"testing"

	"github\.com/danielpickens/particle engine/tests/helper"
)

func TestIntegration(t *testing.T) {
	helper.RunTestSpecs(t, "Integration Suite")
}
