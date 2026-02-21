package gateway

import (
	"fmt"
	"real-state-api/internal/domain"

	"resty.dev/v3"
)

type ZipCodeGateway struct {
}

func NewZipCodeGateway() *ZipCodeGateway {
	return &ZipCodeGateway{}
}

func (g *ZipCodeGateway) GetZipCode(zipCodeStr string) (*domain.ZipCode, error) {
	zipcode := &domain.ZipCode{}
	var err error

	c := resty.New()
	defer c.Close()

	res, err := c.R().
		SetPathParam("zipCode", zipCodeStr).
		SetResult(&zipcode).
		Get("https://viacep.com.br/ws/{zipCode}/json")
	if err != nil {
		return nil, domain.ErrConnection
	}

	if res.IsError() {
		if res.StatusCode() == 404 {
			return nil, domain.ErrNotFound
		}
		if res.StatusCode() == 400 {
			return nil, domain.BadRequest
		}
		if res.StatusCode() == 500 {
			return nil, domain.ThirdPartyErr
		}
		return nil, fmt.Errorf("status code error: %d", res.StatusCode())
	}

	if zipcode.Error == "true" {
		return nil, domain.ErrNotFound
	}

	return zipcode, nil
}
