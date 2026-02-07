package domain

type Prize struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Pic            string `json:"pic"`
	Link           string `json:"link"`
	Type           int32  `json:"type"`
	Data           string `json:"data"`
	Total          int64  `json:"total"`
	Left           int64  `json:"left"`
	IsUse          int32  `json:"is_use"`
	Probability    int64  `json:"probability"`
	ProbabilityMax int64  `json:"probability_max"`
	ProbabilityMin int64  `json:"probability_min"`
}
