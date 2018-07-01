#!/bin/bash

go build main.go
chmod +x main
cp main ./ASharedJourney.app/Contents/MacOS/ASharedJourney
cp -r ./assets ./ASharedJourney.app/Contents/MacOS/
cp -r ./assets ./ASharedJourney.app/Contents/Resources
