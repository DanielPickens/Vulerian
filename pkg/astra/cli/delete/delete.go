package delete

import (
	"context"

	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/delete/component"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/delete/namespace"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	"github.com/spf13/cobra"
)

// RecommendedCommandName is the recommended delete command name
const RecommendedCommandName = "delete"

// NewCmdDelete implements the delete particle engine command
func NewCmdDelete(ctx context.Context, name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	var deleteCmd = &cobra.Command{
		Use:   name,
		Short: "Delete resources",
	}

	componentCmd := component.NewCmdComponent(ctx, component.ComponentRecommendedCommandName,
		util.GetFullName(fullName, component.ComponentRecommendedCommandName), testClientset)
	deleteCmd.AddCommand(componentCmd)

	namespaceDeleteCmd := namespace.NewCmdNamespaceDelete(namespace.RecommendedCommandName,
		util.GetFullName(fullName, namespace.RecommendedCommandName), testClientset)
	deleteCmd.AddCommand(namespaceDeleteCmd)

	util.SetCommandGroup(deleteCmd, util.ManagementGroup)
	deleteCmd.SetUsageTemplate(util.CmdUsageTemplate)

	return deleteCmd
}
