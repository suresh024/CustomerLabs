package models

type WebHookRequest struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	UserTraits      map[string]Attribute `json:"user_traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type ErrResponse struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

type RequestBody map[string]string
