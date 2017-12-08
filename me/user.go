package me

import (
	"github.com/toorop/govh"
)

// User in an OVH user
type User struct {
	FirstName                    string         `json:"firstname"`
	Name                         string         `json:"name"`
	Vat                          string         `json:"vat"`
	OvhSubsidiary                string         `json:"ovhSubsidiary"`
	Area                         string         `json:"area"`
	BirthDay                     govh.DateTime2 `json:"birthDay,omitempty"`
	NationalIdentificationNumber string         `json:"nationalIdentificationNumber"`
	SpareEmail                   string         `json:"spareEmail"`
	OvhCompany                   string         `json:"ovhCompany"`
	State                        string         `json:"state"`
	PhoneCountry                 string         `json:"phoneCountry"`
	Email                        string         `json:"email"`
	Currency                     struct {
		Symbol string `json:"symbol"`
		Code   string `json:"code"`
	} `json:"currency"`
	City                                string `json:"city"`
	Fax                                 string `json:"fax"`
	NicHandle                           string `json:"nichandle"`
	Address                             string `json:"address"`
	CompanyNationalIdentificationNumber string `json:"companyNationalIdentificationNumber"`
	BirthCity                           string `json:"birthCity"`
	Country                             string `json:"country"`
	Language                            string `json:"language"`
	Organisation                        string `json:"organisation"`
	Phone                               string `json:"phone"`
	Sex                                 string `json:"sex"`
	Zip                                 string `json:"zip"`
	CorporationType                     string `json:"corporationType"`
	CustomerCode                        string `json:"customerCode"`
	Legalform                           string `json:"legalform"`
}
