#!/bin/bash

DESTINATION="./build/ASharedJourney.app"

make build_assets
go build main.go
chmod +x main
mkdir -p $DESTINATION/Contents/MacOS/
cp main $DESTINATION/Contents/MacOS/ASharedJourney
cp -r ./assets $DESTINATION/Contents/MacOS/

echo "Built!"
echo "You may now send" $DESTINATION "to your beautiful friends!"
