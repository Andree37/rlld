package models

type Shorter struct {
	Url string `json:"url" binding:"required"`
}

// Shorter logic here
func (s *Shorter) Shortens() string {
	return s.Url + "potato"
}
