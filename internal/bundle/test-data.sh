#!/bin/bash
mkdir -p /Users/mkong/Projects/opcap/internal/bundle/data.test ; cd /Users/mkong/Projects/opcap/internal/bundle/data.test
git init --bare test-data.git

mkdir -p certified-operators/operators

git init
cp -R /Users/mkong/Projects/opcap/internal/bundle/.test/certified-operators/operators/acc-operator /Users/mkong/Projects/opcap/internal/bundle/data.test/certified-operators/operators
touch README.md
git add .
git commit -m "fake repo initial commit"

git remote add origin /Users/mkong/Projects/opcap/internal/bundle/data.test/test-data.git
git push -u origin master
# URL=`git config --get remote.origin.url`
# echo $URL
git config --get remote.origin.url
