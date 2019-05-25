#!/bin/sh

if [ $# != 1 ]; then
  exit 0
fi

# $1 is the speaker serial number in the middle of the SSID
/root/i2crw0 w 33 40 $1

