package models

type Customer struct {
	CustomerName   string `json:"requestNumber"`
	CustomerPhone  string `json:"customerPhone"`
	ProfilePicture string `json:"profilePicture"`
}
