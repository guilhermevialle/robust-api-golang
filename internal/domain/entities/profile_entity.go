package entities

type Profile struct {
	ID         string `json:"id"`
	CustomerID string `json:"-"`
	Name       string `json:"name"`
	Summary    string `json:"summary"`
}
