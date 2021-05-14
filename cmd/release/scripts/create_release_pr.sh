#!/usr/bin/env bash
# Copyright 2020 Amazon.com Inc. or its affiliates. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -e
set -o pipefail
set -x

RELEASE_FILEPATH="${1?....}"
RELEASE_ENVIRONMENT="${2?Should be 'development' or 'production'}"
RELEASE_VERSION="${3?Release branch}"


echo "hellooo"

IS_BOT=false

ORIGIN_ORG=$(git remote get-url origin | sed -n -e "s|git@github.com:\(.*\)/eks-distro.git|\1|p")

PR_TITLE="Increment ${RELEASE_ENVIRONMENT} RELEASE for ${RELEASE_VERSION}"
COMMIT_MESSAGE="TEST!! [PR BOT] Increment RELEASE for"
#
PR_BODY=$(cat <<EOF
TEST!! Bumping RELEASE version

By submitting this pull request, I confirm that you can use, modify, copy, and redistribute this contribution, under the terms of your choice.
EOF
)
#
PR_BRANCH="increment-${RELEASE_ENVIRONMENT}-RELEASE-${RELEASE_VERSION}" #"automated-release-update" #"increment-development-RELEASE-1.19-28"

echo $PR_BRANCH

git checkout -B $PR_BRANCH


if [[ "$(git status --porcelain | wc -l)" -eq 1 ]]; then
  git add "${RELEASE_FILEPATH}"
  if [[ $(git diff --staged --name-only) == "" ]]; then
    exit 0
  fi
  git commit -m "${COMMIT_MESSAGE}" || true
else
  git restore "${RELEASE_FILEPATH}"
  echo "Unexpected files."
  echo "Restored ${RELEASE_FILEPATH}"
  exit 1
fi

echo "pushing..."
git push origin ${PR_BRANCH}

#echo $PR_BRANCH
echo "pushing?"

PR_EXISTS=$(gh pr list | grep -c "${PR_BRANCH}" || true)
  echo "PR_EXISTS?"
  echo $PR_EXISTS

if [ "${PR_EXISTS}" -eq 0 ]; then
    echo "INSIDE"
  gh pr create --title "${PR_TITLE}" --body "${PR_BODY}" --web --repo "aws/eks-distro"
fi
