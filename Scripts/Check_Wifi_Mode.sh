#!/bin/sh

cnt=`ps -aux | grep hostapd | wc -l`

if [ $cnt -gt 1 ]; then
  echo AP
else
  echo STA
fi
