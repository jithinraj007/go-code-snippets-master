package dto

type CustomerResponse struct {
	Id        int    `json:"id"`
	Company   string `json:"company"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	JobTitle  string `json:"job_title"`
}
