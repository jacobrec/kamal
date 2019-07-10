#!/bin/bash -x
export CONFIG="../.config"
go build
rm -f ../.config
./cli --add test.rac.reckhard.ca localhost:80
cat ../.config
echo ""
./cli --add loop.rac.reckhard.ca localhost:8888
cat ../.config
echo ""
./cli --add localhost loop.rac.reckhard.ca:8888
cat ../.config
echo ""
./cli --add test.rac.reckhard.ca localhost:81
cat ../.config
echo ""
./cli --rm test.rac.reckhard.ca localhost:81
cat ../.config
echo ""
./cli --run prog.rac.reckhard.ca 8080 ./binary arg1 arg2
rm cli
