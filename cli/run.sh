#!/bin/bash -x
go build
./cli --help
./cli --add test.rac.reckhard.ca localhost:80
./cli --add loop.rac.reckhard.ca localhost:8888
./cli --add localhost loop.rac.reckhard.ca:8888
./cli --add test.rac.reckhard.ca localhost:81
./cli --rm test.rac.reckhard.ca localhost:81
./cli --run prog.rac.reckhard.ca 8080 ./binary arg1 arg2
rm cli
