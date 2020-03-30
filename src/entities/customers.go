package entities

//Customer modelo de empleado
type Customer struct {
	CustomerID   string `json:"customerId"`
	CompanyName  string `json:"companyName"`
	ContactName  string `json:"contactName"`
	ContactTitle string `json:"contactTitle"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postalCode"`
	Country      string `json:"country"`
	Phone        string `json:"phone"`
	Fax          string `json:"fax"`
}
