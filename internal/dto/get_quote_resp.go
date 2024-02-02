package dto

type QuoteResp struct {
	Quote        string `json:"quote,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

func (q QuoteResp) IsValid() bool {
	return q.Quote != "" && q.ErrorMessage == ""
}
