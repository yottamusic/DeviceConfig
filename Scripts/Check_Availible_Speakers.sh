#!/bin/sh

SpkName=YottaSpk

if [ "$1" = "-c" ]; then
  cnt=`iw wlan0 scan | grep SSID | grep $SpkName | wc -l`
  printf '{"todo": "Check speaker count","result": "success","message": "'$cnt'"}\n'
  exit 0
fi

if [ "$1" = "-l" ]; then
  printf '{"todo": "Check speaker list","result": "success","message": "{'
  iw wlan0 scan | grep SSID | grep $SpkName | awk '{printf "\""$2"\" "}'
  #iw wlan0 scan | grep SSID | awk '{printf "\""$2"\" "}'
  printf '}""}\n'
  exit 0
fi

#{"todo": "Check speaker count","result": "success","message": "YottaSpk_123_Left,...."}