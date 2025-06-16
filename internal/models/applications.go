package models

type Application struct {
	FullName       string `json:"full_name" xml:"FullName"`
	Email          string `json:"email" xml:"Email"`
	DOB            string `json:"dob" xml:"DOB"`
	ProgramApplied string `json:"program_applied" xml:"ProgramApplied"`
}

type ApplicationWithID struct {
	ID             int    `json:"id"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	DOB            string `json:"dob"`
	ProgramApplied string `json:"program_applied"`
}
