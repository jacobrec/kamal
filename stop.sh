#!/bin/bash
# For development only
kill $(ps -aux | ag kamal-server | ag root | tail -n 1 | ag '\d+' --silent -o -m 1)
