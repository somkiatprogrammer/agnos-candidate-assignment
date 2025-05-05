package models

type Patient struct {
  ID int
  FirstNameTH string
  MiddleNameTH string
  LastNameTH string
  FirstNameEN string
  MiddleNameEN string
  LastNameEN string
  DateOfBirth string 
  PatientHN string
  NationalID int
  PassportID string
  PhoneNumber string
  Email string
  Gender string
}

type PatientSearchInput struct {
  FirstName string  `json:"first_name"`
  MiddleName string  `json:"middle_name"`
  LastName string  `json:"last_name"`
  DateOfBirth string  `json:"date_of_birth"`
  NationalID string  `json:"national_id"`
  PassportID string  `json:"passport_id"`
  PhoneNumber string  `json:"phone_number"`
  Email string  `json:"email"`
}