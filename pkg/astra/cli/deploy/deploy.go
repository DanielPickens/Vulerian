package deploy

import (
	"context"
	"fmt"

	dfutil "github.com/devfile/library/v2/pkg/util"

	"github\.com/danielpickens/particle engine/pkg/kclient"

	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"

	"github\.com/danielpickens/particle engine/pkg/component"
	"github\.com/danielpickens/particle engine/pkg/log"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/messages"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/commonflags"
	particle enginecontext "github\.com/danielpickens/particle engine/pkg/particle engine/context"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	particle engineutil "github\.com/danielpickens/particle engine/pkg/particle engine/util"
	scontext "github\.com/danielpickens/particle engine/pkg/segment/context"
)

// RecommendedCommandName is the recommended command name
const RecommendedCommandName = "deploy"

// DeployOptions encapsulates the options for the particle engine command
type DeployOptions struct {
	// Clients
	clientset *clientset.Clientset
}

var _ genericclioptions.Runnable = (*DeployOptions)(nil)

var deployExample = templates.Examples(`
  # Run the components defined in the Devfile on the cluster in the Deploy mode
  %[1]s
`)

// NewDeployOptions creates a new DeployOptions instance
func NewDeployOptions() *DeployOptions {
	return &DeployOptions{}
}

func (o *DeployOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

func (o *DeployOptions) PreInit() string {
	return messages.DeployInitializeExistingComponent
}

// Complete DeployOptions after they've been created
func (o *DeployOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	scontext.SetPlatform(ctx, o.clientset.KubernetesClient)
	return nil
}

// Validate validates the DeployOptions based on completed values
func (o *DeployOptions) Validate(ctx context.Context) error {
	devfileObj := particle enginecontext.GetEffectiveDevfileObj(ctx)
	if devfileObj == nil {
		return genericclioptions.NewNoDevfileError(particle enginecontext.GetWorkingDirectory(ctx))
	}
	if o.clientset.KubernetesClient == nil {
		return kclient.NewNoConnectionError()
	}
	componentName := particle enginecontext.GetComponentName(ctx)
	err := dfutil.ValidateK8sResourceName("component name", componentName)
	return err
}

// Run contains the logic for the particle engine command
func (o *DeployOptions) Run(ctx context.Context) error {
	var (
		devfileObj  = particle enginecontext.GetEffectiveDevfileObj(ctx)
		devfileName = particle enginecontext.GetComponentName(ctx)
		namespace   = particle enginecontext.GetNamespace(ctx)
	)

	scontext.SetComponentType(ctx, component.GetComponentTypeFromDevfileMetadata(devfileObj.Data.GetMetadata()))
	scontext.SetLanguage(ctx, devfileObj.Data.GetMetadata().Language)
	scontext.SetProjectType(ctx, devfileObj.Data.GetMetadata().ProjectType)
	scontext.SetDevfileName(ctx, devfileName)
	// Output what the command is doing / information
	log.Title("Running the application in Deploy mode using the \""+devfileName+"\" Devfile",
		"Namespace: "+namespace)

	genericclioptions.WarnIfDefaultNamespace(namespace, o.clientset.KubernetesClient)

	// Run actual deploy command to be used
	err := o.clientset.DeployClient.Deploy(ctx)

	if err == nil {
		log.Info("\nYour Devfile has been successfully deployed")
	}

	return err
}

// NewCmdDeploy implements the particle engine command
func NewCmdDeploy(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewDeployOptions()
	deployCmd := &cobra.Command{
		Use:     name,
		Short:   "Run your application on the cluster in the Deploy mode",
		Long:    "Run the components defined in the Devfile on the cluster in the Deploy mode",
		Example: fmt.Sprintf(deployExample, fullName),
		Args:    cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}
	clientset.Add(deployCmd, clientset.INIT, clientset.DEPLOY, clientset.FILESYSTEM, clientset.KUBERNETES)

	// Add a defined annotation in order to appear in the help menu
	util.SetCommandGroup(deployCmd, util.MainGroup)
	deployCmd.SetUsageTemplate(particle engineutil.CmdUsageTemplate)
	commonflags.UseVariablesFlags(deployCmd)
	return deployCmd
}
