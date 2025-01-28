package integration

import (
	"fmt"
	"path/filepath"

	"github\.com/danielpickens/particle engine/tests/helper"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("particle engine run command tests", Label(helper.LabelSkipOnOpenShift), func() {
	var cmpName string
	var commonVar helper.CommonVar

	// This is run before every Spec (It)
	var _ = BeforeEach(func() {
		commonVar = helper.CommonBeforeEach()
		cmpName = helper.RandString(6)
		_ = cmpName // Tparticle engine remove when used
		helper.Chdir(commonVar.Context)
		Expect(helper.VerifyFileExists(".particle engine/env/env.yaml")).To(BeFalse())
	})

	// This is run after every Spec (It)
	var _ = AfterEach(func() {
		helper.CommonAfterEach(commonVar)
	})

	When("directory is empty", Label(helper.LabelNoCluster), func() {
		BeforeEach(func() {
			Expect(helper.ListFilesInDir(commonVar.Context)).To(HaveLen(0))
		})

		It("should error", func() {
			output := helper.Cmd("particle engine", "run", "my-command").ShouldFail().Err()
			Expect(output).To(ContainSubstring("The current directory does not represent an particle engine component"))
		})
	})

	When("a component is bootstrapped", func() {
		BeforeEach(func() {
			helper.CopyExample(filepath.Join("source", "devfiles", "nodejs", "project"), commonVar.Context)
			helper.Cmd("particle engine", "init", "--name", cmpName, "--devfile-path", helper.GetExamplePath("source", "devfiles", "nodejs", "devfile-for-run.yaml")).ShouldPass()
			Expect(helper.VerifyFileExists(".particle engine/env/env.yaml")).To(BeFalse())
		})

		It("should fail if command is not found in devfile", Label(helper.LabelNoCluster), func() {
			output := helper.Cmd("particle engine", "run", "unknown-command").ShouldFail().Err()
			Expect(output).To(ContainSubstring(`no command named "unknown-command" found in the devfile`))

		})

		It("should fail if platform is not available", Label(helper.LabelNoCluster), func() {
			By("failing when trying to run on default platform", func() {
				output := helper.Cmd("particle engine", "run", "build").ShouldFail().Err()
				Expect(output).To(ContainSubstring(`unable to access the cluster`))

			})
			By("failing when trying to run on cluster", func() {
				output := helper.Cmd("particle engine", "run", "build", "--platform", "cluster").ShouldFail().Err()
				Expect(output).To(ContainSubstring(`unable to access the cluster`))

			})
			By("failing when trying to run on podman", func() {
				output := helper.Cmd("particle engine", "run", "build", "--platform", "podman").AddEnv("PODMAN_CMD=false").ShouldFail().Err()
				Expect(output).To(ContainSubstring(`unable to access podman`))
			})
		})

		It("should fail if particle engine dev is not running", func() {
			output := helper.Cmd("particle engine", "run", "build").ShouldFail().Err()
			Expect(output).To(ContainSubstring(`unable to get pod for component`))
			Expect(output).To(ContainSubstring(`Please check the command 'particle engine dev' is running`))
		})

		for _, podman := range []bool{false, true} {
			podman := podman
			for _, noCommands := range []bool{false, true} {
				noCommands := noCommands
				When(fmt.Sprintf("particle engine dev is executed with --no-commands=%v and ready", noCommands), helper.LabelPodmanIf(podman, func() {

					var devSession helper.DevSession

					BeforeEach(func() {
						var err error
						devSession, err = helper.StartDevMode(helper.DevSessionOpts{
							RunOnPodman: podman,
							NoCommands:  noCommands,
						})
						Expect(err).ToNot(HaveOccurred())
					})

					AfterEach(func() {
						devSession.Stop()
						devSession.WaitEnd()
					})

					It("should execute commands", func() {
						platform := "cluster"
						if podman {
							platform = "podman"
						}

						By("executing an exec command and displaying output", func() {
							output := helper.Cmd("particle engine", "run", "list-files", "--platform", platform).ShouldPass().Out()
							Expect(output).To(ContainSubstring("etc"))
						})

						By("executing an exec command in another container and displaying output", func() {
							output := helper.Cmd("particle engine", "run", "list-files-in-other-container", "--platform", platform).ShouldPass().Out()
							Expect(output).To(ContainSubstring("etc"))
						})

						if !podman {
							By("executing apply command on Kubernetes component", func() {
								output := helper.Cmd("particle engine", "run", "deploy-config", "--platform", platform).ShouldPass().Out()
								Expect(output).To(ContainSubstring("Creating resource ConfigMap/my-config"))
								out := commonVar.CliRunner.Run("get", "configmap", "my-config", "-n",
									commonVar.Project).Wait().Out.Contents()
								Expect(out).To(ContainSubstring("my-config"))
							})
						}

						if podman {
							By("executing apply command on Image component", func() {
								// Will fail because Dockerfile is not present, but we just want to check the build is started
								// We cannot use PODMAN_CMD=echo with --platform=podman
								output := helper.Cmd("particle engine", "run", "build-image", "--platform", platform).ShouldFail().Out()
								Expect(output).To(ContainSubstring("Building image locally"))
							})
						} else {
							By("executing apply command on Image component", func() {
								output := helper.Cmd("particle engine", "run", "build-image", "--platform", platform).AddEnv("PODMAN_CMD=echo").ShouldPass().Out()
								Expect(output).To(ContainSubstring("Building image locally"))
								Expect(output).To(ContainSubstring("Pushing image to container registry"))

							})
						}

						By("exiting with a status 1 when the exec command fails and displaying error output", func() {
							out := helper.Cmd("particle engine", "run", "error-cmd", "--platform", platform).ShouldFail().Err()
							Expect(out).To(ContainSubstring("No such file or directory"))
						})
					})
				}))
			}
		}
	})
})
