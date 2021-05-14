package main

import (
	//releaseUtils "../internal"
	releaseManager "../release_manager"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

var (
	outputStream io.Writer = os.Stdout
	errStream    io.Writer = os.Stderr
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

	//cmd, err := exec.Command(
	//	"/bin/bash", pathway,
	//	release.EnvironmentReleasePath,
	//	release.Environment,
	//	release.Version,
	//).Output()
	//if err != nil {
	//	fmt.Printf("error %s", err)
	//}
	//output := string(cmd)
	//fmt.Println(output)

	//c.makeArgs = []string{
	//	fmt.Sprintf("RELEASE_BRANCH=%s", c.releaseBranch),
	//	fmt.Sprintf("RELEASE=%s", c.release),
	//	fmt.Sprintf("AWS_REGION=%s", *region),
	//	fmt.Sprintf("AWS_ACCOUNT_ID=%s", *accountId),
	//	fmt.Sprintf("IMAGE_REPO=%s", *imageRepo),
	//}

	//commandlrgs = append(commandArgs, c.makeArgs...)

	cmdTwo := exec.Command("/bin/bash", pathway, release.EnvironmentReleasePath, release.Environment, release.Version)
	//log.Printf("Executing: %s", strings.Join(cmd.Args, " "))
	cmdTwo.Stdout = outputStream
	cmdTwo.Stderr = errStream
	//if !c.dryRun {
	err = cmdTwo.Run()
	if err != nil {
		log.Fatalf("Error running make: %v", err)

	}
}
