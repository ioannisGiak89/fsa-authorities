package model

// CustomResponse is used to keep the status code in FSA client
type CustomResponse struct {
	StatusCode   int
	ResponseBody []byte
}

// AuthoritiesResponse is a FSA response for /Authorities/basic
type AuthoritiesResponse struct {
	Authorities []Authority
	Meta        Meta
	Links       []Link
}

// EstablishmentsResponse is a FSA response for /Establishments
type EstablishmentsResponse struct {
	Establishments []Establishment
	Meta           Meta
	Links          []Link
}

// Authority as returns from FSA's API
type Authority struct {
	LocalAuthorityId     int
	LocalAuthorityIdCode string
	Name                 string
	EstablishmentCount   int
	SchemeType           SchemeType
	Links                []Link
}

// Establishment as returns from FSA's API
type Establishment struct {
	FhrsId                   int
	LocalAuthorityBusinessID string
	BusinessName             string
	BusinessType             string
	BusinessTypeID           int
	AddressLine1             string
	AddressLine2             string
	AddressLine3             string
	AddressLine4             string
	PostCode                 string
	Phone                    string
	RatingValue              string
	RatingKey                string
	RatingDate               string
	LocalAuthorityCode       string
	LocalAuthorityName       string
	LocalAuthorityWebSite    string
	Scores                   Scores
	SchemeType               string
	Geocode                  Geocode
	RightToReply             string
	Distance                 string
	NewRatingPending         bool
	Meta                     Meta
	Links                    []Link
}

// Geocode as returns from FSA's API
type Geocode struct {
	Longitude string
	Latitude  string
}

// Scores as returns from FSA's API
type Scores struct {
	Hygiene                int
	Structural             int
	ConfidenceInManagement int
}

// Link as returns from FSA's API
type Link struct {
	Rel  string
	Href string
}

// Meta as returns from FSA's API
type Meta struct {
	DataSource  string
	ExtractDate string
	ItemCount   int
	ReturnCode  string
	TotalCount  int
	TotalPages  int
	PageSize    int
	PageNumber  int
}
