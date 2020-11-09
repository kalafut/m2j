package m2j

import (
	"bytes"
	"os/exec"
	"strings"
)

// MDToJira takes a string in Github Markdown, and outputs Jira text styling
func MDToJira(str string) (string, error) {
	cmd := exec.Command("docker", "run", "--interactive", "--rm", "pandoc/core:2.11.0.4", "--from=gfm", "--to=jira")
	cmd.Stdin = strings.NewReader(str)
	out := bytes.Buffer{}
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}
