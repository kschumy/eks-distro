package pkg

import (
	"os/exec"
	"strings"
)



type Project interface {
	//getGitRoot()	string
	getNumber()	string
	getBranch() string
}

func getRootDirectory() (string, error) {
	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}

	return strings.Join(strings.Fields(string(gitRootOutput)), ""), nil
}


func GetRootDirectory() (string, error) {
	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", err
	}
	return strings.Join(strings.Fields(string(gitRootOutput)), ""), nil
}

//// Returns: <project_root>/docs/contents/releases/${RELEASE_BRANCH}/${RELEASE}
//// Does not check if this path is valid or already existing
//func getReleasePath(p Project) (string, error) {
//	//gitRoot := p.getGitRoot()
//	//if p.getGitRoot() == "" { // or nil?
//	gitRoot, err := GetRootDirectory() // FIXME
//	if err != nil {
//		return "", err
//	}
//	//}
//	return fmt.Sprintf("%v/docs/contents/releases/%v/%v", gitRoot, p.getBranch(), p.getNumber()), nil
//}


