package query

import "time"

type MeowResponse struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedOn time.Time `json:"created_on"`
}
