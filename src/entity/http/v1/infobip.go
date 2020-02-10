package v1

// InfobipRequestPayload ...
type InfobipRequestPayload struct {
	Messages []InfobipMessages `json:"messages"`
}

// InfobipMessages .../
type InfobipMessages struct {
	From             string               `json:"from"`
	Destinations     []InfobipDestination `json:"destinations"`
	Text             string               `json:"text"`
	NotifyURL        string               `json:"notifyUrl"`
	NotifyContenType string               `json:"notifyContentType"`
	CallbackData     string               `json:"callbackData"`
}

// InfobipDestination ...
type InfobipDestination struct {
	To string `json:"to"`
}

// InfobipResponse ...
type InfobipResponse struct {
	Messages []struct {
		To     string `json:"to"`
		Status struct {
			GroupID     uint   `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          uint   `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
		MessageID string `json:"messageId"`
	} `json:"messages"`
}

// InfobipCallBackRequest ...
type InfobipCallBackRequest struct {
	Results []InfobipRequestChild
}

// InfobipRequestChild ...
type InfobipRequestChild struct {
	BulkID     string
	MessagesID string
	To         string
	SentAt     string
	DoneAt     string
	SmsCount   uint
	MccMnc     string
	Price      struct {
		PricePerMessages string
		Currency         string
	}
	Status struct {
		GroupID     string
		GroupName   string
		ID          string
		Name        string
		Description string
	}
	Error struct {
		GroupID     string
		GroupName   string
		ID          string
		Name        string
		Description string
		Permanent   bool
	}
	CallBackData string
}
