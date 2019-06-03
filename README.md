# DeviceConfig
Web Application for Device &amp; Speaker Configuration

## API Documentation

###Get WiFi Mode
curl -X GET http://192.168.1.24/wifi/api/v1/mode

###Get SSID List 
curl -X GET http://192.168.1.24/wifi/api/v1/list

###WiFi Connection
curl -X POST http://192.168.1.24/wifi/api/v1/connect -d '{"ssid": "myWiFi", "password": "88888888"}'

###Get Speakers List 
curl -X GET http://192.168.1.24/speakers/api/v1/list
  
###Select Speakers
curl -X POST http://192.168.1.24/speakers/api/v1/select -d '{"speakerID": "123", "mode": "mono"}'
Note:
>For Mono Configuration:      {"speakerID": "123", "mode": "mono"}
>For Stereo Configuration:    {"speakerID": "123", "mode": "left"}
>For Stereo Configuration:    {"speakerID": "123", "mode": "right"}

###Configure Speakers
curl -X POST http://192.168.1.24/speakers/api/v1/config -d '{"speakerID": "123,456", "mode": "stereo"}'
Note: 
>For Mono Configuration:      {"speakerID": "123", "mode": "mono"}
>For Stereo Configuration:    {"speakerID": "123,456", "mode": "stereo"}