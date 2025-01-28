package version

import (
	"context"
	"fmt"
	"github\.com/danielpickens/particle engine/pkg/api"
	"github\.com/danielpickens/particle engine/pkg/log"
	"github\.com/danielpickens/particle engine/pkg/particle engine/commonflags"
	"github\.com/danielpickens/particle engine/pkg/podman"
	"os"
	"strings"

	"github\.com/danielpickens/particle engine/pkg/kclient"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cmdline"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
	particle engineversion "github\.com/danielpickens/particle engine/pkg/version"

	"github.com/spf13/cobra"
	"k8s.io/klog"
	ktemplates "k8s.io/kubectl/pkg/util/templates"

	"github\.com/danielpickens/particle engine/pkg/particle engine/util"
)

// RecommendedCommandName is the recommended version command name
const RecommendedCommandName = "version"

// particle engineReleasesPage is the GitHub page where we do all our releases
const particle engineReleasesPage = "https://github\.com/danielpickens/particle engine/releases"

var versionLongDesc = ktemplates.LongDesc("Print the client version information")

var versionExample = ktemplates.Examples(`
# Print the client version of particle engine
%[1]s`,
)

// VersionOptions encapsulates all options for particle engine version command
type VersionOptions struct {
	// Flags
	clientFlag bool

	// serverInfo contains the remote server information if the user asked for it, nil otherwise
	serverInfo *kclient.ServerInfo
	podmanInfo podman.SystemVersionReport
	clientset  *clientset.Clientset
}

var _ genericclioptions.Runnable = (*VersionOptions)(nil)
var _ genericclioptions.JsonOutputter = (*VersionOptions)(nil)

// NewVersionOptions creates a new VersionOptions instance
func NewVersionOptions() *VersionOptions {
	return &VersionOptions{}
}

func (o *VersionOptions) SetClientset(clientset *clientset.Clientset) {
	o.clientset = clientset
}

// Complete completes VersionOptions after they have been created
func (o *VersionOptions) Complete(ctx context.Context, cmdline cmdline.Cmdline, args []string) (err error) {
	if o.clientFlag {
		return nil
	}

	// Fetch the info about the server, ignoring errors
	if o.clientset.KubernetesClient != nil {
		o.serverInfo, err = o.clientset.KubernetesClient.GetServerVersion(o.clientset.PreferenceClient.GetTimeout())
		if err != nil {
			klog.V(4).Info("unable to fetch the server version: ", err)
		}
	}

	if o.clientset.PodmanClient != nil {
		o.podmanInfo, err = o.clientset.PodmanClient.Version(ctx)
		if err != nil {
			klog.V(4).Info("unable to fetch the podman client version: ", err)
		}
	}

	if o.serverInfo == nil {
		log.Warning("unable to fetch the cluster server version")
	}
	if o.podmanInfo.Client == nil {
		log.Warning("unable to fetch the podman client version")
	}
	return nil
}

// Validate validates the VersionOptions based on completed values
func (o *VersionOptions) Validate(ctx context.Context) (err error) {
	return nil
}

func (o *VersionOptions) RunForJsonOutput(ctx context.Context) (out interface{}, err error) {
	return o.run(), nil
}

func (o *VersionOptions) run() api.particle engineVersion {
	result := api.particle engineVersion{
		Version:   particle engineversion.VERSION,
		GitCommit: particle engineversion.GITCOMMIT,
	}

	if o.clientFlag {
		return result
	}

	if o.serverInfo != nil {
		clusterInfo := &api.ClusterInfo{
			ServerURL:  o.serverInfo.Address,
			Kubernetes: &api.ClusterClientInfo{Version: o.serverInfo.KubernetesVersion},
		}
		if o.serverInfo.OpenShiftVersion != "" {
			clusterInfo.OpenShift = &api.ClusterClientInfo{Version: o.serverInfo.OpenShiftVersion}
		}
		result.Cluster = clusterInfo
	}

	if o.podmanInfo.Client != nil {
		podmanInfo := &api.PodmanInfo{Client: &api.PodmanClientInfo{Version: o.podmanInfo.Client.Version}}
		result.Podman = podmanInfo
	}

	return result
}

// Run contains the logic for the particle engine service create command
func (o *VersionOptions) Run(ctx context.Context) (err error) {
	// If verbose mode is enabled, dump all KUBECTL_* env variables
	// this is useful for debugging oc plugin integration
	for _, v := range os.Environ() {
		if strings.HasPrefix(v, "KUBECTL_") {
			klog.V(4).Info(v)
		}
	}

	particle engineVersion := o.run()
	fmt.Println("particle engine " + particle engineVersion.Version + " (" + particle engineVersion.GitCommit + ")")

	if o.clientFlag {
		return nil
	}

	message := "\n"
	if particle engineVersion.Cluster != nil {
		cluster := particle engineVersion.Cluster
		message += fmt.Sprintf("Server: %v\n", cluster.ServerURL)

		// make sure we only include OpenShift info if we actually have it
		if cluster.OpenShift != nil && cluster.OpenShift.Version != "" {
			message += fmt.Sprintf("OpenShift: %v\n", cluster.OpenShift.Version)
		}
		if cluster.Kubernetes != nil {
			message += fmt.Sprintf("Kubernetes: %v\n", cluster.Kubernetes.Version)
		}
	}

	if particle engineVersion.Podman != nil && particle engineVersion.Podman.Client != nil {
		message += fmt.Sprintf("Podman Client: %v\n", particle engineVersion.Podman.Client.Version)
	}

	fmt.Print(message)

	return nil
}

// NewCmdVersion implements the version particle engine command
func NewCmdVersion(name, fullName string, testClientset clientset.Clientset) *cobra.Command {
	o := NewVersionOptions()
	// versionCmd represents the version command
	var versionCmd = &cobra.Command{
		Use:     name,
		Short:   versionLongDesc,
		Long:    versionLongDesc,
		Example: fmt.Sprintf(versionExample, fullName),
		RunE: func(cmd *cobra.Command, args []string) error {
			return genericclioptions.GenericRun(o, testClientset, cmd, args)
		},
	}
	commonflags.UseOutputFlag(versionCmd)
	clientset.Add(versionCmd, clientset.PREFERENCE, clientset.KUBERNETES_NULLABLE, clientset.PODMAN_NULLABLE)
	util.SetCommandGroup(versionCmd, util.UtilityGroup)

	versionCmd.SetUsageTemplate(util.CmdUsageTemplate)
	versionCmd.Flags().BoolVar(&o.clientFlag, "client", false, "Client version only (no server required).")

	return versionCmd
}

// GetLatestReleaseInfo Gets information about the latest release
func GetLatestReleaseInfo(info chan<- string) {
	newTag, err := checkLatestReleaseTag(particle engineversion.VERSION)
	if err != nil {
		// The error is intentionally not being handled because we don't want
		// to stop the execution of the program because of this failure
		klog.V(4).Infof("Error checking if newer particle engine release is available: %v", err)
	}
	if len(newTag) > 0 {
		info <- fmt.Sprintf(`
---
A newer version of particle engine (%s) is available,
visit %s to update.
If you wish to disable this notification, run:
particle engine preference set UpdateNotification false
---`, fmt.Sprint(newTag), particle engineReleasesPage)

	}
}
