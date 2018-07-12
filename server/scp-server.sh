#!/bin/sh

server=$1

echo "building server..."
go build -o web server.go
echo "server done.\n"

echo "building deploy-server..."
go build -o deploy deploy-server/a.go
echo "deploy-server done.\n"

echo "scp..."
scp web deploy deploy.sh $1

rm web
rm deploy

echo "\ndone.\n"

