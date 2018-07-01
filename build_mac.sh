#!/bin/bash

go build main.go
chmod +x main
mkdir -p ./ASharedJourney.app/Contents/MacOS/
cp main ./ASharedJourney.app/Contents/MacOS/ASharedJourney
cp -r ./assets ./ASharedJourney.app/Contents/MacOS/
