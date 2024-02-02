package dto

type QuoteResp struct {
	Quote        string `json:"quote,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}
