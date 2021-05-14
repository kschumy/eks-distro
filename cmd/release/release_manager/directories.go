package release_manager

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

var gitRootDirectory = GetGitRootDirectory()

func GetGitRootDirectory() string {
	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(fmt.Sprintf("Unable to get git root directory: %v", err))
	}
	return strings.Join(strings.Fields(string(gitRootOutput)), "")
}


func getReleasePath(release *Release) string {
	return filepath.Join(gitRootDirectory, "release", release.branch, release.Environment, "RELEASE")
}