package incident

import "time"

type Summary struct {
	ID        string    `json:"id"`
	Summary   string    `json:"summary"`
	Status    string    `json:"status"`
	CreatedBy string    `json:"createdBy"`
	CreatedOn time.Time `json:"createdOn"`
}
