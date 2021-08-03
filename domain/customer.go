package domain

import (
	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking/dto"
)

type Customer struct {
	Id        int    `db:"id"`
	Company   string `db:"company"`
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
	JobTitle  string `db:"job_title"`
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:        c.Id,
		Company:   c.Company,
		LastName:  c.LastName,
		FirstName: c.FirstName,
		JobTitle:  c.JobTitle,
	}
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
