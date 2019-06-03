package devexec

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// ConfigStereoMode Config Speakers in Mono Mode
func ConfigStereoMode(speakerID string, mode string) (string, error) {

	speakerList := strings.FieldsFunc(speakerID, func(r rune) bool {
		if r == ',' {
			return true
		}
		return false
	})

	commandSpkr1 := exec.Command("/bin/bash", "Set_Speaker.sh", speakerList[0], mode)
	commandSpkr1.Dir = "/root/DeviceScripts/"
	commandSpkr1Output, err := commandSpkr1.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speaker1Output := string(commandSpkr1Output[:])
	log.Println(speaker1Output)

	commandSpkr2 := exec.Command("/bin/bash", "Set_Speaker.sh", speakerList[1], mode)
	commandSpkr2.Dir = "/root/DeviceScripts/"
	commandSpkr2Output, err := commandSpkr2.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speaker2Output := string(commandSpkr2Output[:])
	log.Println(speaker2Output)

	return "{\"result\" : \"Success\"}", nil
}
