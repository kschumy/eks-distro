package main

import (
	//releaseUtils "../internal"
	releaseManager "../release_manager"
	"log"
	"os"

	"flag"
	"fmt"
	"os/exec"
)

type Input struct {
	releaseBranch      string
	releaseEnvironment string
}

func (input Input) GetReleaseBranch() string {
	return input.releaseBranch
}

func (input Input) GetReleaseEnvironment() string {
	return input.releaseEnvironment
}

func main() {
	releaseBranch := flag.String("pkg-branch", "", "Release releaseBranch")
	releaseEnvironment := flag.String("pkg-environment", "development", "Must be 'development' or 'production'")

	flag.Parse()

	releaseInput := &Input{
		releaseBranch:      *releaseBranch,
		releaseEnvironment: *releaseEnvironment,
	}

	release, err := releaseManager.IntiRelease(releaseInput)
	if err != nil {
		log.Fatalf("error initializing release values: %v", err)
	}

	err = os.WriteFile(release.EnvironmentReleasePath, []byte(release.Number+"\n"), 0644)
	if err != nil {
		log.Fatalf("error wrirting to RELEASE: %v", err)
	}
	rootdir := releaseManager.GetGitRootDirectory()

	pathway := rootdir + "/cmd/release/scripts/create_release_pr.sh"
	fmt.Println(pathway)

	cmd, err := exec.Command(
		"/bin/bash", pathway,
		release.EnvironmentReleasePath,
		release.Environment,
		release.Version,
	).Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	fmt.Println(output)
}
