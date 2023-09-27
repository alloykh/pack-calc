package models

import "errors"

type PacksResponse struct {
	Total  int         `json:"total"`
	Result map[int]int `json:"result"`
}

type PacksRequest struct {
	Amount int `json:"amount"`
}

func (r PacksRequest) Validate() (err error) {

	if r.Amount <= 0 {
		return errors.New("input must be greater than 0")
	}

	return
}
