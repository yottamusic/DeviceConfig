package devapi

import (
	"log"
	"net/http"
	"runtime"

	"github.com/yottamusic/DeviceConfig/devexec"
)

// SpeakerSelectHandler Handler for Selecting the Speaker
func SpeakerSelectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Expected a GET Request:", http.StatusBadRequest)
		log.Printf("Expected a GET Request")
		return
	}

	if runtime.GOOS == "windows" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{" + " \"result\": \"error\", \"message\": \"Cannot Execute on a Windows Machine\" " + "}"))
	} else {
		//
		execOutput, err := devexec.SelectSpeaker()
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
	return
}
