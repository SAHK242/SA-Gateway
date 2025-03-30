#!/usr/bin/env bash

echo -n "Source branch: "
read -r FROM

echo -n "Target branch: "
read -r TO

echo -e "$(printf "Merging & Push to branch %s from branch %s" "$TO" "$FROM")"

git checkout "${TO}"
git pull origin "${TO}"
git merge "${FROM}"
git push
git checkout "${FROM}"