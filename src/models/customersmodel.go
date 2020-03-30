package models

import (
	"database/sql"
	"entities"
)

//CustomersModel Conexion al modelo
type CustomersModel struct {
	Db *sql.DB
}

//FindAllCustomer Modelo retorna la consulta al a base de datos de los empleados
func (customersModel CustomersModel) FindAllCustomer() (customer []entities.Customer, err error) {
	rows, err := customersModel.Db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	} else {
		var customers []entities.Customer
		for rows.Next() {
			var customerID string
			var companyName string
			var contactName string
			var contactTitle string
			var address string
			var city string
			var region string
			var postalCode string
			var country string
			var phone string
			var fax string

			err2 := rows.Scan(
				&customerID,
				&companyName,
				&contactName,
				&contactTitle,
				&address,
				&city,
				&region,
				&postalCode,
				&country,
				&phone,
				&fax)
			if err2 != nil {
				return nil, err2
			} else {
				customer := entities.Customer{
					CustomerID:   customerID,
					CompanyName:  companyName,
					ContactName:  contactName,
					ContactTitle: contactTitle,
					Address:      address,
					City:         city,
					Region:       region,
					PostalCode:   postalCode,
					Country:      country,
					Phone:        phone,
					Fax:          fax,
				}
				customers = append(customers, customer)
			}
		}
		return customers, nil
	}
}
