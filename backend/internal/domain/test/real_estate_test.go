package test

import (
	"real-state-api/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validAddress() domain.Address {
	return domain.Address{
		ZipCode:      "98200000",
		Street:       "flores da cunha",
		Complement:   "apto 100",
		Neighborhood: "Centro",
		City:         "ibirubá",
		StateAbbr:    "RS",
	}
}

func TestRealEstatePropertyInput_IsValid_Sell(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "SELL",
		Address: validAddress(),
		Value:   350000,
	}
	errs := input.IsValid()
	assert.Empty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_Rent(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "RENT",
		Address: validAddress(),
		Value:   2500,
	}
	errs := input.IsValid()
	assert.Empty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_EmptyType(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "",
		Address: validAddress(),
		Value:   350000,
	}
	errs := input.IsValid()
	assert.NotEmpty(t, errs)
	assert.GreaterOrEqual(t, len(errs), 1)
}

func TestRealEstatePropertyInput_IsValid_InvalidType(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "EXCHANGE",
		Address: validAddress(),
		Value:   350000,
	}
	errs := input.IsValid()
	assert.NotEmpty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_InvalidAddress(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "SELL",
		Address: domain.Address{},
		Value:   350000,
	}
	errs := input.IsValid()
	assert.NotEmpty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_ZeroValue(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "SELL",
		Address: validAddress(),
		Value:   0,
	}
	errs := input.IsValid()
	assert.NotEmpty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_NegativeValue(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "RENT",
		Address: validAddress(),
		Value:   -100,
	}
	errs := input.IsValid()
	assert.NotEmpty(t, errs)
}

func TestRealEstatePropertyInput_IsValid_MultipleErrors(t *testing.T) {
	input := domain.RealEstatePropertyInput{
		Type:    "",
		Address: domain.Address{},
		Value:   -1,
	}
	errs := input.IsValid()
	assert.GreaterOrEqual(t, len(errs), 3)
}

func TestRealEstatePropertyInput_ToDb(t *testing.T) {
	addr := validAddress()
	input := domain.RealEstatePropertyInput{
		Type:    "SELL",
		Address: addr,
		Value:   500000,
		Image:   "http://teste.com",
	}
	prop := input.ToDb()

	require.NotNil(t, prop)
	assert.Equal(t, "SELL", prop.Type)
	assert.Equal(t, addr, prop.Address)
	assert.Equal(t, 500000.0, prop.Value)
	assert.Equal(t, "http://teste.com", prop.Image)
}

func TestRealEstateProperty_FromDb(t *testing.T) {
	addr := validAddress()
	prop := domain.RealEstateProperty{
		Id:      1,
		Type:    "RENT",
		Address: addr,
		Value:   2500,
		Image:   "http://teste.com",
	}
	input := prop.FromDb()

	assert.Equal(t, "RENT", input.Type)
	assert.Equal(t, addr, input.Address)
	assert.Equal(t, 2500.0, input.Value)
	assert.Equal(t, "http://teste.com", input.Image)
}
