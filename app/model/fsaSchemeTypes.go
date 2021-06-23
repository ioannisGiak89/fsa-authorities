package model

type SchemeType int

const (
	FHRS SchemeType = iota + 1
	FHIS
)

func (s SchemeType) String() string {
	return [...]string{"FHRS", "FHIS"}[s-1]
}
