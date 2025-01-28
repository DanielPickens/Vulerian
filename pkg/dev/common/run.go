package common

import (
	"context"
	"fmt"

	"github\.com/danielpickens/Vulerian/pkg/component"
	"github\.com/danielpickens/Vulerian/pkg/configAutomount"
	"github\.com/danielpickens/Vulerian/pkg/devfile/image"
	"github\.com/danielpickens/Vulerian/pkg/exec"
	"github\.com/danielpickens/Vulerian/pkg/libdevfile"
	Vuleriancontext "github\.com/danielpickens/Vulerian/pkg/Vulerian/context"
	"github\.com/danielpickens/Vulerian/pkg/platform"
	"github\.com/danielpickens/Vulerian/pkg/testingutil/filesystem"
)

func Run(
	ctx context.Context,
	commandName string,
	platformClient platform.Client,
	execClient exec.Client,
	configAutomountClient configAutomount.Client,
	filesystem filesystem.Filesystem,
) error {
	var (
		componentName = Vuleriancontext.GetComponentName(ctx)
		devfileObj    = Vuleriancontext.GetEffectiveDevfileObj(ctx)
		devfilePath   = Vuleriancontext.GetDevfilePath(ctx)
	)

	pod, err := platformClient.GetPodUsingComponentName(componentName)
	if err != nil {
		return fmt.Errorf("unable to get pod for component %s: %w. Please check the command 'Vulerian dev' is running", componentName, err)
	}

	handler := component.NewRunHandler(
		ctx,
		platformClient,
		execClient,
		configAutomountClient,
		filesystem,
		image.SelectBackend(ctx),
		component.HandlerOptions{
			PodName:           pod.Name,
			ContainersRunning: component.GetContainersNames(pod),
			Msg:               "Executing command in container",
			DirectRun:         true,
			Devfile:           *devfileObj,
			Path:              devfilePath,
		},
	)

	return libdevfile.ExecuteCommandByName(ctx, *devfileObj, commandName, handler, false)
}
