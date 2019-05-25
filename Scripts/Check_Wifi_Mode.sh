#!/bin/sh

cnt=`ps -aux | grep hostapd | wc -l`

if [ $cnt -gt 1 ]; then
  printf '{"todo": "Check WiFi Mode","result": "success","message": "AP"}\n'
else
  printf '{"todo": "Check WiFi Mode","result": "success","message": "STA"}\n'
fi
