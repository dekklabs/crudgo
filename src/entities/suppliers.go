package entities

//Suppliers entidad de suppliers(Proveedores)
type Suppliers struct {
	SupplierID   int64  `json:"supplierId"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactname"`
	ContactTitle string `json:"contactTitle"`
	Address      string `json:"addres"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
	HomePage     string `json:"homepage"`
}
