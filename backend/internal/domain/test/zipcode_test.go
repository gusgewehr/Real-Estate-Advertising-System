package test

import (
	"real-estate-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddress_IsValid_AllFieldsFilled(t *testing.T) {
	addr := domain.Address{
		ZipCode:      "98200000",
		Street:       "flores da cunha",
		Complement:   "apto 100",
		Neighborhood: "centro",
		City:         "ibirubá",
		StateAbbr:    "RS",
	}
	assert.True(t, addr.IsValid())
}

func TestAddress_IsValid_MissingZipCode(t *testing.T) {
	addr := domain.Address{
		Street:       "flores da cunha",
		Neighborhood: "Centro",
		City:         "ibirubá",
		StateAbbr:    "RS",
	}
	assert.False(t, addr.IsValid())
}

func TestAddress_IsValid_MissingCity(t *testing.T) {
	addr := domain.Address{
		ZipCode:      "98200000",
		Street:       "flores da cunha",
		Neighborhood: "Centro",
		StateAbbr:    "RS",
	}
	assert.False(t, addr.IsValid())
}

func TestAddress_IsValid_MissingStateAbbr(t *testing.T) {
	addr := domain.Address{
		ZipCode:      "98200000",
		Street:       "flores da cunha",
		Neighborhood: "Centro",
		City:         "ibirubá",
	}
	assert.False(t, addr.IsValid())
}

func TestAddress_IsValid_MissingStreet(t *testing.T) {
	addr := domain.Address{
		ZipCode:      "98200000",
		Neighborhood: "Centro",
		City:         "ibirubá",
		StateAbbr:    "RS",
	}
	assert.False(t, addr.IsValid())
}

func TestAddress_IsValid_MissingNeighborhood(t *testing.T) {
	addr := domain.Address{
		ZipCode:   "98200000",
		Street:    "flores da cunha",
		City:      "ibirubá",
		StateAbbr: "RS",
	}
	assert.False(t, addr.IsValid())
}

func TestZipCode_ToAddress(t *testing.T) {
	zc := domain.ZipCode{
		Cep:         "98200000",
		Logradouro:  "flores da cunha",
		Complemento: "apto 100",
		Bairro:      "Centro",
		Localidade:  "ibirubá",
		Uf:          "RS",
	}

	addr := zc.ToAddress()

	assert.NotNil(t, addr)
	assert.Equal(t, "98200000", addr.ZipCode)
	assert.Equal(t, "flores da cunha", addr.Street)
	assert.Equal(t, "apto 100", addr.Complement)
	assert.Equal(t, "Centro", addr.Neighborhood)
	assert.Equal(t, "ibirubá", addr.City)
	assert.Equal(t, "RS", addr.StateAbbr)
}
