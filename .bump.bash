#!/usr/bin/env bash

set -x

git describe --abbrev=0 --tags &>/dev/null || git tag -am "first tag" "0.0.1"

VERSION=($git describe --abbrev=0 --tags)
VERSION_BITS=(${VERSION//./ })
MAJOR=${VERSION_BITS[0]}
MINOR=${VERSION_BITS[1]}
PATCH=${VERSION_BITS[2]}

case $1 in
"major")
	((MAJOR++))
	MINOR=0
	PATCH=0
	;;
"minor")
	((MINOR++))
	PATCH=0
	;;
"patch")
	((PATCH++))
	;;
*)
	echo "
    Usage: ./bump.sh major|minor|patch
    "
	exit 1
	;;
esac

NEW_TAG="$MAJOR.$MINOR.$PATCH"

echo "Bumping version from $VERSION to $NEW_TAG"
read -p "[enter] to confirm / [ctrl-c] to cancel" -n 1 -r
git pull --rebase
git tag -d "$NEW_TAG" &>/dev/null
git tag -a -m "release v$NEW_TAG" "$NEW_TAG"
git push origin --tags
