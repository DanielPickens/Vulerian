package main

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github\.com/danielpickens/particle engine/pkg/config"
	envcontext "github\.com/danielpickens/particle engine/pkg/config/context"
	"github\.com/danielpickens/particle engine/pkg/particle engine/cli"
	"github\.com/danielpickens/particle engine/pkg/particle engine/genericclioptions/clientset"
)

var (
	intro = `Usage:
  particle engine [flags]
  particle engine [command]

Examples:
  Initializing your component by taking your pick from multiple languages or frameworks:
  particle engine init
  
  After creating your initial component, start development with:
  particle engine dev
  
  Want to deploy after development? See it live with:
  particle engine deploy`

	mainCommands = `Main Commands:
  build-images Build images
  deploy       Run your application on the cluster in the Deploy mode
  dev          Run your application on the cluster in the Dev mode
  init         Init bootstraps a new project
  logs         Show logs of all containers of the component
  registry     List all components from the Devfile registry
  run          Run a specific command in the Dev mode

`

	managementCommands = `Management Commands:
  add          Add resources to devfile (binding)
  create       Perform create operation (namespace)
  delete       Delete resources (component, namespace)
  describe     Describe resource (binding, component)
  list         List all components in the current namespace (binding, component, namespace, services)
  remove       Remove resources from devfile (binding)
  set          Perform set operation (namespace)

`

	openshiftCommands = `OpenShift Commands:
  login        Login to cluster 
  logout       Logout of the cluster`

	utilityCommands = `Utility Commands:
  analyze      Detect devfile to use based on files present in current directory
  completion   Add particle engine completion support to your development environment
  preference   Modifies preference settings (add, remove, set, unset, view)
  version      Print the client version information

`
)

func Testparticle engineHelp(t *testing.T) {
	ctx := context.Background()
	envConfig, err := config.GetConfiguration()
	if err != nil {
		t.Fatal(err)
	}
	ctx = envcontext.WithEnvConfig(ctx, *envConfig)

	resetGlobalFlags()

	root, err := cli.NewCmdparticle engine(ctx, cli.particle engineRecommendedName, cli.particle engineRecommendedName, nil, clientset.Clientset{})
	if err != nil {
		t.Fatal(err)
	}

	var stdoutB, stderrB bytes.Buffer
	root.SetOut(&stdoutB)
	root.SetErr(&stderrB)

	root.SetArgs([]string{"help"})

	err = root.ExecuteContext(ctx)
	if err != nil {
		t.Fatal(err)
	}

	stdout := stdoutB.String()
	stderr := stderrB.String()

	if stderr != "" {
		t.Fatal("stderr should be empty")
	}

	if !strings.Contains(stdout, intro) {
		t.Fatalf("stdout should contain \n%s\nbut is\n%s\n", intro, stdout)
	}
	if !strings.Contains(stdout, mainCommands) {
		t.Fatalf("stdout should contain \n%s\nbut is\n%s\n", mainCommands, stdout)
	}
	if !strings.Contains(stdout, managementCommands) {
		t.Fatalf("stdout should contain \n%s\nbut is\n%s\n", managementCommands, stdout)
	}
	if !strings.Contains(stdout, openshiftCommands) {
		t.Fatalf("stdout should contain \n%s\nbut is\n%s\n", openshiftCommands, stdout)
	}
	if !strings.Contains(stdout, utilityCommands) {
		t.Fatalf("stdout should contain \n%s\nbut is\n%s\n", utilityCommands, stdout)
	}
}
