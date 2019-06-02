package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// ConnectWiFi Connect to WiFi using SSID and Password
func ConnectWiFi(ssid string, password string) (string, error) {

	commandConnect := exec.Command("/bin/bash", "Wifi_Connect.sh", ssid, password)
	commandConnect.Dir = "/root/DeviceScripts/"
	commandConnectOutput, err := commandConnect.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	commandOutput := string(commandConnectOutput[:])
	log.Println(commandOutput)
	return "{\"result\" : \"" + commandOutput + "\"}", nil
}
