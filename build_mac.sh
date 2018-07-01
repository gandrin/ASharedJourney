#!/bin/bash

go build main.go
chmod +x main
mkdir -p ./build/ASharedJourney.app/Contents/MacOS/
cp main ./build/ASharedJourney.app/Contents/MacOS/ASharedJourney
cp -r ./assets ./build/ASharedJourney.app/Contents/MacOS/
