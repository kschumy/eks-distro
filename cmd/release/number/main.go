package main

import (
	pkgDir "../internal"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
)

var (
//gitRoot string = getRootDirectorySafe()
)



type Release struct {
	releaseBranch string
	number        string
	prevNumber    string
	//isCompleteNeeded   bool
	releaseEnvironment string
	ReleaseVersionName string
}

func (release Release) GetReleaseNumber() string {
	return release.number
}

func (release Release) GetReleaseBranch() string {
	return release.releaseBranch
}

// TODO splite structs into two
func (release Release) GetReleaseEnvironment() string {
	return release.releaseEnvironment
}


//func getRootDirectorySafe() string {
//	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
//	if err != nil {
//		return ""
//	}
//	return strings.Join(strings.Fields(string(gitRootOutput)), "")
//}

//func getRootDirectory() (string, error) {
//	gitRootOutput, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
//	if err != nil {
//		return "", err
//	}
//
//	return strings.Join(strings.Fields(string(gitRootOutput)), ""), nil
//}



func getFileContentsAsString(filepath string) (string, error) {
	fileOutput, err := ioutil.ReadFile(filepath)
	return string(fileOutput), err
}

func getFileContentsTrimmedAsString(filepath string) (string, error) {
	fileOutput, err := getFileContentsAsString(filepath)
	return strings.TrimSpace(fileOutput), err
}

// TODO: change to error handle
func check(e error) {
	if e != nil {
		panic(e)
	}
}



func main() {
	releaseBranch := flag.String("pkg-branch", "", "Release releaseBranch")
	releaseEnvironment := flag.String("pkg-environment", "development", "Must be 'development' or 'production'")
	number := flag.String("number", "", "Release to test")
	//prevNumber := flag.String("prevNumber", "", "Release to test")
	//isCompleteNeeded := flag.Bool("is-complete-needed", false, "True if is automates")

	flag.Parse()

	release := &Release{
		releaseBranch:      *releaseBranch,
		releaseEnvironment: *releaseEnvironment,
		number:             *number,
		//prevNumber:       *prevNumber,
		//isCompleteNeeded: *isCompleteNeeded,
	}

	releasePath, err := pkgDir.GetReleasePath(release)
	check(err)
	release.prevNumber, err = getFileContentsTrimmedAsString(releasePath) //getFileContent(releasePath, 2)
	check(err)

	prevNumberAsInt, err := strconv.Atoi(release.prevNumber)
	check(err)
	if len(release.number) == 0 {
		release.number = strconv.Itoa(prevNumberAsInt + 1)
	} else {
		numberAsInt, err := strconv.Atoi(release.number)
		check(err)

		if numberAsInt <= prevNumberAsInt {
			panic("cannot have this") // TODO better message
		} else if numberAsInt != prevNumberAsInt+1 {
			fmt.Println("WARNING! Increase in numbers is greater than 1") // TODO better message
		}
	}

	release.ReleaseVersionName = fmt.Sprintf("v%v-eks-%v", release.releaseBranch, release.number)

	err = os.WriteFile(releasePath, []byte(release.number+"\n"), 0644)
	check(err)

	//////////////////////////////

	kubeGitVersionFilePath, _ := pkgDir.GetKubeGitVersionFilePath(release)

	b, err := ioutil.ReadFile(kubeGitVersionFilePath)
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(strings.Join([]string{"eks", release.releaseBranch, release.prevNumber}, "-"))
	if !re.Match(b) {
		panic("no match")
	}
	b = re.ReplaceAll(b, []byte(strings.Join([]string{"eks", release.releaseBranch, release.number}, "-")))

	err = os.WriteFile(kubeGitVersionFilePath, b, 0644)
	check(err)

	//////////////////////////////

	releaseDocsPath, _ := pkgDir.GetReleaseDocsPath(release)

	os.Mkdir(releaseDocsPath,  0777)

	t := template.Must(template.New("changeLogText").Parse(pkgDir.ChangeLogOnlyAL2))

	f, err := os.Create(fmt.Sprintf(releaseDocsPath+"/CHANGELOG-%v.md", release.ReleaseVersionName))
	if err != nil {
		log.Fatalf("Error !!!: %v", err)
	}
	w := io.Writer(f)
	if err := t.Execute(w, release); err != nil {
		panic(err)
	}

}
