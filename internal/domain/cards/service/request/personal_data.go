package request

type PersonalData struct {
	Address    Address `json:"address,omitempty"`
	BirthDate  string  `json:"birth_date"`
	Document   string  `json:"document"`
	LastName   string  `json:"last_name"`
	MotherName string  `json:"mother_name"`
	Name       string  `json:"name"`
	Phone      Phone   `json:"phone,omitempty"`
}

type Phone struct {
	Code   string `json:"code"`
	Number string `json:"number"`
}
type Address struct {
	Number  string `json:"number"`
	ZipCode string `json:"zip_code"`
}
