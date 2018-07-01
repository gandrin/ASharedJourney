#!/bin/bash

go build main.go
chmod +x main
cp main ./ASharedJourney.app/Contents/MacOS/ASharedJourney
cp -r music ./ASharedJourney.app/Contents/MacOS/
cp -r tiles ./ASharedJourney.app/Contents/MacOS/
