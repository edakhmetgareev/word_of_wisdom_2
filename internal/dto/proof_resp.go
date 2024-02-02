package dto

type ProofResp struct {
	Proof string `json:"proof"`
}

func (p ProofResp) IsValid() bool {
	return p.Proof != ""
}
