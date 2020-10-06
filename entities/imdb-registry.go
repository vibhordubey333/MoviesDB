package entities

type IMDBRegistry struct {
	MovieName   string     `json:"moviename,omitempty"`
	Ratings     []Ratings  `json:"ratings,omitempty"`
	RatingCount int        `json:"peoplecount,omitempty"`
	Comments    []Comments `json:"comments,omitempty"`
}

type Comments struct {
	MovieNameCom string `json:"movienamecom,omitempty"`
	UserName     string `json:"username,omitempty"`
	Comment      string `json:"comment,omitempty"`
}

type Ratings struct {
	MovieName string  `json:"moviename,omitempty"`
	UserName  string  `json:"username,omitempty"`
	Rating    float64 `json:"rating,omitempty"`
}
