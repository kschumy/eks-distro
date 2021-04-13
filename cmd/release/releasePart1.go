package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

// PR #1
//
//Create new directory
//    - docs/contents/releases/${RELEASE_BRANCH}/${RELEASE}
//Create new files and populate them
//    - docs/contents/releases/${RELEASE_BRANCH}/${RELEASE}/index.md
//    - docs/contents/releases/${RELEASE_BRANCH}/${RELEASE}/CHANGELOG
//Update existing files
//    - projects/kubernetes/kubernetes/${RELEASE_BRANCH}/KUBE_GIT_VERSION_FILE
//    - number/${RELEASE_BRANCH}/development/RELEASE
//    - docs/contents/index.md
//    - README
// PR #2
//
//Create new files and populate them
//
//
//    - docs/contents/releases/${RELEASE_BRANCH}/ ${RELEASE}/number-announcement.txt
//
//Update existing files
//    - number/${RELEASE_BRANCH}/production/RELEASE



type Release struct {
	versionName string
	branch      string
	gitRoot     string
	number      string
	isDevelopment bool
}

func main() {
	releaseBranch := flag.String("number-branch", "1-19", "Release branch to test")
	number := flag.String("number", "1", "Release to test")
	isDevelopment := flag.Bool("development", false, "Build as a development build")
	//dryRun := flag.Bool("dry-run", false, "Echo out commands, but don't run them")
	//
	//flag.Parse()
	//log.Printf("Running postsubmit - dry-run: %t", *dryRun)
	//
	release := &Release{
		branch:      *releaseBranch,
		number:      *number,
		isDevelopment: *isDevelopment,
		//gitRoot:     "",
		//dryRun:         *dryRun,
	}

	gitRoot, err := getRootDirectory()
	if err != nil {
		log.Fatalf("Error running finding git root: %v", err)
	}
	release.gitRoot = gitRoot

	//createEmptyFile := func(name string) {
	//	d := []byte("")
	//	ioutil.WriteFile(name, d, 0644)
	//}






	releasePath, _ := getReleasePath(release)
	print("releasePath: ")
	println(releasePath)

	//_, err = os.Stat(releasePath)
	//if errors.Is(err, os.ErrNotExist) {
	//	err = os.Mkdir(releasePath, os.ModePerm)
	//	if err != nil {
	//		log.Fatalf("Error making dir: %v", err)
	//	}
	//} else if err == nil {
	//	log.Fatalf("Expected directory %v not to exist already", releasePath)
	//} else {
	//	log.Fatalf("Encountered unepected error when trying to check directory %v", releasePath)
	//}

	//
	//		ioutil.WriteFile("view.html", []byte(`<html>
	//<head>
	//    <title>First Program</title>
	//</head>
	//<body>
	//    {{ . }}
	//</body>
	//</html>`), 0666)
	//

	releaseVersionName := fmt.Sprintf("v%v-eks-%v", releaseBranch, number)
	//createEmptyFile(fmt.Sprintf(releasePath + "/CHANGELOG-%v.md", releaseVersionName))

	//releaseInfo := Release{
	//	ReleaseVersionName: releaseVersionName,
	//}
	t := template.Must(template.New("changeLogText").Parse(changeLogOnlyAL2))

	f, err := os.Create(fmt.Sprintf(releasePath+"/CHANGELOG-%v.md", releaseVersionName))
	if err != nil {
		log.Fatalf("Error !!!: %v", err)
	}
	w := io.Writer(f)
	//var buff bytes.Buffer

	//_, err = w.Write([]byte("foo"))

	if err := t.Execute(w, release); err != nil {
		panic(err)
	}
	d1 := []byte(release.getNumber()+"\n")
	err = f.Truncate(0)
	//err = ioutil.WriteFile(path, []byte(newContents), 0)
	err = ioutil.WriteFile(fmt.Sprintf("%v/number/%v/development/RELEASE", gitRoot, releaseBranch), d1, 0644)
//docs/contents/releases/
	//c.makeArgs = []string{
	//	fmt.Sprintf("RELEASE_BRANCH=%s", c.branch),
	//	fmt.Sprintf("RELEASE=%s", c.number),
	//	fmt.Sprintf("DEVELOPMENT=%t", *development),
	//}
	//
	//cmd := exec.Command("git", "-C", *gitRoot, "diff", "--name-only", "HEAD^", "HEAD")
	//log.Printf("Executing command: %s", strings.Join(cmd.Args, " "))
	//gitDiffOutput, err := cmd.Output()
	//if err != nil {
	//	log.Fatalf("error running git diff: %v\n%s", err, string(gitDiffOutput))
	//}
	//filesChanged := strings.Fields(string(gitDiffOutput))
	//
	//buildOrder := [...]string{
	//	"kubernetes/number",
	//	"kubernetes/kubernetes",
	//	"containernetworking/plugins",
	//	"coredns/coredns",
	//	"etcd-io/etcd",
	//	"kubernetes-sigs/aws-iam-authenticator",
	//	"kubernetes-sigs/metrics-server",
	//	"kubernetes-csi/external-attacher",
	//	"kubernetes-csi/external-resizer",
	//	"kubernetes-csi/livenessprobe",
	//	"kubernetes-csi/node-driver-registrar",
	//	"kubernetes-csi/external-snapshotter",
	//	"kubernetes-csi/external-provisioner",
	//}
	//type changedStruct struct {
	//	changed bool
	//}
	//projects := make(map[string]*changedStruct)
	//for _, projectPath := range buildOrder {
	//	projects[projectPath] = &changedStruct{}
	//}
	//
	//for _, file := range filesChanged {
	//	for projectPath := range projects {
	//		if strings.Contains(file, projectPath) {
	//			projects[projectPath].changed = true
	//		}
	//	}
	//	release := regexp.MustCompile("Makefile|cmd/main_postsubmit.go|EKS_DISTRO_BASE_TAG_FILE|number/.*")
	//	if release.MatchString(file) {
	//		allChanged = true
	//	}
	//}

}

func (r *Release) getGitRoot() string {
	return r.gitRoot
}

func (r *Release) getNumber() string {
	return r.number
}

func (r *Release) getBranch() string {
	return r.number
}

