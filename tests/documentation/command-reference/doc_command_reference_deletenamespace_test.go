package docautomation

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/go-cmp/cmp"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github\.com/danielpickens/particle engine/tests/helper"
)

var _ = Describe("doc command reference particle engine delete namespace", func() {
	var commonVar helper.CommonVar
	var commonPath = filepath.Join("command-reference", "docs-mdx", "delete-namespace")
	var outputStringFormat = "```console\n$ particle engine %s\n%s```\n"
	var ns string

	BeforeEach(func() {
		commonVar = helper.CommonBeforeEach()
		helper.Chdir(commonVar.Context)
		ns = helper.GenerateProjectName()
	})

	AfterEach(func() {
		helper.CommonAfterEach(commonVar)
	})

	Context("To delete a namespace resource", func() {

		BeforeEach(func() {
			helper.Cmd("particle engine", "create", "namespace", ns).ShouldPass()
		})

		AfterEach(func() {
			if commonVar.CliRunner.HasNamespaceProject(ns) {
				commonVar.CliRunner.DeleteNamespaceProject(ns, false)
			}
		})

		It("Deletes a namespace resource for a kubernetes cluster", func() {
			args := []string{"particle engine", "delete", "namespace", ns}
			out, err := helper.RunInteractive(args, []string{"particle engine_LOG_LEVEL=0"}, func(ctx helper.InteractiveContext) {
				helper.ExpectString(ctx, fmt.Sprintf("? Are you sure you want to delete namespace %q?", ns))
				helper.SendLine(ctx, "Yes")

			})
			Expect(err).To(BeNil())
			got := helper.StripAnsi(out)
			got = helper.StripInteractiveQuestion(got)
			got = fmt.Sprintf(outputStringFormat, args[1], helper.StripSpinner(got))
			got = strings.ReplaceAll(got, ns, "particle engine-dev")
			file := "delete_namespace.mdx"
			want := helper.GetMDXContent(filepath.Join(commonPath, file))
			diff := cmp.Diff(want, got)
			Expect(diff).To(BeEmpty(), file)
		})

		It("Deletes a project resource for a openshift cluster", func() {
			args := []string{"particle engine", "delete", "project", ns}
			out, err := helper.RunInteractive(args, []string{"particle engine_LOG_LEVEL=0"}, func(ctx helper.InteractiveContext) {
				helper.ExpectString(ctx, fmt.Sprintf("? Are you sure you want to delete project %q?", ns))
				helper.SendLine(ctx, "Yes")

			})
			Expect(err).To(BeNil())
			got := helper.StripAnsi(out)
			got = helper.StripInteractiveQuestion(got)
			got = fmt.Sprintf(outputStringFormat, args[1], helper.StripSpinner(got))
			got = strings.ReplaceAll(got, ns, "particle engine-dev")
			file := "delete_project.mdx"
			want := helper.GetMDXContent(filepath.Join(commonPath, file))
			diff := cmp.Diff(want, got)
			Expect(diff).To(BeEmpty(), file)
		})

	})
})
