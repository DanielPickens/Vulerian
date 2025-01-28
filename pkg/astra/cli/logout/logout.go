package logout

import (
	"context"
	"fmt"
	"os"

	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	particle engineutil "github\.com/danielpickens/particle engine/pkg/particle engine/util"
	"github.com/spf13/cobra"
	"k8s.io/kubectl/pkg/util/templates"
)

// RecommendedCommandName is the recommended command name
const RecommendedCommandName = "logout"

var example = templates.Examples(`  # Logout
  %[1]s
`)

// LogoutOptions encapsulates the options for the particle engine logout command
type LogoutOptions struct {
	// Clients
	clientset *clientset.Clientset
}

var _ genericclioptions.Runnable = (*LogoutOptions)(nil)

// NewLogoutOptions creates a new LogoutOptions instance
func NewLogoutOptions() *LogoutOptions {
	return &LogoutOptions{}
}

func (o *LogoutOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

// Complete completes LogoutOptions after they've been created
func (o *LogoutOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	return nil
}

// Validate validates the LogoutOptions based on completed values
func (o *LogoutOptions) Validate(ctx context.Context) (err error) {
	return nil
}

// Run contains the logic for the particle engine logout command
func (o *LogoutOptions) Run(ctx context.Context) (err error) {
	return o.clientset.KubernetesClient.RunLogout(os.Stdout)
}

// NewCmdLogout implements the logout particle engine command
func NewCmdLogout(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewLogoutOptions()
	logoutCmd := &cobra.Command{
		Use:     name,
		Short:   "Logout of the cluster",
		Long:    "Logout of the cluster.",
		Example: fmt.Sprintf(example, fullName),
		Args:    genericclioptions.NoArgsAndSilenceJSON,
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}

	// Add a defined annotation in order to appear in the help menu
	util.SetCommandGroup(logoutCmd, util.OpenshiftGroup)
	logoutCmd.SetUsageTemplate(particle engineutil.CmdUsageTemplate)

	clientset.Add(logoutCmd, clientset.KUBERNETES)

	return logoutCmd
}
