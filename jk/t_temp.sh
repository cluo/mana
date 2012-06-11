#!/bin/sh

for i in `awk '/°C/{print $3}' sensors.txt |sed -nr \
    's/\+(\w+)\.\w°C/\1/p'`; do
    if [ $i -gt 47 ]; then
        echo Y
        break
    fi
done
