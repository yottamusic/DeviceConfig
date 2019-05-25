#!/bin/sh

if [ $# != 1 ]; then
  exit 0
fi

# 1: blinking
# 2: on
# 3: off

if [ $1 -le 3 -a $1 -ge 1 ]; then
  /root/i2crw0 w 33 20 $1
fi


