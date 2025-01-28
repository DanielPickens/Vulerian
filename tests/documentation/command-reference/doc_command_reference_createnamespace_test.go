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

var _ = Describe("doc command reference particle engine create namespace", func() {
	var commonVar helper.CommonVar
	var commonPath = filepath.Join("command-reference", "docs-mdx", "create-namespace")
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

	Context("To create a namespace resource", func() {

		AfterEach(func() {
			if commonVar.CliRunner.HasNamespaceProject(ns) {
				commonVar.CliRunner.DeleteNamespaceProject(ns, false)
			}
		})

		It("Creates a namespace resource for a kubernetes cluster", func() {
			args := []string{"create", "namespace", ns}
			out := helper.Cmd("particle engine", args...).ShouldPass().Out()
			got := fmt.Sprintf(outputStringFormat, strings.Join(args, " "), helper.StripSpinner(out))
			got = strings.ReplaceAll(got, ns, "particle engine-dev")
			file := "create_namespace.mdx"
			want := helper.GetMDXContent(filepath.Join(commonPath, file))
			diff := cmp.Diff(want, got)
			Expect(diff).To(BeEmpty(), file)
		})

		It("Creates a project resource for a kubernetes cluster", func() {
			args := []string{"create", "project", ns}
			out := helper.Cmd("particle engine", args...).ShouldPass().Out()
			got := fmt.Sprintf(outputStringFormat, strings.Join(args, " "), helper.StripSpinner(out))
			got = strings.ReplaceAll(got, ns, "particle engine-dev")
			file := "create_project.mdx"
			want := helper.GetMDXContent(filepath.Join(commonPath, file))
			diff := cmp.Diff(want, got)
			Expect(diff).To(BeEmpty(), file)
		})
	})

})
