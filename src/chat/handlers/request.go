package handlers

type SendNewMessageRequest struct {
	Message string `json:"message"`
}

type RetrievePreviousMessageRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
