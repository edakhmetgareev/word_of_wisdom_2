package dto

type ChallengeResp struct {
	Challenge    string `json:"challenge"`
	LeadingZeros int    `json:"leading_zeros"`
}

func (c ChallengeResp) IsValid() bool {
	return c.Challenge != "" && c.LeadingZeros > 0
}
