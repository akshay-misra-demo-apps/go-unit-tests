#!/usr/bin/env bash

# echo ""
# echo "Before testing set environment variable SSO_ACCESS_TOKEN to valid access token like:"
# echo "\"export SSO_ACCESS_TOKEN=\"Bearer aad85ba7-e56e-3a91-8908-6906287968b4\""
# echo ""

set -e
echo "" > coverage.txt
baseDir="github.com/misraak/app"
for d in $(go list ./... | grep -v vendor); do
    go test -race -coverprofile=tmp.out -covermode=atomic $d
    if [ -f tmp.out ]; then
        echo $d
        cat tmp.out >> coverage.txt
        targetDir=".${d#$baseDir}"
        cp tmp.out $targetDir/coverage.out
        rm tmp.out
    fi
done

# go-acc -o coverage.out ./... -- -timeout 30m