package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// ConfigSpeakers Config Speakers in Mono/Stereo Mode
func ConfigSpeakers() (string, error) {

	command := exec.Command("/bin/bash", "Set_Speaker.sh")
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
