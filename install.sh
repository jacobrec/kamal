#!/bin/bash
SCRIPT=`realpath $0`
SCRIPTPATH=`dirname $SCRIPT`
cd $SCRIPTPATH

cd server
go build -o a.out
cd ../

cd cli
go build -o a.out
cd ../

sudo mv server/a.out /bin/kamal-server
sudo mv cli/a.out /bin/kamal
sudo mkdir -p /etc/kamal

echo $PWD
