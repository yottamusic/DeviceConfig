package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// ConfigMonoMode Config Speakers in Mono Mode
func ConfigMonoMode(speakerID string, mode string) (string, error) {

	command := exec.Command("/bin/bash", "Set_Speaker.sh", speakerID, mode)
	command.Dir = "/root/DeviceScripts/"
	commandOutput, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerOutput := string(commandOutput[:])
	log.Println(speakerOutput)
	return "{\"result\" : \"Success\"}", nil
}
