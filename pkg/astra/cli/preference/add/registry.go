package add

import (
	"context"
	// Built-in packages
	"fmt"

	// Third-party packages
	dfutil "github.com/devfile/library/v2/pkg/util"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
	ktemplates "k8s.io/kubectl/pkg/util/templates"

	// particle engine packages
	"github\.com/danielpickens/particle engine/pkg/log"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/registry"
	"github\.com/danielpickens/particle engine/pkg/util"
)

const registryCommandName = "registry"

// "particle engine preference add registry" command description and examples
var (
	addLongDesc = ktemplates.LongDesc(`Add devfile registry`)

	addExample = ktemplates.Examples(`# Add devfile registry
	%[1]s CheRegistry https://che-devfile-registry.openshift.io
	`)
)

// RegistryOptions encapsulates the options for the "particle engine preference add registry" command
type RegistryOptions struct {
	// Clients
	clientset *clientset.Clientset

	// Parameters
	registryName string
	registryURL  string

	// Flags
	tokenFlag string

	operation string
	user      string
}

var _ genericclioptions.Runnable = (*RegistryOptions)(nil)

// NewRegistryOptions creates a new RegistryOptions instance
func NewRegistryOptions() *RegistryOptions {
	return &RegistryOptions{}
}

func (o *RegistryOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

func (o *RegistryOptions) UseDevfile(ctx context.Context, cmdline cmdline.Cmdline, args []string) bool {
	return false
}

// Complete completes RegistryOptions after they've been created
func (o *RegistryOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	o.operation = "add"
	o.registryName = args[0]
	o.registryURL = args[1]
	o.user = "default"
	return nil
}

// Validate validates the RegistryOptions based on completed values
func (o *RegistryOptions) Validate(ctx context.Context) (err error) {
	err = util.ValidateURL(o.registryURL)
	if err != nil {
		return err
	}
	isGithubRegistry, err := registry.IsGithubBasedRegistry(o.registryURL)
	if err != nil {
		return err
	}
	if isGithubRegistry {
		return &registry.ErrGithubRegistryNotSupported{}
	}
	return nil
}

// Run contains the logic for "particle engine preference add registry" command
func (o *RegistryOptions) Run(ctx context.Context) (err error) {
	isSecure := false
	if o.tokenFlag != "" {
		isSecure = true
	}

	err = o.clientset.PreferenceClient.RegistryHandler(o.operation, o.registryName, o.registryURL, false, isSecure)
	if err != nil {
		return err
	}

	if o.tokenFlag != "" {
		err = keyring.Set(dfutil.CredentialPrefix+o.registryName, o.user, o.tokenFlag)
		if err != nil {
			return fmt.Errorf("unable to store registry credential to keyring: %w", err)
		}
	}

	log.Info("New registry successfully added")
	return nil
}

// NewCmdRegistry implements the "particle engine preference add registry" command
func NewCmdRegistry(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewRegistryOptions()
	registryCmd := &cobra.Command{
		Use:     fmt.Sprintf("%s <registry name> <registry URL>", name),
		Short:   addLongDesc,
		Long:    addLongDesc,
		Example: fmt.Sprintf(fmt.Sprint(addExample), fullName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}
	clientset.Add(registryCmd, clientset.PREFERENCE)

	registryCmd.Flags().StringVar(&o.tokenFlag, "token", "", "Token to be used to access secure registry")

	return registryCmd
}
