#!/bin/sh

if [ $# != 2 ]; then
  exit 0
fi

name=`iw wlan0 scan | grep SSID | grep $2`
if [ "$name" = "" ]; then
  exit 0
fi

if [ "$1" = "Mono" ]; then
  # other commands herer
  echo Mono for $name OK
  exit 0
fi

if [ "$1" = "Stereo" ]; then
  # other commands herer
  echo Stereo for $name OK
  exit 0
fi


