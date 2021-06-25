package testUtils

import "github.com/ioannisGiak89/fsa-authorities/app/model"

func GetFakeAuthoritiesResponse() *model.AuthoritiesResponse {
	return &model.AuthoritiesResponse{
		Authorities: GetFakeAuthorities(),
	}
}

func GetFakeAuthorities() []model.Authority {
	return []model.Authority{
		{LocalAuthorityId: 197, Name: "Aberdeen City", SchemeType: 1},
		{LocalAuthorityId: 198, Name: "Aberdeenshire", SchemeType: 2},
		{LocalAuthorityId: 277, Name: "Adur", SchemeType: 1},
	}
}

func GetFakeAuthority() *model.Authority {
	return &model.Authority{
		LocalAuthorityId: 197,
		Name:             "Aberdeen City",
		SchemeType:       1,
	}
}

func GetFakeEstablishmentsResponse() *model.EstablishmentsResponse {
	return &model.EstablishmentsResponse{
		Establishments: GetFakeFhrsEstablishments(),
	}
}

func GetFakeFhrsEstablishments() []model.Establishment {
	return []model.Establishment{
		{RatingValue: "1", SchemeType: "FHRS"},
		{RatingValue: "1", SchemeType: "FHRS"},
		{RatingValue: "2", SchemeType: "FHRS"},
		{RatingValue: "2", SchemeType: "FHRS"},
		{RatingValue: "0", SchemeType: "FHRS"},
		{RatingValue: "3", SchemeType: "FHRS"},
		{RatingValue: "AwaitingInspection", SchemeType: "FHRS"},
		{RatingValue: "5", SchemeType: "FHRS"},
		{RatingValue: "4", SchemeType: "FHRS"},
		{RatingValue: "Exempt", SchemeType: "FHRS"},
	}
}

func GetFakeFhisEstablishments() []model.Establishment {
	return []model.Establishment{
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Pass and Eat Safe", SchemeType: "FHIS"},
		{RatingValue: "Pass and Eat Safe", SchemeType: "FHIS"},
		{RatingValue: "Improvement Required", SchemeType: "FHIS"},
		{RatingValue: "Awaiting Inspection", SchemeType: "FHIS"},
		{RatingValue: "Exempt", SchemeType: "FHIS"},
		{RatingValue: "Pass", SchemeType: "FHIS"},
		{RatingValue: "Awaiting Inspection", SchemeType: "FHIS"},
		{RatingValue: "Awaiting Publication", SchemeType: "FHIS"},
	}
}
