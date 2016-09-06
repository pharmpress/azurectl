#!/bin/bash

version=$(grep -E -o "[0-9]+\.[0-9]+\.[0-9]+\+git" ./version/version.go 
)

if [[ "$version" =~ ^(.*\.)([0-9]+)\+git$ ]];
then
	inc=$((BASH_REMATCH[2]+1))
	currentversion="${BASH_REMATCH[1]}${BASH_REMATCH[2]}";
	nextversion="${BASH_REMATCH[1]}$inc";

	git tag -a "v$currentversion" -m "Version $currentversion"
	git push origin "v$currentversion"
	sed -i -e "s/${currentversion}/${nextversion}/" ./version/version.go
	git add ./version/version.go
	git commit -m "Setting version to ${nextversion}"
	git push
else
	echo "Version wrong format"
	exit 1
fi
