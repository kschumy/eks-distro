package main

import (
	"fmt"
	"strconv"
)

// Returns 'v<RELEASE_BRANCH>-eks-<RELEASE>'. e.g. 'v1-19-eks-3'
func getVersionName(project Project) string {
	return fmt.Sprintf("v%v-eks-%v", project.getBranch(), project.getNumber())
}


func getPreviousVersionName(project Project) (string, error) {
	numberAsInt, err := strconv.Atoi(project.getNumber())
	if err != nil {
		return "", err
	} else if numberAsInt == 1 {
		return "", fmt.Errorf("no previous patch versinos for release %v",  project.getBranch())
	}

	previousNumber := numberAsInt - 1

	return fmt.Sprintf("v%v-eks-%v", project.getBranch(), strconv.Itoa(previousNumber)), nil // TODO: remove convert?
}
