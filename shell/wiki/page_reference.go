package wiki

import "time"

type PageReference struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	ShortLink string    `json:"shortLink"`
	CreatedBy string    `json:"createdBy"`
	CreatedOn time.Time `json:"createdOn"`
}
