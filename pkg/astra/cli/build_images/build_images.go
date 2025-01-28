package build_images

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"

	"github\.com/danielpickens/particle engine/pkg/devfile/image"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/commonflags"
	particle enginecontext "github\.com/danielpickens/particle engine/pkg/particle engine/context"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	particle engineutil "github\.com/danielpickens/particle engine/pkg/particle engine/util"
)

// RecommendedCommandName is the recommended command name
const RecommendedCommandName = "build-images"

// BuildImagesOptions encapsulates the options for the particle engine command
type BuildImagesOptions struct {
	// Clients
	clientset *clientset.Clientset

	// Flags
	pushFlag bool
}

var _ genericclioptions.Runnable = (*BuildImagesOptions)(nil)

var buildImagesExample = templates.Examples(`
  # Build images defined in the devfile
  %[1]s

  # Build images and push them to their registries
  %[1]s --push
`)

// NewBuildImagesOptions creates a new BuildImagesOptions instance
func NewBuildImagesOptions() *BuildImagesOptions {
	return &BuildImagesOptions{}
}

func (o *BuildImagesOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

// Complete completes LoginOptions after they've been created
func (o *BuildImagesOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	return nil
}

// Validate validates the LoginOptions based on completed values
func (o *BuildImagesOptions) Validate(ctx context.Context) (err error) {
	devfileObj := particle enginecontext.GetEffectiveDevfileObj(ctx)
	if devfileObj == nil {
		return genericclioptions.NewNoDevfileError(particle enginecontext.GetWorkingDirectory(ctx))
	}
	return nil
}

// Run contains the logic for the particle engine command
func (o *BuildImagesOptions) Run(ctx context.Context) (err error) {
	return image.BuildPushImages(ctx, image.SelectBackend(ctx), o.clientset.FS, o.pushFlag)
}

// NewCmdBuildImages implements the particle engine command
func NewCmdBuildImages(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewBuildImagesOptions()
	buildImagesCmd := &cobra.Command{
		Use:     name,
		Short:   "Build images",
		Long:    "Build images defined in the devfile",
		Example: fmt.Sprintf(buildImagesExample, fullName),
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}

	util.SetCommandGroup(buildImagesCmd, util.MainGroup)
	buildImagesCmd.SetUsageTemplate(particle engineutil.CmdUsageTemplate)
	commonflags.UseVariablesFlags(buildImagesCmd)
	buildImagesCmd.Flags().BoolVar(&o.pushFlag, "push", false, "If true, build and push the images")
	clientset.Add(buildImagesCmd, clientset.FILESYSTEM)

	return buildImagesCmd
}
