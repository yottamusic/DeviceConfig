package devexec

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

// BashResponseSpeaker Struct for parsing the Request as todo, result, message getting Speakers Count/List
type BashResponseSpeaker struct {
	ToDo    string `json:"todo"`
	Result  string `json:"result"`
	Message string `json:"message"`
}

// GetSpeakersList Get Speakers List by Running a Script from Bash
func GetSpeakersList() (string, error) {

	command := exec.Command("/bin/bash", "Check_Availible_Speakers.sh", "-l")
	command.Dir = "/root/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	var bashResponseSpeaker BashResponseSpeaker
	err = json.Unmarshal(commandOutput, &bashResponseSpeaker)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		return "", err
	}
	//output := string(commandOutput[:])
	log.Println(string(commandOutput[:]))
	if bashResponseSpeaker.Result != "success" {
		return bashResponseSpeaker.Result, nil
	}
	return bashResponseSpeaker.Message, nil
}
