package utils

import (
	"github.com/suresh024/CustomerLabs/models"
	"strings"
)

func extractAttrFromMap(key, id string, oldEvent map[string]string) string {
	if val, ok := oldEvent[key+id]; ok {
		return val
	} else {
		return ""
	}
}

func ConvertToWebHookRequest(oldEvent models.RequestBody) models.WebHookRequest {
	var webHokkRequest = models.WebHookRequest{
		Event:           oldEvent["Ev"],
		EventType:       oldEvent["Et"],
		AppID:           oldEvent["ID"],
		UserID:          oldEvent["UID"],
		MessageID:       oldEvent["MID"],
		PageTitle:       oldEvent["T"],
		PageURL:         oldEvent["P"],
		BrowserLanguage: oldEvent["L"],
		ScreenSize:      oldEvent["SC"],
		Attributes:      make(map[string]models.Attribute),
		UserTraits:      make(map[string]models.Attribute),
	}
	for key, value := range oldEvent {
		if strings.HasPrefix(key, "atrk") {
			id := key[len("atrk"):]
			webHokkRequest.Attributes[value] = models.Attribute{
				Value: extractAttrFromMap("atrv", id, oldEvent),
				Type:  extractAttrFromMap("atrt", id, oldEvent),
			}
		} else if strings.HasPrefix(key, "uatrk") {
			id := key[len("uatrk"):]
			webHokkRequest.UserTraits[value] = models.Attribute{
				Value: extractAttrFromMap("uatrv", id, oldEvent),
				Type:  extractAttrFromMap("uatrt", id, oldEvent),
			}
		}
	}
	return webHokkRequest

}
