Check_Availible_Speaker.sh

Check_Availible_Speaker.sh -c to get speaker country
Check_Availible_Speaker.sh -l to get speaker list

Check_Wifi_Mode.sh will return STA/AP as the WIFI mode

Set_LED.sh $1 $2 $3
$1 is the speaker_serial_number
$2 is the LED color, 0x20 for red, 0x21 for green 0x22 for blue
$3 is the LED mode
# bit 0-2:
#   1: fast blinking
#   2: slow blinking
#   3: breathing
#   4: on
#   5: off
# bit 7 will be polarity control

Set_Speaker.sh $1 $2
$1 is the speaker_serial_number name
$2 is the speaker_serial_number mode

# to play a default sound in specific speaker
Play_Sound.sh $1
$1 is the serial number of speaker SSID

Wifi_Connect.sh SSID PIN    -> connect to SSID using PIN