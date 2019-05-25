package devexec

import (
	"fmt"
	"log"
	"os/exec"
)

// GetSpeakersList Get Speakers List by Running a Script from Bash
func GetSpeakersList() (string, error) {

	out, err := exec.Command("bash", "-c", "pwd").Output()
	if err != nil {
		fmt.Printf("%s", err)
		return "", err
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	log.Println(output)
	return output, nil
}
