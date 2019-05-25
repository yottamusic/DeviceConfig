#!/bin/sh

SpkName=YottaSpk

if [ "$1" = "" ]; then
  cnt=`iw wlan0 scan | grep SSID | grep $SpkName | wc -l`
  echo $cnt
else
  iw wlan0 scan | grep SSID | grep $SpkName | awk '{print $2";"}'
fi
