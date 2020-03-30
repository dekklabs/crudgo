package models

import (
	"database/sql"
	"encoding/json"
	"entities"
	"net/http"
)

//ProductModel Model Producto Database
type ProductModel struct {
	Db *sql.DB
}

//FindAll retorna todos los datos
func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var productID int64
			var productName string
			var supplierID int64
			var categoryID int64
			var quantityPerUnit string
			var unitPrice float64
			var unitsInStock int64
			var unitsOnOrder int64
			var reorderLevel int64
			var discontinued int64
			err2 := rows.Scan(&productID, &productName, &supplierID, &categoryID, &quantityPerUnit, &unitPrice, &unitsInStock, &unitsOnOrder, &reorderLevel, &discontinued)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					ProductID:       productID,
					ProductName:     productName,
					SupplierID:      supplierID,
					CategoryID:      categoryID,
					QuantityPerUnit: quantityPerUnit,
					UnitPrice:       unitPrice,
					UnitsInStock:    unitsInStock,
					UnitsOnOrder:    unitsOnOrder,
					ReorderLevel:    reorderLevel,
					Discontinued:    discontinued,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

/* Tools */
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithjson(w, code, map[string]string{"error": msg})
}

func respondWithjson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
