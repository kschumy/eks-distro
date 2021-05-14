package release_manager

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Release struct {
	branch             string
	Number             string
	prevNumber         string
	Environment        string

	Version string

	EnvironmentReleasePath string
}

type ReleaseInput interface {
	GetReleaseBranch() string
	GetReleaseEnvironment() string
}

func IntiRelease(input ReleaseInput) (*Release, error) {
	release := &Release{
		branch:      input.GetReleaseBranch(),
		Environment: input.GetReleaseEnvironment(),
	}

	var err error

	release.EnvironmentReleasePath = getReleasePath(release)

	release.prevNumber, err = determinePreviousReleaseNumber(release)
	if err != nil {
		return &Release{}, err
	}
	log.Printf("Determined %q is the previous release number\n", release.prevNumber)

	release.Number, err = determineReleaseNumber(release)
	if err != nil {
		return &Release{}, err
	}
	log.Printf("Determined %q is the release number\n", release.Number)

	release.Version = createReleaseVersion(release)

	return release, nil
}


func determinePreviousReleaseNumber(release *Release) (string, error) {
	if len(release.prevNumber) > 0 {
		log.Printf("release number %q already known and is not re-sought\n", release.prevNumber)
		return release.prevNumber, nil
	}

	fileOutput, err := ioutil.ReadFile(release.EnvironmentReleasePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(fileOutput)), nil
}

func determineReleaseNumber(release *Release) (string, error) {
	prevReleaseNumber := release.prevNumber

	if len(prevReleaseNumber) == 0 {
		prevReleaseNumber, err := determinePreviousReleaseNumber(release)
		if err != nil {
			return "", err
		}
		log.Printf("previous release number not provided to determime release number. It is assumed to be %q\n",
			prevReleaseNumber)
	}

	prevNumberAsInt, err := strconv.Atoi(prevReleaseNumber)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(prevNumberAsInt + 1), nil
}
