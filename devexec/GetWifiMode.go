package devexec

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// BashResponseWifiMode Struct for parsing the Request as todo, result, message getting Wifi Mode
type BashResponseWifiMode struct {
	ToDo    string `json:"todo"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

// GetWifiMode Get Wifi Mode by Running a Script from Bash
func GetWifiMode() (string, error) {

	command := exec.Command("/bin/bash", "Check_Wifi_Mode.sh", "-l")
	command.Dir = "/root/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	//output := string(commandOutput[:])
	log.Println(string(commandOutput[:]))
	var bashResponseWifiMode BashResponseWifiMode
	err = json.Unmarshal(commandOutput, &bashResponseWifiMode)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		return "", err
	}
	if bashResponseWifiMode.Result != "success" {
		return bashResponseWifiMode.Result, nil
	}
	return bashResponseWifiMode.Message, nil
}
