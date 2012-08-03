#!/bin/bash

FILE='/proc/net/dev'

cat $FILE |awk '
BEGIN {printf "%s %s %s\n","adapter", "receive", "transmit(bytes)"};
/eth/ {printf "%s %s %s\n", $1, $2, $10}'
