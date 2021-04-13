package main


const changeLogBase = `# Changelog for {{.ReleaseVersionName}}

This changelog highlights the changes for [{{.ReleaseVersionName}}](https://github.com/aws/eks-distro/tree/{{.ReleaseVersionName}}).

`

const changeLogOnlyAL2 = changeLogBase + `## Changes
Security updates to Amazon Linus 2.

`
