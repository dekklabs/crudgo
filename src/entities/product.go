package entities

//Product entidad modelo
type Product struct {
	ProductID       int64   `json:"productId"`
	ProductName     string  `json:"produtName"`
	SupplierID      int64   `json:"supplierId"`
	CategoryID      int64   `json:"categoryId"`
	QuantityPerUnit string  `json:"quantityPerUnit"`
	UnitPrice       float64 `json:"unitPrice"`
	UnitsInStock    int64   `json:"unitsInStock"`
	UnitsOnOrder    int64   `json:"unitsOnOrder"`
	ReorderLevel    int64   `json:"reorderLevel"`
	Discontinued    int64   `json:"discontinued"`
}
