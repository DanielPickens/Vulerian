package integration

import (
	"os"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github\.com/danielpickens/particle engine/tests/helper"
)

var _ = Describe("particle engine login and logout command tests", func() {
	// user related constants
	const loginTestUserForNoProject = "particle engineloginnoproject"
	const loginTestUserPassword = "password@123"
	var session1 string
	var testUserToken string
	var oc helper.OcRunner
	var currentUserToken string

	BeforeEach(func() {
		SetDefaultEventuallyTimeout(10 * time.Minute)
		SetDefaultConsistentlyDuration(30 * time.Second)
		oc = helper.NewOcRunner("oc")
	})

	Context("when running help for login command", func() {
		It("should display the help", func() {
			appHelp := helper.Cmd("particle engine", "login", "-h").ShouldPass().Out()
			Expect(appHelp).To(ContainSubstring("Login to cluster"))
		})
	})

	Context("when running help for logout command", func() {
		It("should display the help", func() {
			appHelp := helper.Cmd("particle engine", "logout", "-h").ShouldPass().Out()
			Expect(appHelp).To(ContainSubstring("Logout of the cluster"))
		})
	})

	It("should show cluster version when not logged in", func() {
		version := helper.Cmd("particle engine", "version").ShouldPass().Out()
		helper.MatchAllInOutput(version, []string{"particle engine", "Server", "Kubernetes"})
	})

	Context("when running login tests", func() {
		It("should successful with correct credentials and fails with incorrect token", func() {
			if strings.ToLower(os.Getenv("CI")) != "openshift" {
				Skip("Skipping if not running on OpenShift. Set CI environment variable to openshift.")
			}
			// skip if requested
			skipLogin := os.Getenv("SKIP_USER_LOGIN_TESTS")
			if skipLogin == "true" {
				Skip("Skipping login command tests as SKIP_USER_LOGIN_TESTS is true")
			}
			// Current user login token
			currentUserToken = oc.GetToken()

			// Login successful without any projects with appropriate message
			session1 = helper.Cmd("particle engine", "login", "-u", loginTestUserForNoProject, "-p", loginTestUserPassword).ShouldPass().Out()
			Expect(session1).To(ContainSubstring("Login successful"))
			Expect(session1).To(ContainSubstring("You don't have any projects. You can try to create a new project, by running"))
			Expect(session1).To(ContainSubstring("particle engine create project <project-name>"))
			session1 = oc.GetLoginUser()
			Expect(session1).To(ContainSubstring(loginTestUserForNoProject))

			// particle engineloginnoproject user login token
			testUserToken = oc.GetToken()

			// Login successful with token without any projects with appropriate message
			session1 = helper.Cmd("particle engine", "login", "-t", testUserToken).ShouldPass().Out()
			Expect(session1).To(ContainSubstring("Logged into"))
			Expect(session1).To(ContainSubstring("You don't have any projects. You can try to create a new project, by running"))
			Expect(session1).To(ContainSubstring("particle engine create project <project-name>"))
			session1 = oc.GetLoginUser()
			Expect(session1).To(ContainSubstring(loginTestUserForNoProject))

			// Login fails on invalid token with appropriate message
			sessionErr := helper.Cmd("particle engine", "login", "-t", "verybadtoken").ShouldFail().Err()
			Expect(sessionErr).To(ContainSubstring("The token provided is invalid or expired"))

			// loging back to current user
			helper.Cmd("particle engine", "login", "--token", currentUserToken).ShouldPass()
		})
	})
})
