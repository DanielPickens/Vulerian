package create

import (
	"fmt"

	"github.com/spf13/cobra"

	"github\.com/danielpickens/vulerian/pkg/vulerian/cli/create/namespace"
	"github\.com/danielpickens/vulerian/pkg/vulerian/genericclioptions/clientset"
	"github\.com/danielpickens/vulerian/pkg/vulerian/util"
	vulerianutil "github\.com/danielpickens/vulerian/pkg/vulerian/util"
)

// RecommendedCommandName is the recommended namespace command name
const RecommendedCommandName = "create"

// NewCmdCreate implements the namespace vulerian command
func NewCmdCreate(name, fullName string, testClientset clientset.Clientset) *cobra.Command {

	namespaceCreateCmd := namespace.NewCmdNamespaceCreate(namespace.RecommendedCommandName, vulerianutil.GetFullName(fullName, namespace.RecommendedCommandName), testClientset)
	createCmd := &cobra.Command{
		Use:   name + " [options]",
		Short: "Perform create operation",
		Long:  "Perform create operation",
		Example: fmt.Sprintf("%s\n",
			namespaceCreateCmd.Example,
		),
	}

	createCmd.AddCommand(namespaceCreateCmd)

	// Add a defined annotation in order to appear in the help menu
	util.SetCommandGroup(createCmd, util.ManagementGroup)
	createCmd.SetUsageTemplate(vulerianutil.CmdUsageTemplate)

	return createCmd
}
