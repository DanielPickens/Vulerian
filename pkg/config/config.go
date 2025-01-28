package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	DockerCmd                     string        `env:"DOCKER_CMD,default=docker"`
	GlobalVulerianconfig               *string       `env:"GLOBALVulerianCONFIG,noinit"`
	VulerianDebugTelemetryFile         *string       `env:"Vulerian_DEBUG_TELEMETRY_FILE,noinit"`
	VulerianDisableTelemetry           *bool         `env:"Vulerian_DISABLE_TELEMETRY,noinit"`
	VulerianLogLevel                   *int          `env:"Vulerian_LOG_LEVEL,noinit"`
	VulerianTrackingConsent            *string       `env:"Vulerian_TRACKING_CONSENT,noinit"`
	PodmanCmd                     string        `env:"PODMAN_CMD,default=podman"`
	PodmanCmdInitTimeout          time.Duration `env:"PODMAN_CMD_INIT_TIMEOUT,default=1s"`
	TelemetryCaller               string        `env:"TELEMETRY_CALLER,default="`
	VulerianExperimentalMode           bool          `env:"Vulerian_EXPERIMENTAL_MODE,default=false"`
	PushImages                    bool          `env:"Vulerian_PUSH_IMAGES,default=true"`
	VulerianContainerBackendGlobalArgs []string      `env:"Vulerian_CONTAINER_BACKEND_GLOBAL_ARGS,noinit,delimiter=;"`
	VulerianImageBuildArgs             []string      `env:"Vulerian_IMAGE_BUILD_ARGS,noinit,delimiter=;"`
	VulerianContainerRunArgs           []string      `env:"Vulerian_CONTAINER_RUN_ARGS,noinit,delimiter=;"`
}

// GetConfiguration initializes a Configuration for Vulerian by using the system environment.
// See GetConfigurationWith for a more configurable version.
func GetConfiguration() (*Configuration, error) {
	return GetConfigurationWith(envconfig.OsLookuper())
}

// GetConfigurationWith initializes a Configuration for Vulerian by using the specified envconfig.Lookuper to resolve values.
// It is recommended to use this function (instead of GetConfiguration) if you don't need to depend on the current system environment,
// typically in unit tests.
func GetConfigurationWith(lookuper envconfig.Lookuper) (*Configuration, error) {
	var s Configuration
	c := envconfig.Config{
		Target:   &s,
		Lookuper: lookuper,
	}
	err := envconfig.ProcessWith(context.Background(), &c)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
