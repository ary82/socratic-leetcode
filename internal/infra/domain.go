package infra

import "fmt"

type TextGenerationRequest struct {
	Question string `json:"question"`
	Lines    int    `json:"lines"`
	Lang     string `json:"lang"`
	Code     string `json:"code"`
}

type TextGenerationResponse struct {
	Line int    `json:"line"`
	Q    string `json:"q"`
}

func (req *TextGenerationRequest) Validate() error {
	if req.Question == "" {
		return fmt.Errorf("no question given")
	}
	if req.Lines == 0 {
		return fmt.Errorf("0 lines solution not allowed")
	}
	if req.Lang == "" {
		return fmt.Errorf("lang cannot be empty")
	}
	if req.Code == "" {
		return fmt.Errorf("no code input given")
	}

	return nil
}
