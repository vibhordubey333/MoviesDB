package entities

type IMDBRegistry struct {
	MovieName   string     `json:"moviename,omitempty"`
	Rating      string     `json:"rating,omitempty"`
	RatingCount int        `json:"peoplecount,omitempty"`
	Comments    []Comments `json:"comments,omitempty"`
}

type Comments struct {
	MovieNameCom string `json:"movienamecom,omitempty"`
	UserName     string `json:"username,omitempty"`
	Comment      string `json:"comment,omitempty"`
}
