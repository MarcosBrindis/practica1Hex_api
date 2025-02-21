package model

type FilmHub struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Type        string  `json:"type"`
	Genre       string  `json:"genre"`
	Duration    float32 `json:"duration"`
	ReleaseYear int     `json:"release_year"`
}
