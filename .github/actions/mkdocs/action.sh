#!/bin/bash

# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -ex

REQUIREMENTS="${GITHUB_WORKSPACE}/vertical-pod-autoscaler/docs/requirements.txt"

if [ -f "${REQUIREMENTS}" ]; then
    pip install -r "${REQUIREMENTS}"
fi

if [ -n "${GITHUB_TOKEN}" ]; then
    remote_repo="https://x-access-token:${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
elif [ -n "${PERSONAL_TOKEN}" ]; then
    remote_repo="https://x-access-token:${PERSONAL_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"
fi

git config --global user.name "$GITHUB_ACTOR"
git config --global user.email "$GITHUB_ACTOR@users.noreply.github.com"

mkdocs build --config-file "${GITHUB_WORKSPACE}/vertical-pod-autoscaler/mkdocs.yml"

git clone --branch=gh-pages --depth=1 "${remote_repo}" gh-pages
cd gh-pages
# remove current content in branch gh-pages
git rm -r .
# copy new doc.
mkdir vertical-pod-autoscaler/
cp -r ../vertical-pod-autoscaler/site/* ./vertical-pod-autoscaler/

cat << EOF > ./index.html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Redirecting...</title>
    <meta http-equiv="refresh" content="0; URL='/autoscaler/vertical-pod-autoscaler/'" />
</head>
<body>
    <p>If you are not redirected automatically, follow this <a href="/autoscaler/vertical-pod-autoscaler/">link to /autoscaler/vertical-pod-autoscaler/</a>.</p>
</body>
</html>
EOF

# commit changes
git add .
git commit -m "Deploy GitHub Pages"
git push --force --quiet "${remote_repo}" gh-pages > /dev/null 2>&1
