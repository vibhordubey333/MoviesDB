package entities

type IMDBRegistry struct {
	MovieName   string                 `json:"moviename,omitempty"`
	Genre       string                 `json:"genre,omitempty"`
	Rating      string                 `json:"rating,omitempty"`
	PeopleCount int                    `json:"peoplecount,omitempty"`
	Comments    map[string]interface{} `json:"comments,omitempty"`
}