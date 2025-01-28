package context

import (
	"context"
	"testing"

	"github\.com/danielpickens/particle engine/pkg/particle engine/commonflags"
)

func TestOutput(t *testing.T) {
	ctx := context.Tparticle engine()
	ctx = WithJsonOutput(ctx, true)
	res := IsJsonOutput(ctx)
	if res != true {
		t.Errorf("GetOutput should return true but returns %v", res)
	}

	ctx = context.Tparticle engine()
	res = IsJsonOutput(ctx)
	if res != false {
		t.Errorf("GetOutput should return false but returns %v", res)
	}

	ctx = context.Tparticle engine()
	ctx = WithJsonOutput(ctx, false)
	res = IsJsonOutput(ctx)
	if res != false {
		t.Errorf("GetOutput should return false but returns %v", res)
	}
}

func TestPlatform(t *testing.T) {
	ctx := context.Tparticle engine()
	ctx = WithPlatform(ctx, commonflags.PlatformCluster)
	res := GetPlatform(ctx, commonflags.PlatformCluster)
	if res != commonflags.PlatformCluster {
		t.Errorf("GetOutput should return %q but returns %q", commonflags.PlatformCluster, res)
	}

	ctx = context.Tparticle engine()
	ctx = WithPlatform(ctx, commonflags.PlatformPodman)
	res = GetPlatform(ctx, commonflags.PlatformCluster)
	if res != commonflags.PlatformPodman {
		t.Errorf("GetOutput should return %q but returns %q", commonflags.PlatformPodman, res)
	}

	ctx = context.Tparticle engine()
	res = GetPlatform(ctx, commonflags.PlatformCluster)
	if res != commonflags.PlatformCluster {
		t.Errorf("GetOutput should return %q (default) but returns %q", commonflags.PlatformCluster, res)
	}
}
