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

	command := exec.Command("/bin/bash", "Check_Availible_Speakers.sh", "-c")
	command.Dir = "/root/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	log.Println(string(commandOutput[:]))
	var bashResponseSpeakerCount BashResponseSpeaker
	err = json.Unmarshal(commandOutput, &bashResponseSpeakerCount)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		return "", err
	}
	if bashResponseSpeakerCount.Result != "success" {
		return bashResponseSpeakerCount.Result, nil
	}
	speakersList := string(commandOutput[:])
	if bashResponseSpeakerCount.Message == "0" {
		log.Printf("No Wireless Speakers Found!")
	} else {
		// If Speaker Count Not Zero, Get the list
		command := exec.Command("/bin/bash", "Check_Availible_Speakers.sh", "-l")
		command.Dir = "/root/"
		commandOutput, err := command.CombinedOutput()
		if err != nil {
			fmt.Printf("%s", err)
			return "", err
		}
		log.Println(string(commandOutput[:]))
		var bashResponseSpeakerList BashResponseSpeaker
		err = json.Unmarshal(commandOutput, &bashResponseSpeakerList)
		if err != nil {
			log.Printf("Error Parsing Message into JSON: %v", err)
			return "", err
		}
		if bashResponseSpeakerList.Result != "success" {
			return bashResponseSpeakerList.Result, nil
		}
	}
	return speakersList, nil
}
