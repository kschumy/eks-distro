package release_manager

import (
	"fmt"
	"strings"
)

// Returns <major version>.<minor version>-<patch version> (e.g. 1.20-2)
func createReleaseVersion(release *Release) string {
	if len(release.Version) == 0 {
		return fmt.Sprintf("%s-%s", strings.Replace(release.branch, "-", ".", 1), release.Number)
	}
	return release.Version
}