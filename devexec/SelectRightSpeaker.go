package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// SelectRightSpeaker Select Right Speaker by Blink LED and create Sound
func SelectRightSpeaker(speakerID string, mode string) (string, error) {

	commandSound := exec.Command("/bin/bash", "Play_Sound.sh", speakerID)
	commandSound.Dir = "/root/"
	commandSoundOutput, err := commandSound.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerSoundOutput := string(commandSoundOutput[:])
	log.Println(speakerSoundOutput)

	commandLED := exec.Command("/bin/bash", "Set_Right_Speaker_LED.sh", speakerID)
	commandLED.Dir = "/root/"
	commandLEDOutput, err := commandLED.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerLEDOutput := string(commandLEDOutput[:])
	log.Println(speakerLEDOutput)
	return "{\"result\" : \"Success\"}", nil
}
