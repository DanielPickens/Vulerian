package list

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github\.com/danielpickens/particle engine/pkg/api"
	"github\.com/danielpickens/particle engine/pkg/component"
	"github\.com/danielpickens/particle engine/pkg/log"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/feature"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/list/binding"
	clicomponent "github\.com/danielpickens/particle engine/pkg/particle engine/cli/list/component"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/list/namespace"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/list/services"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/commonflags"
	fcontext "github\.com/danielpickens/particle engine/pkg/particle engine/commonflags/context"
	particle enginecontext "github\.com/danielpickens/particle engine/pkg/particle engine/context"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	particle engineutil "github\.com/danielpickens/particle engine/pkg/particle engine/util"

	ktemplates "k8s.io/kubectl/pkg/util/templates"
)

// RecommendedCommandName is the recommended list name
const RecommendedCommandName = "list"

var listExample = ktemplates.Examples(`  # List all components in the application
%[1]s
  `)

// ListOptions ...
type ListOptions struct {
	// Clients
	clientset *clientset.Clientset

	// Local variables
	namespaceFilter string

	// Flags
	namespaceFlag string
}

var _ genericclioptions.Runnable = (*ListOptions)(nil)
var _ genericclioptions.JsonOutputter = (*ListOptions)(nil)

// NewListOptions ...
func NewListOptions() *ListOptions {
	return &ListOptions{}
}

func (o *ListOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

// Complete ...
func (lo *ListOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	// If the namespace flag has been passed, we will search there.
	// if it hasn't, we will search from the default project / namespace.
	if lo.namespaceFlag != "" {
		lo.namespaceFilter = lo.namespaceFlag
	} else if lo.clientset.KubernetesClient != nil {
		lo.namespaceFilter = particle enginecontext.GetNamespace(ctx)
	}
	// Set the namespace; this ensures we fetch resources from the given namespace
	if lo.clientset.KubernetesClient != nil {
		lo.clientset.KubernetesClient.SetNamespace(lo.namespaceFilter)
	}

	return nil
}

// Validate ...
func (lo *ListOptions) Validate(ctx context.Context) (err error) {
	return nil
}

// Run has the logic to perform the required actions as part of command
func (lo *ListOptions) Run(ctx context.Context) error {
	listSpinner := log.Spinnerf("Listing resources from the namespace %q", lo.namespaceFilter)
	defer listSpinner.End(false)

	list, err := lo.run(ctx)
	if err != nil {
		return err
	}

	listSpinner.End(true)

	fmt.Printf("\nComponents:\n")
	clicomponent.HumanReadableOutput(ctx, list)
	fmt.Printf("\nBindings:\n")
	binding.HumanReadableOutput(list)
	return nil
}

// Run contains the logic for the particle engine command
func (lo *ListOptions) RunForJsonOutput(ctx context.Context) (out interface{}, err error) {
	return lo.run(ctx)
}

func (lo *ListOptions) run(ctx context.Context) (list api.ResourcesList, err error) {
	var (
		devfileObj    = particle enginecontext.GetEffectiveDevfileObj(ctx)
		componentName = particle enginecontext.GetComponentName(ctx)

		kubeClient   = lo.clientset.KubernetesClient
		podmanClient = lo.clientset.PodmanClient
	)

	switch fcontext.GetPlatform(ctx, "") {
	case commonflags.PlatformCluster:
		podmanClient = nil
	case commonflags.PlatformPodman:
		kubeClient = nil
	}

	allComponents, componentInDevfile, err := component.ListAllComponents(
		kubeClient, podmanClient, lo.namespaceFilter, devfileObj, componentName)
	if err != nil {
		return api.ResourcesList{}, err
	}

	var bindings []api.ServiceBinding
	var inDevfile []string

	workingDir := particle enginecontext.GetWorkingDirectory(ctx)
	bindings, inDevfile, err = lo.clientset.BindingClient.ListAllBindings(devfileObj, workingDir)
	if err != nil {
		return api.ResourcesList{}, err
	}

	// RunningOn is displayed only when Platform is active
	if !feature.IsEnabled(ctx, feature.GenericPlatformFlag) {
		for i := range allComponents {
			//lint:ignore SA1019 we need to output the deprecated value, before to remove it in a future release
			allComponents[i].RunningOn = ""
			allComponents[i].Platform = ""
		}
	}

	return api.ResourcesList{
		ComponentInDevfile: componentInDevfile,
		Components:         allComponents,
		BindingsInDevfile:  inDevfile,
		Bindings:           bindings,
	}, nil
}

// NewCmdList implements the list particle engine command
func NewCmdList(ctx context.Context, name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewListOptions()

	var listCmd = &cobra.Command{
		Use:     name,
		Short:   "List all components in the current namespace",
		Long:    "List all components in the current namespace.",
		Example: fmt.Sprintf(listExample, fullName),
		Args:    genericclioptions.NoArgsAndSilenceJSON,
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}
	clientset.Add(listCmd, clientset.KUBERNETES_NULLABLE, clientset.BINDING, clientset.FILESYSTEM)
	if feature.IsEnabled(ctx, feature.GenericPlatformFlag) {
		clientset.Add(listCmd, clientset.PODMAN_NULLABLE)
	}

	namespaceCmd := namespace.NewCmdNamespaceList(namespace.RecommendedCommandName, particle engineutil.GetFullName(fullName, namespace.RecommendedCommandName), testClientset)
	bindingCmd := binding.NewCmdBindingList(binding.RecommendedCommandName, particle engineutil.GetFullName(fullName, binding.RecommendedCommandName), testClientset)
	componentCmd := clicomponent.NewCmdComponentList(ctx, clicomponent.RecommendedCommandName, particle engineutil.GetFullName(fullName, clicomponent.RecommendedCommandName), testClientset)
	servicesCmd := services.NewCmdServicesList(services.RecommendedCommandName, particle engineutil.GetFullName(fullName, services.RecommendedCommandName), testClientset)
	listCmd.AddCommand(namespaceCmd, bindingCmd, componentCmd, servicesCmd)

	util.SetCommandGroup(listCmd, util.ManagementGroup)
	listCmd.SetUsageTemplate(particle engineutil.CmdUsageTemplate)
	listCmd.Flags().StringVar(&o.namespaceFlag, "namespace", "", "Namespace for particle engine to scan for components")

	commonflags.UseOutputFlag(listCmd)
	commonflags.UsePlatformFlag(listCmd)

	return listCmd
}
