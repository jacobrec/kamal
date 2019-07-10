#!/bin/bash -x
export CONFIG="../.config"
go build
rm -f ../.config
./cli --add test.rac.reckhard.ca localhost
echo ""
./cli --add loop.rac.reckhard.ca localhost:8888
echo ""
./cli --add localhost loop.rac.reckhard.ca:8888
echo ""
./cli --add test.rac.reckhard.ca localhost:81
echo ""
./cli --rm test.rac.reckhard.ca localhost:81
echo ""
./cli --ls
echo ""
./cli --run prog.rac.reckhard.ca 8080 ./binary arg1 arg2
rm cli
