#!/bin/sh
#rmmod 88x2bu
#insmod /work/88x2bu.ko
ifconfig wlan0 up
ssid=$1
passwd=$2
ifconfig wlan0 up
pkill wpa_supplicant
`mkdir /etc/wpa_supplicant/`
echo "network={\n    ssid=\"$ssid\"\n    psk=\"$passwd\"\n    priority=1\n}"   >/etc/wpa_supplicant/wpa_supplicant.conf 
wpa_supplicant -i wlan0 -c /etc/wpa_supplicant/wpa_supplicant.conf  &
sleep 3
dhclient wlan0

