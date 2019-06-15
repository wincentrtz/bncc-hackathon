package models

// Album models
type Album struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"desc"`
	Users       []*User    `json:"users"`
	Payments    []*Payment `json:"payments"`
	IsPaidOff   bool       `json:"is_paid_off"`
	ImagePaths  []string   `json:"image_paths"`
}
