package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/suresh024/CustomerLabs/models"
	"github.com/suresh024/CustomerLabs/utils"
	"log"
	"net/http"
)

var webhookKey = "aeaf5387-3d4f-44bf-80b5-2c5ca6adbf30"

func PostToWebHook(requestBody models.WebHookRequest) error {
	funcDesc := "PostToWebHook"
	url := "https://webhook.site/" + webhookKey

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s | %v", funcDesc, err)
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Printf("%s | POST request successfully sent.", funcDesc)
	return nil
}

func EventSender(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody models.RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		utils.ErrorResponse(w, "EventSender decode error", http.StatusBadRequest, err)
		return
	}
	go func(payload models.RequestBody) {
		webHookRequest := utils.ConvertToWebHookRequest(requestBody)
		err = PostToWebHook(webHookRequest)
		if err != nil {
			log.Printf("Error sending webhook request: %v, body: %v", err, webHookRequest)
			return
		}
	}(requestBody)

	utils.ReturnResponse(w, http.StatusOK, "Posted to WebHook")

}

func Start() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/webhook", EventSender).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}
