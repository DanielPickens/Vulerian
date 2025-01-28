package integration

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github\.com/danielpickens/particle engine/tests/helper"
)

var _ = Describe("particle engine generic", Label(helper.LabelSkipOnOpenShift), func() {
	// Tparticle engine: A neater way to provide particle engine path. Currently we assume \
	// particle engine and oc in $PATH already
	var oc helper.OcRunner
	var commonVar helper.CommonVar

	// This is run before every Spec (It)
	var _ = BeforeEach(func() {
		oc = helper.NewOcRunner("oc")
		commonVar = helper.CommonBeforeEach()
	})

	// Clean up after the test
	// This is run after every Spec (It)
	var _ = AfterEach(func() {
		helper.CommonAfterEach(commonVar)
	})

	for _, label := range []string{
		helper.LabelNoCluster, helper.LabelUnauth,
	} {
		label := label
		Context("label "+label, Label(label), func() {
			When("running particle engine --help", func() {
				var output string
				BeforeEach(func() {
					output = helper.Cmd("particle engine", "--help").ShouldPass().Out()
				})
				It("returns full help contents including usage, examples, commands, utility commands, component shortcuts, and flags sections", func() {
					helper.MatchAllInOutput(output, []string{"Usage:", "Examples:", "Main Commands:", "OpenShift Commands:", "Utility Commands:", "Flags:"})
					helper.DontMatchAllInOutput(output, []string{"--kubeconfig"})
				})
				It("does not support the --kubeconfig flag", func() {
					helper.DontMatchAllInOutput(output, []string{"--kubeconfig"})
				})
			})

			When("running particle engine without subcommand and flags", func() {
				var output string
				BeforeEach(func() {
					output = helper.Cmd("particle engine").ShouldPass().Out()
				})
				It("a short vesion of help contents is returned, an error is not expected", func() {
					Expect(output).To(ContainSubstring("To see a full list of commands, run 'particle engine --help'"))
				})
			})

			It("returns error when using an invalid command", func() {
				output := helper.Cmd("particle engine", "hello").ShouldFail().Err()
				Expect(output).To(ContainSubstring("Invalid command - see available commands/subcommands above"))
			})

			It("returns JSON error", func() {
				By("using an invalid command with JSON output", func() {
					res := helper.Cmd("particle engine", "unknown-command", "-o", "json").ShouldFail()
					stdout, stderr := res.Out(), res.Err()
					Expect(stdout).To(BeEmpty())
					Expect(helper.IsJSON(stderr)).To(BeTrue())
				})

				By("using an invalid describe sub-command with JSON output", func() {
					res := helper.Cmd("particle engine", "describe", "unknown-sub-command", "-o", "json").ShouldFail()
					stdout, stderr := res.Out(), res.Err()
					Expect(stdout).To(BeEmpty())
					Expect(helper.IsJSON(stderr)).To(BeTrue())
				})

				By("using an invalid list sub-command with JSON output", func() {
					res := helper.Cmd("particle engine", "list", "unknown-sub-command", "-o", "json").ShouldFail()
					stdout, stderr := res.Out(), res.Err()
					Expect(stdout).To(BeEmpty())
					Expect(helper.IsJSON(stderr)).To(BeTrue())
				})

				By("omitting required subcommand with JSON output", func() {
					res := helper.Cmd("particle engine", "describe", "-o", "json").ShouldFail()
					stdout, stderr := res.Out(), res.Err()
					Expect(stdout).To(BeEmpty())
					Expect(helper.IsJSON(stderr)).To(BeTrue())
				})
			})

			It("returns error when using an invalid command with --help", func() {
				output := helper.Cmd("particle engine", "hello", "--help").ShouldFail().Err()
				Expect(output).To(ContainSubstring("unknown command 'hello', type --help for a list of all commands"))
			})
		})
	}

	Context("When deleting two project one after the other", func() {
		It("should be able to delete sequentially", func() {
			project1 := helper.CreateRandProject()
			project2 := helper.CreateRandProject()

			helper.DeleteProject(project1)
			helper.DeleteProject(project2)
		})
		It("should be able to delete them in any order", func() {
			project1 := helper.CreateRandProject()
			project2 := helper.CreateRandProject()
			project3 := helper.CreateRandProject()

			helper.DeleteProject(project2)
			helper.DeleteProject(project1)
			helper.DeleteProject(project3)
		})
	})

	Context("executing particle engine version command", func() {
		const (
			reparticle engineVersion        = `^particle engine\s*v[0-9]+.[0-9]+.[0-9]+(?:-\w+)?\s*\(\w+(-\w+)?\)`
			reKubernetesVersion = `Kubernetes:\s*v[0-9]+.[0-9]+.[0-9]+((-\w+\.[0-9]+)?\+\w+)?`
			rePodmanVersion     = `Podman Client:\s*[0-9]+.[0-9]+.[0-9]+((-\w+\.[0-9]+)?\+\w+)?`
			reJSONVersion       = `^v{0,1}[0-9]+.[0-9]+.[0-9]+((-\w+\.[0-9]+)?\+\w+)?`
		)
		When("executing the complete command with server info", func() {
			var particle engineVersion string
			BeforeEach(func() {
				particle engineVersion = helper.Cmd("particle engine", "version").ShouldPass().Out()
			})
			for _, podman := range []bool{true, false} {
				podman := podman
				It("should show the version of particle engine major components including server login URL", helper.LabelPodmanIf(podman, func() {
					By("checking the human readable output", func() {
						Expect(particle engineVersion).Should(MatchRegexp(reparticle engineVersion))

						// particle engine tests setup (CommonBeforeEach) is designed in a way that if a test is labelled with 'podman', it will not have cluster configuration
						// so we only test podman info on podman labelled test, and clsuter info otherwise
						// Tparticle engine (pvala): Change this behavior when we write tests that should be tested on both podman and cluster simultaneously
						// Ref: https://github\.com/danielpickens/particle engine/issues/6719
						if podman {
							Expect(particle engineVersion).Should(MatchRegexp(rePodmanVersion))
							Expect(particle engineVersion).To(ContainSubstring(helper.GetPodmanVersion()))
						} else {
							Expect(particle engineVersion).Should(MatchRegexp(reKubernetesVersion))
							if !helper.IsKubernetesCluster() {
								serverURL := oc.GetCurrentServerURL()
								Expect(particle engineVersion).Should(ContainSubstring("Server: " + serverURL))
								ocpMatcher := ContainSubstring("OpenShift: ")
								if serverVersion := commonVar.CliRunner.GetVersion(); serverVersion == "" {
									// Might indicate a user permission error on certain clusters (observed with a developer account on Prow nightly jobs)
									ocpMatcher = Not(ocpMatcher)
								}
								Expect(particle engineVersion).Should(ocpMatcher)
							}
						}
					})

					By("checking the JSON output", func() {
						particle engineVersion = helper.Cmd("particle engine", "version", "-o", "json").ShouldPass().Out()
						Expect(helper.IsJSON(particle engineVersion)).To(BeTrue())
						helper.JsonPathSatisfiesAll(particle engineVersion, "version", MatchRegexp(reJSONVersion))
						helper.JsonPathExist(particle engineVersion, "gitCommit")
						if podman {
							helper.JsonPathSatisfiesAll(particle engineVersion, "podman.client.version", MatchRegexp(reJSONVersion), Equal(helper.GetPodmanVersion()))
						} else {
							helper.JsonPathSatisfiesAll(particle engineVersion, "cluster.kubernetes.version", MatchRegexp(reJSONVersion))
							if !helper.IsKubernetesCluster() {
								serverURL := oc.GetCurrentServerURL()
								helper.JsonPathContentIs(particle engineVersion, "cluster.serverURL", serverURL)
								m := BeEmpty()
								if serverVersion := commonVar.CliRunner.GetVersion(); serverVersion != "" {
									// A blank serverVersion might indicate a user permission error on certain clusters (observed with a developer account on Prow nightly jobs)
									m = Not(m)
								}
								helper.JsonPathSatisfiesAll(particle engineVersion, "cluster.openshift", m)
							}
						}
					})
				}))
			}

			for _, label := range []string{helper.LabelNoCluster, helper.LabelUnauth} {
				label := label
				It("should show the version of particle engine major components", Label(label), func() {
					Expect(particle engineVersion).Should(MatchRegexp(reparticle engineVersion))
				})
			}
		})

		When("podman client is bound to delay and particle engine version is run", Label(helper.LabelPodman), func() {
			var particle engineVersion string
			BeforeEach(func() {
				delayer := helper.GenerateDelayedPodman(commonVar.Context, 2)
				particle engineVersion = helper.Cmd("particle engine", "version").WithEnv("PODMAN_CMD="+delayer, "PODMAN_CMD_INIT_TIMEOUT=1s").ShouldPass().Out()
			})
			It("should not print podman version if podman cmd timeout has been reached", func() {
				Expect(particle engineVersion).Should(MatchRegexp(reparticle engineVersion))
				Expect(particle engineVersion).ToNot(ContainSubstring("Podman Client:"))
			})
		})
		It("should only print client info when using --client flag", func() {
			By("checking human readable output", func() {
				particle engineVersion := helper.Cmd("particle engine", "version", "--client").ShouldPass().Out()
				Expect(particle engineVersion).Should(MatchRegexp(reparticle engineVersion))
				Expect(particle engineVersion).ToNot(SatisfyAll(ContainSubstring("Server"), ContainSubstring("Kubernetes"), ContainSubstring("Podman Client")))
			})

			By("checking JSON output", func() {
				particle engineVersion := helper.Cmd("particle engine", "version", "--client", "-o", "json").ShouldPass().Out()
				Expect(helper.IsJSON(particle engineVersion)).To(BeTrue())
				helper.JsonPathSatisfiesAll(particle engineVersion, "version", MatchRegexp(reJSONVersion))
				helper.JsonPathExist(particle engineVersion, "gitCommit")
				helper.JsonPathSatisfiesAll(particle engineVersion, "cluster", BeEmpty())
				helper.JsonPathSatisfiesAll(particle engineVersion, "podman", BeEmpty())
			})
		})
	})

	Describe("Experimental Mode", Label(helper.LabelNoCluster), func() {
		AfterEach(func() {
			helper.ResetExperimentalMode()
		})

		When("experimental mode is enabled", func() {
			BeforeEach(func() {
				helper.EnableExperimentalMode()
			})

			AfterEach(func() {
				helper.ResetExperimentalMode()
			})

			It("should display warning message", func() {
				out := helper.Cmd("particle engine", "version", "--client").ShouldPass().Out()
				Expect(out).Should(ContainSubstring("Experimental mode enabled. Use at your own risk."))
			})
		})
	})

})
