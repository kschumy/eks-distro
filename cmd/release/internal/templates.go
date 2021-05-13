package internal

const ChangeLogBase = `# Changelog for {{.ReleaseVersionName}}

This changelog highlights the changes for [{{.ReleaseVersionName}}](https://github.com/aws/eks-distro/tree/{{.ReleaseVersionName}}).

`

const ChangeLogOnlyAL2 = ChangeLogBase + `## Changes
Security updates to Amazon Linux 2.

`
