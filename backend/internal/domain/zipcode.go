package domain

import "errors"

type ZipCode struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`

	Error string `json:"erro"`
}

var (
	ErrNotFound   = errors.New("zipcode does not exist")
	ErrInternal   = errors.New("internal server error")
	ErrConnection = errors.New("error establishing connection")
	BadRequest    = errors.New("bad request to third party")
	ThirdPartyErr = errors.New("third party error")
)

type Address struct {
	ZipCode      string `json:"zip_code"`
	Street       string `json:"street"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	StateAbbr    string `json:"state_abbr"`
}

func (z ZipCode) ToAddress() *Address {
	return &Address{
		ZipCode:      z.Cep,
		Street:       z.Logradouro,
		Complement:   z.Complemento,
		Neighborhood: z.Bairro,
		City:         z.Localidade,
		StateAbbr:    z.Uf,
	}

}

func (a Address) IsValid() bool {

	if a.ZipCode == "" || a.City == "" || a.StateAbbr == "" || a.Street == "" || a.Neighborhood == "" {
		return false
	}

	return true
}
