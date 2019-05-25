#!/bin/sh

if [ $# != 1 ]; then
  exit 0
fi

if [ -f $1 ]; then
  /root/playfile $1
fi


