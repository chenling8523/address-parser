package addressparser

type Address struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
	Area     string `json:"area"`
	Detail   string `json:"detail"`
}

func AddressParse() *Address {
	return nil
}
