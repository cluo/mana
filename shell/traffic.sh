#!/bin/bash

FILE='/proc/net/dev'

cat $FILE |awk '
BEGIN {
    printf "%s\t%s\t%s\n","interface", "receive", "transmit"
};
/eth/ {
    printf "%s\t\t%s\t%s\n", $1, $2, $10
}'
