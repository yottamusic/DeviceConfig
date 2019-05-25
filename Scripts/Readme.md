Check_Availible_Speaker.sh

Check_Availible_Speaker.sh -c to get speaker country
Check_Availible_Speaker.sh -l to get speaker list

Check_Wifi_Mode.sh will return STA/AP as the WIFI mode

Set_Left_Speaker_LED.sh $1
Set_Mono_Speaker_LED.sh $1
Set_Right_Speaker_LED.sh $1
$1 is the speaker_serial_number

Set_Speaker.sh $1 $2
$1 is the speaker_serial_number name
$2 is the speaker_serial_number mode

# to play a default sound in specific speaker
Play_Sound.sh $1
$1 is the serial number of speaker SSID

Wifi_Connect.sh SSID PIN    -> connect to SSID using PIN