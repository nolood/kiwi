package dto

type MatchesPosition struct {
	Alternatenames []*struct {
		Length int
		Start  int
	}
}

type City struct {
	Alternatenames  string
	CountryCode     string
	Longitude       float64
	Latitude        float64
	Name            string
	ID              int64
	GeneralMatch    string
	MatchesPosition MatchesPosition `json:"_matchesPosition"`
}
