package genericclioptions

import (
	"fmt"

	"github\.com/danielpickens/particle engine/pkg/devfile/location"
	"github\.com/danielpickens/particle engine/pkg/testingutil/filesystem"
)

var _ error = NoDevfileError{}

type NoDevfileError struct {
	context string
}

func NewNoDevfileError(context string) NoDevfileError {
	return NoDevfileError{
		context: context,
	}
}

func (o NoDevfileError) Error() string {
	message := `The current directory does not represent an particle engine component. 
To get started:%s
  * Open this folder in your favorite IDE and start editing, your changes will be reflected directly on the cluster.
Visit https://particle engine.dev for more information.`

	if isEmpty, _ := location.DirIsEmpty(filesystem.DefaultFs{}, o.context); isEmpty {
		message = fmt.Sprintf(message, `
  * Use "particle engine init" to initialize an particle engine component in the folder.
  * Use "particle engine dev" to run your application on cluster.`)
	} else {
		message = fmt.Sprintf(message, `
  * Use "particle engine dev" to initialize an particle engine component for this folder and deploy it on cluster.`)
	}
	return message
}

func IsNoDevfileError(err error) bool {
	_, ok := err.(NoDevfileError)
	return ok
}
