package devapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	"github.com/yottamusic/DeviceConfig/devexec"
)

// SpeakerConfigHandler Handler for Configuring the Speaker into Mono/Stereo mode
func SpeakerConfigHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Expected a GET Request:", http.StatusBadRequest)
		log.Printf("Expected a GET Request")
		return
	}
	jsonMsg, err := ioutil.ReadAll(r.Body)
	//log.Printf(string(jsonMsg))

	defer r.Body.Close()
	if err != nil {
		log.Printf("Error Reading Message Body: %v", err)
		http.Error(w, "Can't Read Message Body. "+err.Error(), http.StatusBadRequest) //HTTP 400
		return
	}

	var spkr Speaker
	err = json.Unmarshal(jsonMsg, &spkr)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		http.Error(w, "Error Parsing Message into JSON. "+err.Error(), http.StatusBadRequest) //HTTP 400
		return
	}
	fmt.Printf("Mode: %s & SpeakerID: %s ", spkr.Mode, spkr.SpeakerID)

	if runtime.GOOS == "windows" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{" + " \"result\": \"error\", \"message\": \"Cannot Execute on a Windows Machine\" " + "}"))
	} else {
		// Set Speaker(s) is required Modes
		if spkr.Mode == "mono" {
			execOutput, err := devexec.ConfigMonoMode(spkr.SpeakerID, spkr.Mode)
			if err != nil {
				// Got Failure in Executing Command
				log.Printf("Got Failure in Executing Command: %v", err)
				http.Error(w, "Got Failure in Executing Command. "+err.Error(), http.StatusBadRequest) //HTTP 400
				return
			} else if execOutput == "failure" {
				// Got Failure in Executing Command
				log.Printf("Failed to Configure Speaker!")
				http.Error(w, "Failed to Configure Speaker!", http.StatusBadRequest) //HTTP 400
				return
			}
			w.Write([]byte(execOutput)) // Send Response
		} else if spkr.Mode == "stereo" {
			execOutput, err := devexec.ConfigStereoMode(spkr.SpeakerID, spkr.Mode)
			if err != nil {
				// Got Failure in Executing Command
				log.Printf("Got Failure in Executing Command: %v", err)
				http.Error(w, "Got Failure in Executing Command. "+err.Error(), http.StatusBadRequest) //HTTP 400
				return
			} else if execOutput == "failure" {
				// Got Failure in Executing Command
				log.Printf("Failed to Configure Speaker!")
				http.Error(w, "Failed to Configure Speaker!", http.StatusBadRequest) //HTTP 400
				return
			}
			w.Write([]byte(execOutput)) // Send Response
		}
	}
	return
}
