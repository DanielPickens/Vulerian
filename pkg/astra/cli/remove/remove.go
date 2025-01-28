package remove

import (
	"github.com/spf13/cobra"

	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/remove/binding"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
)

// RecommendedCommandName is the recommended remove command name
const RecommendedCommandName = "remove"

// NewCmdRemove implements the particle engine remove command
func NewCmdRemove(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	var removeCmd = &cobra.Command{
		Use:   name,
		Short: "Remove resources from devfile",
	}

	bindingCmd := binding.NewCmdBinding(binding.BindingRecommendedCommandName, util.GetFullName(fullName, binding.BindingRecommendedCommandName), testClientset)
	removeCmd.AddCommand(bindingCmd)
	util.SetCommandGroup(removeCmd, util.ManagementGroup)
	removeCmd.SetUsageTemplate(util.CmdUsageTemplate)

	return removeCmd
}
