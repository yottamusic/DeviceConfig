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

// Connect Struct for parsing the request as SpeakerID
type Connect struct {
	SSID     string `json:"ssid"`
	Password string `json:"password"`
}

// WifiConnectionHandler Handler for Configuring the Speaker into Mono/Stereo mode
func WifiConnectionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "{"+" \"result\": \"error\", \"message\": \"Expected a POST Request\"}", http.StatusBadRequest)
		log.Printf("Expected a POST Request")
		return
	}
	jsonMsg, err := ioutil.ReadAll(r.Body)
	//log.Printf(string(jsonMsg))

	defer r.Body.Close()
	if err != nil {
		log.Printf("Error Reading Message Body: %v", err)
		http.Error(w, "{"+" \"result\": \"error\", \"message\": \"Can't Read Message Body\""+err.Error()+"}", http.StatusBadRequest) //HTTP 400
		return
	}

	var connect Connect
	err = json.Unmarshal(jsonMsg, &connect)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		http.Error(w, "{"+" \"result\": \"error\", \"message\": \"Error Parsing Message into JSON\""+err.Error()+"}", http.StatusBadRequest) //HTTP 400
		return
	}
	fmt.Printf("SSID: %s & Password: %s ", connect.SSID, connect.Password)

	if runtime.GOOS == "windows" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte("{" + " \"result\": \"error\", \"message\": \"Cannot Execute on a Windows Machine\" " + "}"))
	} else {
		// Set Speaker(s) is required Modes
		execOutput, err := devexec.ConnectWiFi(connect.SSID, connect.Password)
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
	return
}
