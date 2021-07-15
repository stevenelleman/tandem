#!/bin/bash

# 1. Get master commit
MASTER_SHA=$( cat .git/refs/heads/master )

# Commits in order of earliest to latest
SHAs=()

# 2. Gather SHAs of unique commits, going back to masterSHA
INDEX=0
while [ true ]
do
  COMMIT=$( git rev-parse HEAD~$INDEX )
  if [ $COMMIT == $MASTER_SHA ]; then
    break
  else
    SHAs=($COMMIT ${SHAs[@]})
    # echo $commit $index
    # echo "${SHAs[@]}"
    INDEX=$((INDEX+1))
  fi
done

# 3. make branch name
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO=$(basename $(dirname $DIR))
if [ $REPO == "tandem" ]; then
  echo "Already within Tandem -- directly make branch against master."
  exit
fi
BRANCH=$(git branch --show-current)
NEW_BRANCH=$REPO/$BRANCH

# 4. Need to checkout tandem as base and cherry-pick commits onto it
git remote add upstream git@github.com:stevenelleman/tandem.git
git fetch upstream
git checkout -b $NEW_BRANCH upstream/$NEW_BRANCH
for COMMIT in SHAs
do
  git cherry-pick $COMMIT
done



