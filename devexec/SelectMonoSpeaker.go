package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// SelectMonoSpeaker Select One Speaker by Blink LED and create Sound
func SelectMonoSpeaker(speakerID string, mode string) (string, error) {

	commandSound := exec.Command("/bin/bash", "Play_Sound.sh", speakerID)
	commandSound.Dir = "/root/DeviceScripts/"
	commandSoundOutput, err := commandSound.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerSoundOutput := string(commandSoundOutput[:])
	log.Println(speakerSoundOutput)

	commandLED := exec.Command("/bin/bash", "Set_Mono_Speaker_LED.sh", speakerID)
	commandLED.Dir = "/root/DeviceScripts/"
	commandLEDOutput, err := commandLED.CombinedOutput()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	speakerLEDOutput := string(commandLEDOutput[:])
	log.Println(speakerLEDOutput)
	return "{\"result\" : \"Success\"}", nil
}
