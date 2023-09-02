package skeleton

import (
	"context"
	"net/http"
	"testing-nextalent/model"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetCountry(ctx context.Context, person string) (string, error)
	GetCountryAll(ctx context.Context) ([]model.Person, error)
	ScriptInsertData(ctx context.Context, person model.Person) error
}

type External interface {
	GetTimeZone(ctx context.Context, headers http.Header, timeZone string) (interface{}, error)
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data     Data
	external External
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data, external External) Service {
	// Assign variable dari parameter ke object
	return Service{
		data:     data,
		external: external,
	}
}

// GetCountry ...
func (s Service) GetCountry(ctx context.Context, person string) (interface{}, error) {

	var result interface{}
	var err error

	// END CHECK AUTH
	if person == "" {
		result, err = s.data.GetCountryAll(ctx)
	} else {
		result, err = s.data.GetCountry(ctx, person)
	}

	return result, err
}

// GetTimeZone ...
func (s Service) GetTimeZone(ctx context.Context, timeZone string) (interface{}, error) {

	// END CHECK AUTH
	result, err := s.external.GetTimeZone(ctx, http.Header{}, timeZone)

	return result, err
}

// Script ...
func (s Service) ScriptInsertData(ctx context.Context) (interface{}, error) {
	var err error
	datas := []model.Person{}

	datas = append(datas, model.Person{
		Name:    "Adam",
		Country: "Kuala Lumpur",
	})
	datas = append(datas, model.Person{
		Name:    "John",
		Country: "Singapore",
	})
	datas = append(datas, model.Person{
		Name:    "Henry",
		Country: "Singapore",
	})
	datas = append(datas, model.Person{
		Name:    "Dominic",
		Country: "Thailand",
	})

	// END CHECK AUTH
	for x := range datas {
		err = s.data.ScriptInsertData(ctx, datas[x])
	}

	return "Success", err
}
