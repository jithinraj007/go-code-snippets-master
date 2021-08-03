package domain

import (
	"database/sql"

	"github.com/ashishjuyal/banking-lib/errs"
	"github.com/ashishjuyal/banking-lib/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	premiumcustomers := make([]Customer, 0)

	if status == "" {
		findAllSql := "select id, company, last_name, first_name, job_title from customers"
		err = d.client.Select(&premiumcustomers, findAllSql)
	} else {
		findAllSql := "select id, company, last_name, first_name, job_title, status from customers where status = ?"
		err = d.client.Select(&premiumcustomers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return premiumcustomers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select id, company, last_name, first_name, email_address, job_title, status from customers where id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
