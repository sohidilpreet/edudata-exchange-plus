package models

type Application struct {
    FullName       string `json:"full_name" xml:"FullName"`
    Email          string `json:"email" xml:"Email"`
    DOB            string `json:"dob" xml:"DOB"`
    ProgramApplied string `json:"program_applied" xml:"ProgramApplied"`
}
