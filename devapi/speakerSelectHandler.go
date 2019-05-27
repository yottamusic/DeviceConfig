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

// Speaker Struct for parsing the request as SpeakerID
type Speaker struct {
	SpeakerID string `json:"speakerID"`
	Mode      string `json:"mode"`
}

// SpeakerSelectHandler Handler for Selecting the Speaker
func SpeakerSelectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Expected a POST Request:", http.StatusBadRequest)
		log.Printf("Expected a POST Request")
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
	fmt.Printf("SpeakerID is: %s and Mode is: %s", spkr.SpeakerID, spkr.Mode)

	if runtime.GOOS == "windows" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{" + " \"result\": \"error\", \"message\": \"Cannot Execute on a Windows Machine\" " + "}"))
	} else {
		// Select speaker based on mode
		if spkr.Mode == "mono" {
			execOutput, err := devexec.SelectMonoSpeaker(spkr.SpeakerID, spkr.Mode)
			if err != nil {
				// Got Failure in Executing Command
				log.Printf("Got Failure in Executing Command: %v", err)
				http.Error(w, "Got Failure in Executing Command. "+err.Error(), http.StatusBadRequest) //HTTP 400
				return
			} else if execOutput == "failure" {
				// Got Failure in Executing Command
				log.Printf("Failed to Select Speaker!")
				http.Error(w, "Failed to Select Speaker!", http.StatusBadRequest) //HTTP 400
				return
			}
			w.Write([]byte(execOutput)) // Send Response
		} else if spkr.Mode == "left" {
			execOutput, err := devexec.SelectLeftSpeaker(spkr.SpeakerID, spkr.Mode)
			if err != nil {
				// Got Failure in Executing Command
				log.Printf("Got Failure in Executing Command: %v", err)
				http.Error(w, "Got Failure in Executing Command. "+err.Error(), http.StatusBadRequest) //HTTP 400
				return
			} else if execOutput == "failure" {
				// Got Failure in Executing Command
				log.Printf("Failed to Select Speaker!")
				http.Error(w, "Failed to Select Speaker!", http.StatusBadRequest) //HTTP 400
				return
			}
			w.Write([]byte(execOutput)) // Send Response
		} else if spkr.Mode == "right" {
			execOutput, err := devexec.SelectRightSpeaker(spkr.SpeakerID, spkr.Mode)
			if err != nil {
				// Got Failure in Executing Command
				log.Printf("Got Failure in Executing Command: %v", err)
				http.Error(w, "Got Failure in Executing Command. "+err.Error(), http.StatusBadRequest) //HTTP 400
				return
			} else if execOutput == "failure" {
				// Got Failure in Executing Command
				log.Printf("Failed to Select Speaker!")
				http.Error(w, "Failed to Select Speaker!", http.StatusBadRequest) //HTTP 400
				return
			}
			w.Write([]byte(execOutput)) // Send Response
		}
	}
	return
}
