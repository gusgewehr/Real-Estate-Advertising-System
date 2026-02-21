package domain

import (
	"errors"
	"time"
)

type RealEstateProperty struct {
	Id   uint
	Type string
	Address
	Value     float64
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type RealEstatePropertyInput struct {
	Type    string  `json:"type"`
	Address Address `json:"address"`
	Value   float64 `json:"value"`
	Image   string  `json:"image"`
}

func (r RealEstatePropertyInput) IsValid() []error {
	errs := make([]error, 0)
	if r.Type == "" {
		errs = append(errs, errors.New("type cannot be empty"))
	}
	if r.Type != "SELL" && r.Type != "RENT" {
		errs = append(errs, errors.New("type must be either SELL or RENT"))
	}

	if !r.Address.IsValid() {
		errs = append(errs, errors.New("address is invalid"))
	}

	if r.Value <= 0 {
		errs = append(errs, errors.New("value must be greater than zero"))
	}

	return errs
}

func (r RealEstatePropertyInput) ToDb() *RealEstateProperty {
	return &RealEstateProperty{
		Type:    r.Type,
		Address: r.Address,
		Value:   r.Value,
		Image:   r.Image,
	}
}

func (r RealEstateProperty) FromDb() RealEstatePropertyInput {
	return RealEstatePropertyInput{
		Type:    r.Type,
		Address: r.Address,
		Value:   r.Value,
		Image:   r.Image,
	}
}
