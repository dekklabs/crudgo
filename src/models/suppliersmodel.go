package models

import (
	"database/sql"
	"entities"
)

// SuppliersConexion Conexion Db
type SuppliersConexion struct {
	Db *sql.DB
}

//FindAllModelSupplier Trae todos los Suppliers(Proveedores)
func (suppliersConexion SuppliersConexion) FindAllModelSupplier() (supplier []entities.Suppliers, err error) {
	rows, err := suppliersConexion.Db.Query("SELECT * FROM suppliers")
	if err != nil {
		return nil, err
	} else {
		var suppliers []entities.Suppliers
		for rows.Next() {
			var supplierID int64
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
			var homePage string

			err2 := rows.Scan(&supplierID, &companyName, &contactName, &contactTitle, &address, &city, &region,
				&postalCode,
				&country,
				&phone,
				&fax,
				&homePage)
			if err2 != nil {
				return nil, err2
			} else {
				supplier := entities.Suppliers{
					SupplierID:   supplierID,
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
					HomePage:     homePage,
				}
				suppliers = append(suppliers, supplier)
			}
		}
		return suppliers, nil
	}
}
