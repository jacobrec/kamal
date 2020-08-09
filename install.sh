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

mv server/a.out /bin/kamal-server
mv cli/a.out /bin/kamal
mkdir -p /etc/kamal

echo $PWD
