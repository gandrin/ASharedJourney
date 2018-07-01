#!/bin/bash

DESTINATION="/tmp/ASharedJourney_Windows"

make build_assets
GOOS=windows GOARCH=386 go build -o ASharedJourney.exe main.go
chmod +x main
mkdir -p $DESTINATION
cp main $DESTINATION/ASharedJourney.exe
cp -r ./assets $DESTINATION/

echo "Built!"
echo "You may now send" $DESTINATION "to your beautiful friends!"
