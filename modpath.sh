#!/bin/bash
currentPath=$(pwd)
currentFolder=${currentPath##*/}
modPath=${currentPath:18:${#currentPath}}
go mod init $currentFolder
sed -i "s+$currentFolder+${modPath}+g" go.mod
