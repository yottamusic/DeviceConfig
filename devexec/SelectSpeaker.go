package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// SelectSpeaker Select Speaker by Blink LED and create Sound
func SelectSpeaker() (string, error) {

	command := exec.Command("/bin/bash", "Play_Sound.sh")
	command.Dir = "/root/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerOutput := string(commandOutput[:])
	log.Println(speakerOutput)
	return speakerOutput, nil
}
