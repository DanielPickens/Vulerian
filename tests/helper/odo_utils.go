package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/devfile/api/v2/pkg/apis/workspaces/v1alpha2"
	devfilepkg "github.com/devfile/api/v2/pkg/devfile"
	"github.com/devfile/library/v2/pkg/devfile/parser/data/v2/common"

	"github\.com/danielpickens/particle engine/pkg/devfile"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// GetPreferenceValue a global config value of given key or
// returns an empty string if value is not set
func GetPreferenceValue(key string) string {
	stdOut := Cmd("particle engine", "preference", "view").ShouldPass().Out()
	re := regexp.MustCompile(" " + key + `.+`)
	particle engineConfigKeyValue := re.FindString(stdOut)
	if particle engineConfigKeyValue == "" {
		return fmt.Sprintf("%s not found", key)
	}
	trimKeyValue := strings.TrimSpace(particle engineConfigKeyValue)
	if strings.Compare(key, trimKeyValue) != 0 {
		return strings.TrimSpace(strings.SplitN(trimKeyValue, " ", 2)[1])
	}
	return ""
}

// CreateRandProject create new project with random name (10 letters)
// without writing to the config file (without switching project)
func CreateRandProject() string {
	projectName := GenerateProjectName()
	fmt.Fprintf(GinkgoWriter, "Creating a new project: %s\n", projectName)
	session := Cmd("particle engine", "create", "project", projectName, "-w", "-v4").ShouldPass().Out()
	Expect(session).To(ContainSubstring("New project created"))
	Expect(session).To(ContainSubstring(projectName))
	return projectName
}

// DeleteProject deletes a specified project
func DeleteProject(projectName string) {
	fmt.Fprintf(GinkgoWriter, "Deleting project: %s\n", projectName)
	session := Cmd("particle engine", "delete", "project", projectName, "-f").ShouldPass().Out()
	Expect(session).To(ContainSubstring(fmt.Sprintf("Project %q will be deleted asynchronously", projectName)))
}

// GetMetadataFromDevfile retrieves the metadata from devfile
func GetMetadataFromDevfile(devfilePath string) devfilepkg.DevfileMetadata {
	devObj, err := devfile.ParseAndValidateFromFile(devfilePath, "", true)
	Expect(err).ToNot(HaveOccurred())
	return devObj.Data.GetMetadata()
}

func GetDevfileComponents(devfilePath, componentName string) []v1alpha2.Component {
	devObj, err := devfile.ParseAndValidateFromFile(devfilePath, "", true)
	Expect(err).ToNot(HaveOccurred())
	components, err := devObj.Data.GetComponents(common.DevfileOptions{
		FilterByName: componentName,
	})
	Expect(err).ToNot(HaveOccurred())
	return components
}

type particle engineV2Watch struct {
	CmpName               string
	StringsToBeMatched    []string
	StringsNotToBeMatched []string
	FolderToCheck         string
	SrcType               string
}

// VerifyContainerSyncEnv verifies the sync env in the container
func VerifyContainerSyncEnv(podName, containerName, namespace, projectSourceValue, projectsRootValue string, cliRunner CliRunner) {
	envProjectsRoot, envProjectSource := "PROJECTS_ROOT", "PROJECT_SOURCE"
	projectSourceMatched, projectsRootMatched := false, false

	envNamesAndValues := cliRunner.GetContainerEnv(podName, "runtime", namespace)
	envNamesAndValuesArr := strings.Fields(envNamesAndValues)

	for _, envNamesAndValues := range envNamesAndValuesArr {
		envNameAndValueArr := strings.Split(envNamesAndValues, ":")

		if envNameAndValueArr[0] == envProjectSource && strings.Contains(envNameAndValueArr[1], projectSourceValue) {
			projectSourceMatched = true
		}

		if envNameAndValueArr[0] == envProjectsRoot && strings.Contains(envNameAndValueArr[1], projectsRootValue) {
			projectsRootMatched = true
		}
	}

	Expect(projectSourceMatched).To(Equal(true))
	Expect(projectsRootMatched).To(Equal(true))
}
