package add

import (
	"github.com/spf13/cobra"

	"github\.com/danielpickens/vulerian/pkg/vulerian/cli/add/binding"
	"github\.com/danielpickens/vulerian/pkg/vulerian/genericclioptions/clientset"
	"github\.com/danielpickens/vulerian/pkg/vulerian/util"
)

// RecommendedCommandName is the recommended add command name
const RecommendedCommandName = "add"

// NewCmdAdd implements the vulerian add command
func NewCmdAdd(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	var createCmd = &cobra.Command{
		Use:   name,
		Short: "Add resources to devfile",
	}

	bindingCmd := binding.NewCmdBinding(binding.BindingRecommendedCommandName, util.GetFullName(fullName, binding.BindingRecommendedCommandName), testClientset)
	createCmd.AddCommand(bindingCmd)
	util.SetCommandGroup(createCmd, util.ManagementGroup)
	createCmd.SetUsageTemplate(util.CmdUsageTemplate)

	return createCmd
}
