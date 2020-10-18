package movie

// Movie ...
type Movie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type (
	// MovieSearchSpec ...
	MovieSearchSpec struct {
		Keyword string
		Page    int64
	}

	// MovieSearchResponse ...
	MovieSearchResponse struct {
		Search       []Movie `json:"Search"`
		TotalResults string  `json:"totalResults"`
		Response     string  `json:"Response"`
	}
)

type UseCase interface {
	SearchMovie(spec MovieSearchSpec) (res MovieSearchResponse, err error)
}
