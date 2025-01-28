package commonflags

import (
	"context"
	"flag"
	"os"
	"testing"

	"github\.com/danielpickens/particle engine/pkg/config"
	envcontext "github\.com/danielpickens/particle engine/pkg/config/context"

	"github.com/spf13/pflag"
	"k8s.io/klog"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	cfg := config.Configuration{}
	ctx = envcontext.WithEnvConfig(ctx, cfg)
	klog.InitFlags(nil)
	AddOutputFlag()
	AddPlatformFlag(ctx)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	os.Exit(m.Run())
}
