#!/bin/bash
set -e

bash go-mod-hack.sh v1.20.1
go get github.com/argoproj/argo-cd@a4ee25b
touch hack/.hack.argocd.installed
