#!/bin/bash
KAMAL_LOG=/etc/kamal/log
sudo setsid kamal-server > $KAMAL_LOG 2>&1 &

