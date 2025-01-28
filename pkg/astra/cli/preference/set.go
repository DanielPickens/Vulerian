package preference

import (
	"context"
	"fmt"
	"strings"

	"github\.com/danielpickens/particle engine/pkg/log"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
	scontext "github\.com/danielpickens/particle engine/pkg/segment/context"

	"github\.com/danielpickens/particle engine/pkg/particle engine/cli/ui"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	"github\.com/danielpickens/particle engine/pkg/preference"

	"github.com/spf13/cobra"
	ktemplates "k8s.io/kubectl/pkg/util/templates"

	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
)

const setCommandName = "set"

var (
	setLongDesc = ktemplates.LongDesc(`Set an individual value in the particle engine preference file.  
%[1]s`)
	setExample = ktemplates.Examples(`
   # All available preference values you can set`)
)

// SetOptions encapsulates the options for the command
type SetOptions struct {
	// Clients
	clientset *clientset.Clientset

	// Flags
	forceFlag bool

	// Parameters
	paramName  string
	paramValue string
}

var _ genericclioptions.Runnable = (*SetOptions)(nil)

// NewSetOptions creates a new SetOptions instance
func NewSetOptions() *SetOptions {
	return &SetOptions{}
}

func (o *SetOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

func (o *SetOptions) UseDevfile(ctx context.Context, cmdline cmdline.Cmdline, args []string) bool {
	return false
}

// Complete completes SetOptions after they've been created
func (o *SetOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	o.paramName = strings.ToLower(args[0])
	o.paramValue = args[1]
	return
}

// Validate validates the SetOptions based on completed values
func (o *SetOptions) Validate(ctx context.Context) (err error) {
	return
}

// Run contains the logic for the command
func (o *SetOptions) Run(ctx context.Context) (err error) {

	if !o.forceFlag {
		if isSet := o.clientset.PreferenceClient.IsSet(o.paramName); isSet {
			// Tparticle engine: could add a logic to check if the new value set by the user is not same as the current value
			var proceed bool
			proceed, err = ui.Proceed(fmt.Sprintf("%v is already set. Do you want to override it in the config", o.paramName))
			if err != nil {
				return err
			}
			if !proceed {
				log.Info("Aborted by the user")
				return nil
			}
		}
	}

	err = o.clientset.PreferenceClient.SetConfiguration(o.paramName, o.paramValue)
	if err != nil {
		return err
	}

	log.Successf("Value of '%s' preference was set to '%s'", o.paramName, o.paramValue)

	scontext.SetPreferenceParameter(ctx, o.paramName, &o.paramValue)
	return nil
}

// NewCmdSet implements the config set particle engine command
func NewCmdSet(ctx context.Context, name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewSetOptions()
	preferenceSetCmd := &cobra.Command{
		Use:   name,
		Short: "Set a value in the particle engine preference file",
		Long:  fmt.Sprintf(setLongDesc, preference.FormatSupportedParameters()),
		Example: func(exampleString, fullName string) string {
			prefClient, err := preference.NewClient(ctx)
			if err != nil {
				util.LogErrorAndExit(err, "unable to set preference, something is wrong with particle engine, kindly raise an issue at https://github\.com/danielpickens/particle engine/issues/new?template=Bug.md")
			}
			properties := prefClient.NewPreferenceList()
			for _, property := range properties.Items {
				value := property.Default
				exampleString += fmt.Sprintf("\n  %s %s %v", fullName, property.Name, value)
			}
			return "\n" + exampleString
		}(setExample, fullName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}
	clientset.Add(preferenceSetCmd, clientset.PREFERENCE)

	preferenceSetCmd.Flags().BoolVarP(&o.forceFlag, "force", "f", false, "Don't ask for confirmation, set the preference directly")
	return preferenceSetCmd
}
