package devexec

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// BashResponseSSID Struct for parsing the Request as todo, result, message getting SSID List
type BashResponseSSID struct {
	ToDo    string `json:"todo"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

// GetSSIDList Get SSID List by Running a Script from Bash
func GetSSIDList() (string, error) {

	command := exec.Command("/bin/bash", "Scan_Available_SSID.sh")
	command.Dir = "/root/DeviceScripts/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	var bashResponseSSID BashResponseSSID
	err = json.Unmarshal(commandOutput, &bashResponseSSID)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		return "", err
	}
	output := string(commandOutput[:])
	log.Println(string(commandOutput[:]))
	if bashResponseSSID.Result != "success" {
		return bashResponseSSID.Result, nil
	}
	return output, nil
}
