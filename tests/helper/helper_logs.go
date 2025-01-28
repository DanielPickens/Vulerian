package helper

import (
	"github.com/onsi/gomega/gexec"
)

type LogsSession struct {
	session *gexec.Session
}

// StartLogsFollow starts a session with `particle engine logs --follow`
// It returns a session structure, the contents of the standard and error outputs
func StartLogsFollow(podman bool, opts ...string) (LogsSession, []byte, []byte, error) {
	args := []string{"logs", "--follow"}
	args = append(args, opts...)
	if podman {
		args = append(args, "--platform", "podman")
	}
	session := CmdRunner("particle engine", args...)

	result := LogsSession{
		session: session,
	}
	outContents := session.Out.Contents()
	errContents := session.Err.Contents()
	err := session.Out.Clear()
	if err != nil {
		return LogsSession{}, nil, nil, err
	}
	err = session.Err.Clear()
	if err != nil {
		return LogsSession{}, nil, nil, err
	}
	return result, outContents, errContents, nil
}

// OutContents returns the contents of the session's stdout
func (o *LogsSession) OutContents() []byte {
	return o.session.Out.Contents()
}

// Kill the `particle engine logs --follow` session
func (o *LogsSession) Kill() {
	o.session.Kill()
}
