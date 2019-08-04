package devapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// EQ Struct for parsing the request as EQ Settings
type EQ struct {
	EQSettings string `json:"settings"`
}

// SaveConfigHandler Handler for Saving the EQ Settings
func SaveConfigHandler(w http.ResponseWriter, r *http.Request) {
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

	var eqConfig EQ
	err = json.Unmarshal(jsonMsg, &eqConfig)
	if err != nil {
		log.Printf("Error Parsing Message into JSON: %v", err)
		http.Error(w, "Error Parsing Message into JSON. "+err.Error(), http.StatusBadRequest) //HTTP 400
		return
	}
	fmt.Printf("Settings: %s", eqConfig.EQSettings)
	ioutil.WriteFile("./config", jsonMsg, 0600)

	w.Write([]byte("{" + " \"result\": \"Saved EQ Settings Successfully\" " + "}"))
	return
}
